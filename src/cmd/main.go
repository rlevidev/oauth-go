package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rlevidev/oauth-go/src/config/start_db"
	"github.com/rlevidev/oauth-go/src/routes"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("ℹ️  Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Inicializar banco de dados PostgreSQL
	log.Println("🔄 Inicializando conexão com PostgreSQL...")
	db, err := start_db.InitDB()
	if err != nil {
		log.Printf("❌ Erro ao conectar ao banco de dados: %v", err)
		log.Fatal("💀 Não foi possível conectar ao banco de dados. Verifique suas credenciais no arquivo .env")
	}

	log.Println("✅ Banco de dados inicializado com sucesso!")
	log.Printf("📊 Conectado ao banco: %s", os.Getenv("DB_NAME"))

	router := gin.Default()

	// Passar a conexão do banco para as rotas
	routes.InitRoutes(&router.RouterGroup, db)

	log.Printf("🚀 Servidor iniciado na porta %s", os.Getenv("PORT"))
	router.Run()
}

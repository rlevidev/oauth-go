package start_db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/rlevidev/oauth-go/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// validateDBConfig valida as configurações do banco de dados
func validateDBConfig() error {
	log.Println("🔍 Validando configurações do banco de dados...")

	// Verificar variáveis obrigatórias
	requiredVars := map[string]string{
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
	}

	var missingVars []string
	for varName, value := range requiredVars {
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}

	if len(missingVars) > 0 {
		return fmt.Errorf("❌ Variáveis de ambiente obrigatórias não definidas: %v\n💡 Verifique seu arquivo .env", missingVars)
	}

	// Validar se a porta é um número válido
	if port := os.Getenv("DB_PORT"); port != "" {
		if _, err := strconv.Atoi(port); err != nil {
			return fmt.Errorf("❌ DB_PORT deve ser um número válido, recebeu: %s", port)
		}
	}

	// Validar formato do host (não vazio após trim)
	if host := os.Getenv("DB_HOST"); host == "" {
		return errors.New("❌ DB_HOST não pode estar vazio")
	}

	log.Println("✅ Configurações validadas com sucesso!")
	return nil
}

func InitDB() (*gorm.DB, error) {
	// Validar configurações antes de conectar
	if err := validateDBConfig(); err != nil {
		return nil, err
	}

	// Configuração do PostgreSQL via variáveis de ambiente
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// DSN (Data Source Name) para PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	log.Println("Conectando ao banco PostgreSQL...")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Erro ao conectar ao banco PostgreSQL: %v", err)
		return nil, err
	}

	// Testar a conexão
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Erro ao obter conexão do banco: %v", err)
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		log.Printf("Erro ao fazer ping no banco PostgreSQL: %v", err)
		return nil, err
	}

	log.Println("✅ Conectado ao PostgreSQL com sucesso!")

	// Auto-migrate das tabelas
	err = db.AutoMigrate(
		&models.UserDomain{},
	)
	if err != nil {
		log.Printf("Erro ao migrar o banco de dados: %v", err)
		return nil, err
	}

	log.Println("✅ Migrações do banco de dados executadas com sucesso!")

	return db, nil
}

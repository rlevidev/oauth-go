package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/google/uuid"
)

type UserDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
}

func NewUserDomain(
	email string,
	password string,
	name string,
) *UserDomain {
	return &UserDomain{
		ID:       uuid.New().String(),
		Email:    email,
		Password: password,
		Name:     name}
}

func NewUserLoginDomain(
	email string,
	password string,
) *UserDomain {
	return &UserDomain{
		Email:    email,
		Password: password,
	}
}

func (ud *UserDomain) EncryptPassword() {
	encrypt := md5.New()
	defer encrypt.Reset()
	encrypt.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(encrypt.Sum(nil))
}

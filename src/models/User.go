package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/felipecesargomes/social-apiGO/src/security"
)

type User struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty" sql:"type:varchar(255)"` // Ajustar o tamanho do campo Password
	CreatedAt string `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("nick is required and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("email is required and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email is invalid")
	}

	if step == "add" && user.Password == "" {
		return errors.New("password is required and cannot be empty")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	if step == "add" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordWithHash)
	}
	return nil
}

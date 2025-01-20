package models

import (
	"errors"
	"strings"
)

type User struct {
	ID        int64  `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
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

	if step == "add" && user.Password == "" {
		return errors.New("password is required and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}

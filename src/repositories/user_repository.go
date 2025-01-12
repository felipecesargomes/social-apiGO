package repositories

import (
	"database/sql"

	"github.com/felipecesargomes/social-apiGO/src/models"
)

// Users repository
type Users struct {
	DB *sql.DB
}

func NewRepositoryUser(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) CreateUser(users models.User) (uint64, error) {
	statement, err := repository.DB.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(users.Name, users.Nick, users.Email, users.Password)

	if err != nil {
		return 0, err
	}

	lastIDInsert, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastIDInsert), nil
}

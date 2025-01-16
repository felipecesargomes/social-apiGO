package repositories

import (
	"database/sql"
	"fmt"

	"github.com/felipecesargomes/social-apiGO/src/models"
)

// Users repository
type Users struct {
	DB *sql.DB
}

func NewRepositoryUser(db *sql.DB) *Users {
	return &Users{DB: db}
}

func (repository Users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.DB.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastIDInsert, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastIDInsert), nil
}

func (repository Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	rows, err := repository.DB.Query(
		"SELECT id, name, nick, email FROM users WHERE name LIKE ? OR nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repository Users) FindById(id uint64) (models.User, error) {
	rows, err := repository.DB.Query("SELECT id, name, nick, email FROM users WHERE id = ?", id)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

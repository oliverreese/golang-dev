package repository

import (
	"github.com/oliverreese/golang-dev/pkg/database"
	"github.com/oliverreese/golang-dev/pkg/model"
	"fmt"
)

type UserRepository struct {
	DB database.DBInterface
}

func NewUserRepository(db *database.Database) *UserRepository {
	return &UserRepository{
		DB: db.DB,
	}
}

func (repo *UserRepository) FindOneById(id int) (*model.User, error) {
	user := model.User{}
	fmt.Println(user)
	fmt.Println(id)

	query := `SELECT user.id, user.firstname, user.lastname, user.age FROM user WHERE user.id = ?`


	err := repo.DB.QueryRow(query, id).Scan(
		&user.Id,
		&user.Firstname,
		&user.Lastname,
		&user.Age,
	)

	if err != nil {
		return &user, err
	}

	return &user, nil
}
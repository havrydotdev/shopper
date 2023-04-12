package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"shopper"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(user shopper.SignUpInput) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id", users)
	row := r.db.QueryRow(query, user.Username, user.Email, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthRepo) GetUser(email, password string) (shopper.User, error) {
	var user shopper.User
	query := fmt.Sprintf("SELECT * FROM %s u WHERE u.password = $1 AND u.email = $2", users)
	err := r.db.Get(&user, query, password, email)

	return user, err
}

func (r *AuthRepo) GetUserById(userId int) (shopper.User, error) {
	var user shopper.User
	query := fmt.Sprintf("SELECT * FROM %s u WHERE u.id = $1", users)
	err := r.db.Get(&user, query, userId)

	return user, err
}

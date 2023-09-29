package repository

import (
	"blogspot/entity"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) IUserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Login(user entity.User) (entity.User, error) {
	query := `SELECT * FROM users WHERE email=? and password=?`

	err := r.db.Get(&user, query, user.Email, user.Password)
	return user, err
}

func (r *userRepo) Register(user entity.User) (err error) {
	query := `INSERT INTO users(email, password, name, user_type) VALUES(:email, :password, :name, 1)`

	_, err = r.db.NamedExec(query, user)

	if err != nil {
		return
	}
	return
}

func (r *userRepo) Update(user entity.User) (err error) {
	query := `UPDATE users SET email=:email, password=:password, name=:name WHERE id=:id`

	_, err = r.db.NamedExec(query, user)

	if err != nil {
		return
	}
	return
}

package parser

import (
	"blogspot/entity"
	"net/http"
)

type userParser struct{}

func NewUserParser() IUserParser {
	return &userParser{}
}

func (p *userParser) ParseUserEntity(r *http.Request) (entity.User, error) {
	user := entity.User{}

	err := jsonParser(r, &user)
	if err != nil {
		return user, err
	}

	user.Password = HashPassword(user.Password)

	return user, nil
}

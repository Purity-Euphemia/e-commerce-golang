package services

import (
	"errors"

	"ecommerce-go/models"
	"ecommerce-go/repositories"
)

func RegisterUser(name, email, password string) (*models.User, error) {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := repositories.CreateUser(&user)
	return &user, err
}

func LoginUser(email, password string) (*models.User, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil || user.Password != password {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

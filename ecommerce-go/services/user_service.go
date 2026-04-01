package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"ecommerce-go/models"
	"ecommerce-go/repositories"
	"ecommerce-go/utils"
)

func RegisterUser(name, email, password string) (map[string]interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "customer",
	}
	err = repositories.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
		"token": token,
	}, nil
}

func LoginUser(email, password string) (map[string]interface{}, error) {
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
		"token": token,
	}, nil
}

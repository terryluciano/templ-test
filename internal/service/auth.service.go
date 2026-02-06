package service

import (
	"errors"

	"github.com/terryluciano/templ-test/internal/model"
	"github.com/terryluciano/templ-test/internal/repository"
	"github.com/terryluciano/templ-test/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

func AuthSignup(input *validation.SignupSchema) (*validation.UserResponse, error) {

	email_exists, err := repository.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if email_exists != nil {
		return nil, errors.New("email is already taken.")
	}

	// generate password hash with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  string(hashedPassword),
	}

	if err := repository.CreateUser(user); err != nil {
		return nil, err
	}

	userResponse := &validation.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return userResponse, nil

}

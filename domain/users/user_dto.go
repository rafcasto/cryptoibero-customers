package users

import (
	"src/github.com/rafaelc/cryptoibero-customers/utils/errors"
	"strings"
)

type User struct {
	ID       int64  `json:"id"`
	FirtName string `json:"firtName"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirtName = strings.TrimSpace(user.FirtName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
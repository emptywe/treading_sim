package model

import "errors"

type User struct {
	Uid       int     `json:"-" db:"id"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	UserName  string  `json:"user_name,omitempty"`
	FirstName string  `json:"firstName,omitempty" `
	LastName  string  `json:"lastName,omitempty" `
	Status    string  `json:"status" `
	Balance   float64 `json:"balance" `
}

type TUser struct {
	UserName string
	Balance  float64
}

type SignUpRequest struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	ID int `json:"id"`
}

type SignInRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type SignInResponse struct {
	ID      int    `json:"id"`
	Session string `json:"session"`
	Token   string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (user SignUpRequest) ValidateUser() error {

	if len([]rune(user.UserName)) < 2 {
		return errors.New("username must be at least two symbols")
	}

	if len([]rune(user.Password)) < 4 {
		return errors.New("password must be at least four symbols")
	}
	return nil
}

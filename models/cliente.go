package models

import "gopkg.in/validator.v2"

type Users struct {
	Id       int    `json:"id,omitempty"` //Caso seja vazio não pega no Json Marshall
	Username string `json:"username" validate:"nonzero"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email" validate:"regexp=^[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$"`
}

func ValidateDataUsers(users *Users) error {
	if err := validator.Validate(users); err != nil {
		return err
	}
	return nil
}

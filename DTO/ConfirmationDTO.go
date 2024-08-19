package DTO

import "Golang/Models"

type ConfirmationDTO struct {
	User  Models.User
	Token string
}

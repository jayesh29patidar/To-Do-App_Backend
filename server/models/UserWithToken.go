package models

type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

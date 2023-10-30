package models

import "time"

type UserModel struct {
	ID 			int
	Email 		string
	Username 	string
	Password 	string
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

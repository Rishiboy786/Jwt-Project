package models

import "time"

type User struct{
	ID  int					`json:"id"`
	First_name *string		`json:"first_name" validate:"required ,min=2, max=100"`
	Last_name	*string		`json:"last_name" validate:"required ,min=2, max=100"`
	Email		*string 	`json:"email" validate:"email, required"`
	Password	*string		`json:"password" validate:"required , min=6"`
	Phone		*string		`json:"phone" validate:"required "`
	Token		*string		`json:"token"`
	User_type	*string		`json:"user_type" validate:"required eq=ADMIN|eq=USER"`
	Refresh_token *string	`json:"refresh_token"`
	Created_at		time.Time	`json:"created_at"`

	Updated_at		time.Time	`json:"updated_at"`
	User_id			string		`json:"user_id"`
}
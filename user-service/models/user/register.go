package user

import "time"

// UserAccount --
type Register struct {
	Email       string    `json:"email" validate:"required,email"`
	Phone       string    `json:"phoneNumber" validate:"required,max=15,min=9"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Nickname    string    `json:"nickname"`
	Description string    `json:"description"`
	Photo       string    `json:"photo"`
	Password    string    `json:"password" validate:"required"`
	BirthDate   time.Time `json:"birthDate"`
	Gender      string    `json:"gender"`
}

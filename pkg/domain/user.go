package domain

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username" validate:"required,alphanum,min=3,max=25"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10"`
}

type UserSvc interface {
	GetUser(id int) (*User, error)
	ListUser(category string) ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

type UserDB interface {
	GetUser(id int) (*User, error)
	ListUser(category string) ([]*User, error)
	CreateUser(u *User) error
	DeleteUser(id int) error
}

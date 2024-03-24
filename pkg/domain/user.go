package domain

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
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

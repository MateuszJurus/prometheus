// db/user_db.go

package db

type User struct {
	ID    int    `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(user User) error {
	_, err := db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	return err
}

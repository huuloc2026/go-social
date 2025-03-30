package domain

type User struct {
	ID       int
	Username string
	Email    string
}

func NewUser(id int, username, email string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
	}
}

package appname

import "context"

type User struct {
	username string
	id       string
}

// UserService represents a service for managing users.
type UserService interface {
	FindUserByID(ctx context.Context, id int) (*User, error)
}

func Run() {
	for {

	}
}

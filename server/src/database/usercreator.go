package database

// UserCreator .
type UserCreator interface {
	CreateUser(string) (int, error)
	IsUserExist(string) bool
}

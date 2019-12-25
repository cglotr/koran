package database

// UserGetter .
type UserGetter interface {
	GetUser(string) (*User, error)
	GetUserByID(int) (*User, error)
}

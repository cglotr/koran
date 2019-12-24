package database

// UserGetter .
type UserGetter interface {
	GetUser(string) (User, error)
}

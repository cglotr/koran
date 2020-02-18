package database

import "github.com/cglotr/koran/server/src/structs"

// UserGetter .
type UserGetter interface {
	GetUser(string) (*structs.User, error)
	GetUserByID(int) (*structs.User, error)
}

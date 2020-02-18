package user

import "github.com/cglotr/koran/server/src/structs"

// Getter .
type Getter interface {
	Get(string) (*structs.User, error)
	GetByID(int) (*structs.User, error)
}

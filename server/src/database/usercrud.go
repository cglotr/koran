package database

// UserCRUD .
type UserCRUD interface {
	UserCreator
	UserGetter
	UserUpdater
}

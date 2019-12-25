package database

// UserQuranDeleter .
type UserQuranDeleter interface {
	DeleteUserQuran(int, int, int) error
}

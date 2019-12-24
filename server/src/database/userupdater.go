package database

// UserUpdater .
type UserUpdater interface {
	UpdateUserToken(int, string) bool
}

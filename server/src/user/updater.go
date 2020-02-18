package user

// Updater .
type Updater interface {
	UpdateToken(int, string) bool
}

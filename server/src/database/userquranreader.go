package database

// UserQuranReader .
type UserQuranReader interface {
	ReadUserQuran(int, int, int) (bool, error)
}

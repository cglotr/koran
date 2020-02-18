package user

// Creator .
type Creator interface {
	Create(string) (int, error)
	Exist(string) bool
}

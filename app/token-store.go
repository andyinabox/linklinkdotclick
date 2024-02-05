package app

type TokenStore interface {
	Create(uint) (string, error)
	Get(string) (uint, error)
	Delete(string) error
}

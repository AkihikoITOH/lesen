package backend

type Backend interface {
	Write(key string) error
	Read(key string) bool
	Erase(key string) error
}

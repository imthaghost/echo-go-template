package cache

// Service ...
type Service interface {
	Get() error
	Insert() error
}
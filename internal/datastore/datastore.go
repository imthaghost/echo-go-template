package datastore

// Service ...
type Service interface {
	Update() error
	Get() error
	Connect()
}
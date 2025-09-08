package storage

type Storage interface {
	CreateDevice(name string, manufacturer string, year int) (int64, error)
}

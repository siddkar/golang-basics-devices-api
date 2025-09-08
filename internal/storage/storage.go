package storage

import (
	"devices-api/internal/types/entities"
)

type Storage interface {
	CreateDevice(name string, manufacturer string, year int) (int64, error)
	GetDeviceById(id int64) (entities.Device, error)
}

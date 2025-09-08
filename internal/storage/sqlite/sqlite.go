package sqlite

import (
	"database/sql"
	"devices-api/internal/config"
	"devices-api/internal/types/entities"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS devices (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		manufacturer TEXT,
		year INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateDevice(name string, manufacturer string, year int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO devices (name, manufacturer, year) Values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, manufacturer, year)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Sqlite) GetDeviceById(id int64) (entities.Device, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, manufacturer, year FROM devices WHERE id = ? LIMIT 1")
	if err != nil {
		return entities.Device{}, err
	}
	defer stmt.Close()

	var device entities.Device

	err = stmt.QueryRow(id).Scan(&device.Id, &device.Name, &device.Manufacturer, &device.Year)

	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Device{}, fmt.Errorf("no student passed with id: %s", fmt.Sprint(id))
		}
		return entities.Device{}, fmt.Errorf("query error: %w", err)
	}

	return device, nil
}

package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Определяем инстанс
type Storage struct {
	config *Config
	// DataBase FileDescriptor
	db *sql.DB
}

//Storage constructor
func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

//Open connection method
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfully!")
	return nil
}

//Close connection method
func (storage *Storage) Close() {
	storage.db.Close()
}

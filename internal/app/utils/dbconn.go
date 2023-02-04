package utils

import (
	// golang package
	"fmt"
	"log"

	// external package
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

var (
	// mock sql
	sqlxOpen = sqlx.Open
)

// InitDBConn will initialize connection to database.
func InitDBConn(cfg configuration.DatabaseConfig) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sqlxOpen(cfg.Driver, psqlInfo)
	if err != nil {
		log.Printf("[initDBConn] sql.Open() got error: %+v\n", err)
		return nil, err
	}

	log.Println("successfully connect to database")
	return db, nil
}

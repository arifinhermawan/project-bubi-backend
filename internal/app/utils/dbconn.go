package utils

import (
	// golang package
	"database/sql"
	"fmt"
	"log"

	// external package
	_ "github.com/lib/pq"

	// internal package
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
)

var (
	// mock sql
	sqlOpen = sql.Open
)

// InitDBConn will initialize connection to database.
func InitDBConn(cfg configuration.DatabaseConfig) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sqlOpen(cfg.Driver, psqlInfo)
	if err != nil {
		log.Printf("[initDBConn] sql.Open() got error: %+v\n", err)
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Printf("[initDBConn] db.Ping() got error: %+v\n", err)
		return err
	}

	log.Println("successfully connect to database - initDBConn")
	return nil
}

package app

import (
	// golang package

	"log"

	// internal package
	"github.com/arifinhermawan/bubi/internal/app/server"
	"github.com/arifinhermawan/bubi/internal/app/utils"
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
)

var (
	configPath = "files/"
)

func NewApplication() {

	// init infra
	infra := server.NewInfra()

	// init db connection
	dbConfig := infra.Config.GetConfig().Database
	errDBConn := utils.InitDBConn(dbConfig)
	if errDBConn != nil {
		log.Fatalf("[NewApplication] utils.InitDBConn() got an error: %+v", errDBConn)
	}

	// init repo
	_ = pgsql.NewRepository()

	// init app stack
	_ = server.NewResource()
	_ = server.NewService()
	_ = server.NewUsecase()
	handlers := server.NewHandler()
	log.Println("Successfully initialize app stack! - NewApplication")

	// register handler
	utils.HandleRequest(handlers)
}

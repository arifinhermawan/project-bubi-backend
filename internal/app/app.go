package app

import (
	// golang package
	"log"

	// internal package
	"github.com/arifinhermawan/bubi/internal/app/server"
	"github.com/arifinhermawan/bubi/internal/app/utils"
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
	"github.com/arifinhermawan/bubi/internal/infrastructure/golang"
	reader "github.com/arifinhermawan/bubi/internal/infrastructure/reader"
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
)

func NewApplication() {

	cfg := configuration.NewConfiguration()
	golang := golang.NewGolang()
	reader := reader.NewReader()

	// init infra
	infraParam := server.InfraParam{
		Config: cfg,
		Golang: golang,
		Reader: reader,
	}

	infra := server.NewInfra(infraParam)

	// init db connection
	dbConfig := infra.Config.GetConfig().Database
	dbConn, errDBConn := utils.InitDBConn(dbConfig)
	if errDBConn != nil {
		log.Fatalf("[NewApplication] utils.InitDBConn() got an error: %+v", errDBConn)
	}
	defer dbConn.Close()

	// init repo
	dbRepo := pgsql.NewDBRepository(pgsql.DBRepoParam{
		Infra: infra,
		DB:    dbConn,
	})

	// ----------------
	// |init app stack|
	// ----------------

	// init resources
	resourceParam := server.ResourceParam{
		DB: dbRepo,
	}
	resources := server.NewResource(resourceParam)

	// init services
	services := server.NewService(resources)

	// init usecases
	useCases := server.NewUsecase(services)

	// init handlers
	handlers := server.NewHandler(useCases, infra)
	log.Println("Successfully initialize app stack!")

	// register handler
	utils.HandleRequest(handlers)
}

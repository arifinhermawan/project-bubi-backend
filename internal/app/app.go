package app

import (
	// golang package
	"log"

	// internal package
	"github.com/arifinhermawan/bubi/internal/app/server"
	"github.com/arifinhermawan/bubi/internal/app/utils"
	"github.com/arifinhermawan/bubi/internal/infrastructure/authentication"
	"github.com/arifinhermawan/bubi/internal/infrastructure/configuration"
	"github.com/arifinhermawan/bubi/internal/infrastructure/golang"
	reader "github.com/arifinhermawan/bubi/internal/infrastructure/reader"
	"github.com/arifinhermawan/bubi/internal/repository/pgsql"
	"github.com/arifinhermawan/bubi/internal/repository/redis"
)

func NewApplication() {

	cfg := configuration.NewConfiguration()
	auth := authentication.NewAuth(cfg)
	golang := golang.NewGolang()
	reader := reader.NewReader()

	// init infra
	infraParam := server.InfraParam{
		Auth:   auth,
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

	// init redis connection
	redisConfig := infra.Config.GetConfig().Redis
	redisConn, errRedisConn := utils.InitRedisConn(&redisConfig)
	if errRedisConn != nil {
		log.Fatalf("[NewApplication] utils.InitRedisConn() got an error: %+v", errRedisConn)
	}

	// init repo
	dbRepo := pgsql.NewDBRepository(pgsql.DBRepoParam{
		Infra: infra,
		DB:    dbConn,
	})

	// init redis
	redisRepo := redis.NewRedisRepository(redis.RedisRepositoryParam{
		Infra: infra,
		Redis: redisConn,
	})

	// ----------------
	// |init app stack|
	// ----------------

	// init resources
	resourceParam := server.ResourceParam{
		Cache: redisRepo,
		Infra: infra,
		DB:    dbRepo,
	}
	resources := server.NewResource(resourceParam)

	// init services
	services := server.NewService(resources, infra)

	// init usecases
	useCases := server.NewUsecase(services)

	// init handlers
	handlers := server.NewHandler(useCases, infra)
	log.Println("Successfully initialize app stack!")

	// register handler
	utils.HandleRequest(infra, handlers)
}

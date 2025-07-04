package main

import (
	"fmt"

	"github.com/Komilov31/l0/cmd/api"
	"github.com/Komilov31/l0/db"
	"github.com/Komilov31/l0/internal/config"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		panic("could not initialize logger")
	}
	defer logger.Sync()

	dbConfig := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBName,
	)

	pgxpool, err := db.NewSqlStorage(dbConfig)
	if err != nil {
		logger.Error("could not create db instance")
		panic("could not create db instance")
	}

	err = db.InitStorage(pgxpool)
	if err != nil {
		logger.Error("could not initialize db")
		panic("could not initialize db instance " + err.Error())
	}
	logger.Info("sucessfully initialized db")

	apiServer := api.NewAPIServer(":8081", pgxpool, logger)
	apiServer.Run()
}

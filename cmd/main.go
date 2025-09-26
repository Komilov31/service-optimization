package main

import (
	"fmt"
	"log"

	"net/http"
	_ "net/http/pprof"

	"github.com/Komilov31/l0/cmd/api"
	"github.com/Komilov31/l0/db"
	"github.com/Komilov31/l0/internal/config"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("could not initialize logger")
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
		log.Fatal("could not create db instance: ", err)
	}

	err = db.InitStorage(pgxpool)
	if err != nil {
		log.Fatal("could not initialize db instance " + err.Error())
	}
	logger.Info("sucessfully initialized db")

	go func() {
		logger.Info("starting server for pprof on :5000")
		if err := http.ListenAndServe(":5000", http.DefaultServeMux); err != nil {
			logger.Error("could not start pprof server")
		}
	}()

	apiServer := api.NewAPIServer(":8080", pgxpool, logger)
	apiServer.Run()
}

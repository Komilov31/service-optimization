package api

import (
	"github.com/Komilov31/l0/internal/cache"
	"github.com/Komilov31/l0/internal/handler"
	"github.com/Komilov31/l0/internal/kafka"
	"github.com/Komilov31/l0/internal/repository"
	"github.com/Komilov31/l0/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type APIServer struct {
	addr   string
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewAPIServer(addr string, db *pgxpool.Pool, logger *zap.Logger) *APIServer {
	return &APIServer{addr: addr, db: db, logger: logger}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	router.LoadHTMLFiles("/app/static/index.html")
	router.Static("/static", "/app/static")

	repository := repository.New(s.db, s.logger)
	inMemoryCache := cache.New(repository)
	inMemoryCache.LoadFromDbToCache()
	service := service.New(repository, inMemoryCache)
	handler := handler.New(service)

	kafkaConsumer := kafka.NewConsumer(s.logger, service)
	go kafkaConsumer.Consume()

	router.GET("order/:order_uid", handler.GetOrderById)
	router.GET("order/", handler.GetMainPage)

	s.logger.Info("listening on " + s.addr)

	return router.Run(s.addr)
}

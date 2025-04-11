package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc"
	"github.com/sMARCHz/go-secretaria-finance/internal/config"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"

	_ "github.com/lib/pq"
)

func main() {
	// Setting up
	zapLog := logger.NewZapLogger()
	logger.Init(zapLog.Sugar())

	config.LoadConfig()

	db, closeDBConnection := createDBConnection()

	// Starting server
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	grpcServer, shutdown := grpc.NewGRPCServer(db)
	go func() {
		grpcServer.Start()
	}()

	// Shutdown server
	s := <-stopCh
	logger.Infof("Got signal '%v', attempting graceful shutdown", s)
	shutdown()
	closeDBConnection()
	logger.Info("Gracefully shutting down...")
}

func createDBConnection() (*sqlx.DB, func()) {
	cfg := config.Get().DB
	datasourceName := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)
	db, err := sqlx.Open(cfg.Driver, datasourceName)
	if err != nil {
		logger.Fatal("cannot connect to database: ", err)
	}
	return db, func() {
		err := db.Close()
		if err != nil {
			logger.Fatal("failed to close db connection: ", err)
		}
	}
}

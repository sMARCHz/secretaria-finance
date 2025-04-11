package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc"
	"github.com/sMARCHz/go-secretaria-finance/internal/config"
	"github.com/sMARCHz/go-secretaria-finance/internal/logger"

	_ "github.com/lib/pq"
)

var (
	apiType = flag.String("type", "grpc", "Type of API to run")
)

func main() {
	logger := logger.New()
	config := config.LoadConfig(logger, ".")

	// create database connection
	db, closeDBConnection := createDBConnection(config.DB, logger)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)

	if *apiType == "rest" {
		// TODO
	} else {
		grpcServer, shutdown := grpc.NewGRPCServer(config.App, logger, db)
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
}

func createDBConnection(config config.DatabaseConfiguration, logger logger.Logger) (*sqlx.DB, func()) {
	datasourceName := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v", config.Username, config.Password, config.Host, config.Port, config.DBName, config.SSLMode)
	db, err := sqlx.Open(config.Driver, datasourceName)
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

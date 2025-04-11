package grpc

import (
	"fmt"
	"net"

	"github.com/jmoiron/sqlx"
	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driven/db"
	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
	"github.com/sMARCHz/go-secretaria-finance/internal/config"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/services"
	"github.com/sMARCHz/go-secretaria-finance/pkg/logger"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	server   *grpc.Server
	database *sqlx.DB
}

func NewGRPCServer(db *sqlx.DB) (GRPCServer, func()) {
	server := grpc.NewServer(grpc.EmptyServerOption{})
	return GRPCServer{
		server:   server,
		database: db,
	}, server.GracefulStop
}

func (g GRPCServer) Start() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Get().App.Port))
	if err != nil {
		logger.Fatal("failed to listen: ", err)
	}

	// Register service to grpc server
	repo := db.NewFinanceRepository(g.database)
	service := services.NewFinanceService(repo)
	pb.RegisterFinanceServiceServer(g.server, newFinanceServiceServer(service))

	// Start server
	logger.Infof("Starting gRPC server at :%v...", config.Get().App.Port)
	g.server.Serve(lis)
}

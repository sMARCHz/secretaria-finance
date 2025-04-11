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
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	return GRPCServer{
		server:   grpcServer,
		database: db,
	}, grpcServer.GracefulStop
}

func (g GRPCServer) Start() {
	// Wiring
	fRepo := db.NewFinanceRepository(g.database)
	fService := services.NewFinanceService(fRepo)
	fServer := newFinanceServiceServer(fService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Get().App.Port))
	if err != nil {
		logger.Fatal("failed to listen: ", err)
	}

	// Register service to grpc server
	pb.RegisterFinanceServiceServer(g.server, fServer)

	// Start server
	logger.Infof("Starting gRPC server at :%v...", config.Get().App.Port)
	g.server.Serve(lis)
}

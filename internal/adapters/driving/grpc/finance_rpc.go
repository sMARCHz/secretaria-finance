package grpc

import (
	"context"

	"github.com/sMARCHz/go-secretaria-finance/internal/adapters/driving/grpc/pb"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/dto"
	"github.com/sMARCHz/go-secretaria-finance/internal/core/services"
	"github.com/sMARCHz/go-secretaria-finance/internal/utils"
	"google.golang.org/protobuf/types/known/emptypb"
)

type financeServiceServer struct {
	pb.UnimplementedFinanceServiceServer
	service services.FinanceService
}

func newFinanceServiceServer(service services.FinanceService) *financeServiceServer {
	return &financeServiceServer{
		service: service,
	}
}

func (f *financeServiceServer) Withdraw(ctx context.Context, r *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	req := &dto.TransactionRequest{
		AccountName: r.AccountName,
		Category:    r.Category,
		Description: r.Description,
		Amount:      r.Amount,
	}
	response, err := f.service.Withdraw(req)
	if err != nil {
		return &pb.TransactionResponse{
			Status: int32(err.StatusCode),
			Error:  err.Message,
		}, utils.ConvertHttpErrToGRPC(err)
	}
	return response.ToProto(), nil
}

func (f *financeServiceServer) Deposit(ctx context.Context, r *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	req := &dto.TransactionRequest{
		AccountName: r.AccountName,
		Category:    r.Category,
		Description: r.Description,
		Amount:      r.Amount,
	}
	response, err := f.service.Deposit(req)
	if err != nil {
		return &pb.TransactionResponse{
			Status: int32(err.StatusCode),
			Error:  err.Message,
		}, utils.ConvertHttpErrToGRPC(err)
	}
	return response.ToProto(), nil
}

func (f *financeServiceServer) Transfer(ctx context.Context, r *pb.TransferRequest) (*pb.TransferResponse, error) {
	req := &dto.TransferRequest{
		FromAccountName: r.FromAccountName,
		ToAccountName:   r.ToAccountName,
		Description:     r.Description,
		Amount:          r.Amount,
	}
	response, err := f.service.Transfer(req)
	if err != nil {
		return &pb.TransferResponse{
			Status: int32(err.StatusCode),
			Error:  err.Message,
		}, utils.ConvertHttpErrToGRPC(err)
	}
	return response.ToProto(), nil
}

func (f *financeServiceServer) GetBalance(ctx context.Context, r *emptypb.Empty) (*pb.GetBalanceResponse, error) {
	response, err := f.service.GetBalance()

	accountsBalance := make([]*pb.AccountBalance, len(response))
	for i, v := range response {
		accountsBalance[i] = v.ToProto()
	}

	if err != nil {
		return &pb.GetBalanceResponse{
			Status: int32(err.StatusCode),
			Error:  err.Message,
		}, utils.ConvertHttpErrToGRPC(err)
	}

	// TODO: Fix this
	return &pb.GetBalanceResponse{Accounts: accountsBalance}, nil
}

func (f *financeServiceServer) GetOverviewStatement(ctx context.Context, r *pb.OverviewStatementRequest) (*pb.OverviewStatementResponse, error) {
	req := &dto.GetOverviewStatementRequest{
		From: r.From.AsTime(),
		To:   r.To.AsTime(),
	}
	response, err := f.service.GetOverviewStatement(req)
	if err != nil {
		return &pb.OverviewStatementResponse{
			Status: int32(err.StatusCode),
			Error:  err.Message,
		}, utils.ConvertHttpErrToGRPC(err)
	}
	return response.ToProto(), nil
}

func (f *financeServiceServer) GetOverviewMonthlyStatement(ctx context.Context, r *emptypb.Empty) (*pb.OverviewStatementResponse, error) {
	response, err := f.service.GetOverviewMonthlyStatement()
	if err != nil {
		return &pb.OverviewStatementResponse{
			Status: int32(err.StatusCode),
			Error:  err.Message,
		}, utils.ConvertHttpErrToGRPC(err)
	}
	return response.ToProto(), nil
}

func (f *financeServiceServer) GetOverviewAnnualStatement(ctx context.Context, r *emptypb.Empty) (*pb.OverviewStatementResponse, error) {
	response, err := f.service.GetOverviewAnnualStatement()
	if err != nil {
		return &pb.OverviewStatementResponse{
			Status: int32(err.StatusCode),
			Error:  err.Message,
		}, utils.ConvertHttpErrToGRPC(err)
	}
	return response.ToProto(), nil
}

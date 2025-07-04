package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"readly/middleware/auth"
	"readly/pb/readl
	"readly/usecase"
)

type BookServerImpl struct {
	pb.UnimplementedBookServiceServer
	maker           auth.TokenMaker
	registerUseCase usecase.RegisterBookUseCase
	deleteUseCase   usecase.DeleteBookUseCase
}

func NewBookServer(
	maker auth.TokenMaker,
	registerUseCase usecase.RegisterBookUseCase,
	deleteUseCase usecase.DeleteBookUseCase,
) *BookServerImpl {
	return &BookServerImpl{
		maker:           maker,
		registerUseCase: registerUseCase,
		deleteUseCase:   deleteUseCase,
	}
}

func (b *BookServerImpl) RegisterBook(ctx context.Context, req *pb.RegisterBookRequest) (*pb.RegisterBookResponse, error) {
	claims, err := auth.Authenticate(ctx, b.maker)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// TODO:バリデーション

	args := usecase.NewRegisterBookRequest(claims.UserID, req)
	res, err := b.registerUseCase.RegisterBook(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return res.ToProto(), nil
}

func (b *BookServerImpl) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	claims, err := auth.Authenticate(ctx, b.maker)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	args := usecase.NewDeleteBookRequest(claims.UserID, req)
	err = b.deleteUseCase.DeleteBook(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return &emptypb.Empty{}, nil
}

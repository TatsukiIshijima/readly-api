package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"readly/feature/book/usecase"
	"readly/middleware/auth"
	"readly/pb/readly/v1"
)

type BookServerImpl struct {
	pb.UnimplementedBookServiceServer
	maker              auth.TokenMaker
	registerUseCase    usecase.RegisterBookUseCase
	deleteUseCase      usecase.DeleteBookUseCase
	getBookUseCase     usecase.GetBookUseCase
	getBookListUseCase usecase.GetBookListUseCase
}

func NewBookServer(
	maker auth.TokenMaker,
	registerUseCase usecase.RegisterBookUseCase,
	deleteUseCase usecase.DeleteBookUseCase,
	getBookUseCase usecase.GetBookUseCase,
	getBookListUseCase usecase.GetBookListUseCase,
) *BookServerImpl {
	return &BookServerImpl{
		maker:              maker,
		registerUseCase:    registerUseCase,
		deleteUseCase:      deleteUseCase,
		getBookUseCase:     getBookUseCase,
		getBookListUseCase: getBookListUseCase,
	}
}

func (b *BookServerImpl) RegisterBook(ctx context.Context, req *pb.RegisterBookRequest) (*pb.RegisterBookResponse, error) {
	claims, err := auth.AuthenticateGRPC(ctx, b.maker)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	args := usecase.NewRegisterBookRequest(claims.UserID, req)
	res, err := b.registerUseCase.RegisterBook(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return res.ToProto(), nil
}

func (b *BookServerImpl) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	claims, err := auth.AuthenticateGRPC(ctx, b.maker)
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

func (b *BookServerImpl) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	claims, err := auth.AuthenticateGRPC(ctx, b.maker)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	args := usecase.NewGetBookRequest(claims.UserID, req.BookId)
	res, err := b.getBookUseCase.GetBook(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return res.ToProto(), nil
}

func (b *BookServerImpl) GetBookList(ctx context.Context, req *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	claims, err := auth.AuthenticateGRPC(ctx, b.maker)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	args := usecase.NewGetBookListRequest(claims.UserID, req.Limit, req.Offset)
	res, err := b.getBookListUseCase.GetBookList(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return res.ToProto(), nil
}

package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"readly/middleware"
	"readly/pb"
	"readly/service/auth"
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

func (b *BookServerImpl) RegisterBook(ctx context.Context, req *pb.RegisterBookRequest) (*pb.Book, error) {
	claims, err := middleware.Authenticate(ctx, b.maker)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	// TODO:バリデーション

	args := usecase.RegisterBookRequest{
		UserID:        claims.UserID,
		Title:         req.GetTitle(),
		Genres:        req.GetGenres(),
		Description:   req.GetDescription(),
		CoverImageURL: req.GetCoverImageUrl(),
		URL:           req.GetUrl(),
		AuthorName:    req.GetAuthorName(),
		PublisherName: req.GetPublisherName(),
		PublishDate:   req.GetPublishDate(),
		ISBN:          req.GetIsbn(),
		Status:        req.GetReadingStatus(),
		StartDate:     req.GetStartDate(),
		EndDate:       req.GetEndDate(),
	}
	book, err := b.registerUseCase.RegisterBook(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return &pb.Book{
		Id:            book.ID,
		Title:         book.Title,
		Genres:        book.Genres,
		Description:   book.Description,
		CoverImageUrl: book.CoverImageURL,
		Url:           book.URL,
		AuthorName:    book.AuthorName,
		PublisherName: book.PublisherName,
		PublishDate:   book.PublishDate,
		Isbn:          book.ISBN,
		ReadingStatus: pb.ReadingStatus(book.Status),
		StartDate:     book.StartDate,
		EndDate:       book.EndDate,
	}, nil
}

func (b *BookServerImpl) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	claims, err := middleware.Authenticate(ctx, b.maker)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	// TODO:ユーザが登録したBookであるかチェック（認可 Authorization）

	args := usecase.DeleteBookRequest{
		UserID: claims.UserID,
		BookID: req.GetBookId(),
	}
	err = b.deleteUseCase.DeleteBook(ctx, args)
	if err != nil {
		return nil, gRPCStatusError(err)
	}
	return &emptypb.Empty{}, nil
}

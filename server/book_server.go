package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"readly/entity"
	"readly/middleware"
	"readly/pb/readly/v1"
	"readly/service/auth"
	"readly/usecase"
	"readly/util"
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

func (b *BookServerImpl) toReadingStatusEntity(status pb.ReadingStatus) entity.ReadingStatus {
	switch status {
	case pb.ReadingStatus_UNREAD:
		return entity.Unread
	case pb.ReadingStatus_READING:
		return entity.Reading
	case pb.ReadingStatus_DONE:
		return entity.Done
	default:
		return entity.Unknown
	}
}

func (b *BookServerImpl) RegisterBook(ctx context.Context, req *pb.RegisterBookRequest) (*pb.Book, error) {
	claims, err := middleware.Authenticate(ctx, b.maker)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	// TODO:バリデーション

	args := usecase.RegisterBookRequest{
		UserID:        claims.UserID,
		Title:         req.GetTitle(),
		Genres:        req.GetGenres(),
		Description:   util.ToStringOrNil(req.GetDescription()),
		CoverImageURL: util.ToStringOrNil(req.GetCoverImageUrl()),
		URL:           util.ToStringOrNil(req.GetUrl()),
		AuthorName:    util.ToStringOrNil(req.GetAuthorName()),
		PublisherName: util.ToStringOrNil(req.GetPublisherName()),
		PublishDate:   entity.NewDateEntityFromProto(req.GetPublishDate()),
		ISBN:          util.ToStringOrNil(req.GetIsbn()),
		Status:        b.toReadingStatusEntity(req.GetReadingStatus()),
		StartDate:     entity.NewDateEntityFromProto(req.GetStartDate()),
		EndDate:       entity.NewDateEntityFromProto(req.GetEndDate()),
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
		PublishDate:   book.PublishDate.ToProto(),
		Isbn:          book.ISBN,
		ReadingStatus: pb.ReadingStatus(book.Status),
		StartDate:     book.StartDate.ToProto(),
		EndDate:       book.EndDate.ToProto(),
	}, nil
}

func (b *BookServerImpl) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	claims, err := middleware.Authenticate(ctx, b.maker)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
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

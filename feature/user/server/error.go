package server

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"readly/feature/user/usecase"
)

func gRPCStatusError(err error) error {
	var e *usecase.Error
	if !errors.As(err, &e) {
		return status.Error(codes.Internal, err.Error())
	}
	switch e.StatusCode {
	case usecase.BadRequest:
		return status.Error(codes.InvalidArgument, e.Message)
	case usecase.UnAuthorized:
		return status.Error(codes.Unauthenticated, e.Message)
	case usecase.Forbidden:
		return status.Error(codes.PermissionDenied, e.Message)
	case usecase.NotFound:
		return status.Error(codes.NotFound, e.Message)
	case usecase.Conflict:
		return status.Error(codes.AlreadyExists, e.Message)
	case usecase.Internal:
		return status.Error(codes.Internal, e.Message)
	default:
		return status.Error(codes.Internal, e.Message)
	}
}

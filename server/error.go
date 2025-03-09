package server

import (
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"readly/usecase"
)

func gRPCStatusError(err error) error {
	var e *usecase.Error
	if !errors.As(err, &e) {
		return status.Errorf(codes.Internal, err.Error())
	}
	switch e.StatusCode {
	case usecase.BadRequest:
		return status.Errorf(codes.InvalidArgument, e.Message)
	case usecase.UnAuthorized:
		return status.Errorf(codes.Unauthenticated, e.Message)
	case usecase.Forbidden:
		return status.Errorf(codes.PermissionDenied, e.Message)
	case usecase.NotFound:
		return status.Errorf(codes.NotFound, e.Message)
	case usecase.Conflict:
		return status.Errorf(codes.AlreadyExists, e.Message)
	case usecase.Internal:
		return status.Errorf(codes.Internal, e.Message)
	default:
		return status.Errorf(codes.Internal, e.Message)
	}
}

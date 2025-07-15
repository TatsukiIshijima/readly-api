package usecase

import (
	"net"
	pb "readly/pb/readly/v1"
	"readly/util"
)

type SignInRequest struct {
	Email     string
	Password  string
	IPAddress string
	UserAgent string
}

func NewSignInRequest(proto *pb.SignInRequest, ipAddress, userAgent string) SignInRequest {
	return SignInRequest{
		Email:     proto.GetEmail(),
		Password:  proto.GetPassword(),
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
}

func (r SignInRequest) Validate() error {
	// Email validation
	if len(r.Email) == 0 {
		return newError(BadRequest, InvalidRequestError, "email is required")
	}
	if err := util.StringValidator(r.Email).ValidateEmail(); err != nil {
		return newError(BadRequest, InvalidRequestError, "email has invalid format")
	}
	if err := util.StringValidator(r.Email).ValidateNoSQLInjection(); err != nil {
		return newError(BadRequest, InvalidRequestError, "email contains potentially dangerous content")
	}

	// Password validation
	if len(r.Password) == 0 {
		return newError(BadRequest, InvalidRequestError, "password is required")
	}
	if err := util.StringValidator(r.Password).ValidatePassword(); err != nil {
		return newError(BadRequest, InvalidRequestError, err.Error())
	}
	if err := util.StringValidator(r.Password).ValidateNoSQLInjection(); err != nil {
		return newError(BadRequest, InvalidRequestError, "password contains potentially dangerous content")
	}

	// IPAddress validation (optional but validate if provided)
	if len(r.IPAddress) > 0 {
		if net.ParseIP(r.IPAddress) == nil {
			return newError(BadRequest, InvalidRequestError, "IP address has invalid format")
		}
		if err := util.StringValidator(r.IPAddress).ValidateNoSQLInjection(); err != nil {
			return newError(BadRequest, InvalidRequestError, "IP address contains potentially dangerous content")
		}
	}

	// UserAgent validation (optional but validate length if provided)
	if len(r.UserAgent) > 0 {
		if err := util.StringValidator(r.UserAgent).ValidateLength(0, 2048); err != nil {
			return newError(BadRequest, InvalidRequestError, "user agent must be less than 2048 characters")
		}
		if err := util.StringValidator(r.UserAgent).ValidateNoSQLInjection(); err != nil {
			return newError(BadRequest, InvalidRequestError, "user agent contains potentially dangerous content")
		}
	}

	return nil
}

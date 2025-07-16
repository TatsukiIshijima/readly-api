package usecase

import (
	"net"
	pb "readly/pb/readly/v1"
	"readly/util"
)

type SignUpRequest struct {
	Name      string
	Email     string
	Password  string
	IPAddress string
	UserAgent string
}

func NewSignUpRequest(proto *pb.SignUpRequest, ipAddress, userAgent string) SignUpRequest {
	return SignUpRequest{
		Name:      proto.GetName(),
		Email:     proto.GetEmail(),
		Password:  proto.GetPassword(),
		IPAddress: ipAddress,
		UserAgent: userAgent,
	}
}

func (r SignUpRequest) Validate() error {
	// Name validation
	if len(r.Name) == 0 {
		return newError(BadRequest, InvalidRequestError, "name is required")
	}
	if err := util.StringValidator(r.Name).ValidateLength(1, 100); err != nil {
		return newError(BadRequest, InvalidRequestError, "name must be between 1 and 100 characters")
	}

	// Email validation
	if len(r.Email) == 0 {
		return newError(BadRequest, InvalidRequestError, "email is required")
	}
	if err := util.StringValidator(r.Email).ValidateEmail(); err != nil {
		return newError(BadRequest, InvalidRequestError, "email has invalid format")
	}

	// Password validation
	if len(r.Password) == 0 {
		return newError(BadRequest, InvalidRequestError, "password is required")
	}
	if err := util.StringValidator(r.Password).ValidatePassword(); err != nil {
		return newError(BadRequest, InvalidRequestError, err.Error())
	}

	// IPAddress validation (optional but validate if provided)
	if len(r.IPAddress) > 0 {
		if net.ParseIP(r.IPAddress) == nil {
			return newError(BadRequest, InvalidRequestError, "IP address has invalid format")
		}
	}

	// UserAgent validation (optional but validate length if provided)
	if len(r.UserAgent) > 0 {
		if err := util.StringValidator(r.UserAgent).ValidateLength(0, 2048); err != nil {
			return newError(BadRequest, InvalidRequestError, "user agent must be less than 2048 characters")
		}
	}

	return nil
}

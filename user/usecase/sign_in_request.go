package usecase

import pb "readly/pb/readly/v1"

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

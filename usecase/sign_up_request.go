package usecase

import pb "readly/pb/readly/v1"

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

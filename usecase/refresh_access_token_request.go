package usecase

import pb "readly/pb/readly/v1"

type RefreshAccessTokenRequest struct {
	RefreshToken string
}

func NewRefreshAccessTokenRequest(proto *pb.RefreshTokenRequest) RefreshAccessTokenRequest {
	return RefreshAccessTokenRequest{
		RefreshToken: proto.GetRefreshToken(),
	}
}

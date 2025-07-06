package usecase

import pb "readly/pb/readly/v1"

type RefreshAccessTokenResponse struct {
	AccessToken string
}

func NewRefreshAccessTokenResponse(accessToken string) *RefreshAccessTokenResponse {
	return &RefreshAccessTokenResponse{
		AccessToken: accessToken,
	}
}

func (r *RefreshAccessTokenResponse) ToProto() *pb.RefreshTokenResponse {
	return &pb.RefreshTokenResponse{
		AccessToken: r.AccessToken,
	}
}

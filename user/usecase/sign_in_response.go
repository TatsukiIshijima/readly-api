package usecase

import (
	pb "readly/pb/readly/v1"
	"readly/user/repository"
)

type SignInResponse struct {
	AccessToken  string
	RefreshToken string
	UserID       int64
	Name         string
	Email        string
}

func NewSignInResponse(accessToken, refreshToken string, userRes *repository.GetUserResponse) *SignInResponse {
	return &SignInResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userRes.ID,
		Name:         userRes.Name,
		Email:        userRes.Email,
	}
}

func (r *SignInResponse) ToProto() *pb.SignInResponse {
	return &pb.SignInResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
		UserId:       r.UserID,
		Name:         r.Name,
		Email:        r.Email,
	}
}

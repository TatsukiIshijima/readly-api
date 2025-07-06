package usecase

import (
	pb "readly/pb/readly/v1"
	"readly/user/repository"
)

type SignUpResponse struct {
	AccessToken  string
	RefreshToken string
	UserID       int64
	Name         string
	Email        string
}

func NewSignUpResponse(accessToken, refreshToken string, userRes *repository.CreateUserResponse) *SignUpResponse {
	return &SignUpResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userRes.ID,
		Name:         userRes.Name,
		Email:        userRes.Email,
	}
}

func (r *SignUpResponse) ToProto() *pb.SignUpResponse {
	return &pb.SignUpResponse{
		AccessToken:  r.AccessToken,
		RefreshToken: r.RefreshToken,
		UserId:       r.UserID,
		Name:         r.Name,
		Email:        r.Email,
	}
}

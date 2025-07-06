package usecase

import pb "readly/pb/readly/v1"

type DeleteBookRequest struct {
	UserID int64
	BookID int64
}

func NewDeleteBookRequest(userID int64, proto *pb.DeleteBookRequest) DeleteBookRequest {
	return DeleteBookRequest{
		UserID: userID,
		BookID: proto.GetBookId(),
	}
}

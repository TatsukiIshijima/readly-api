package repository

type GetUserByEmailRequest struct {
	Email string
}

func NewGetUserByEmailRequest(email string) GetUserByEmailRequest {
	return GetUserByEmailRequest{
		Email: email,
	}
}

type GetUserByIDRequest struct {
	ID int64
}

func NewGetUserByIDRequest(id int64) GetUserByIDRequest {
	return GetUserByIDRequest{
		ID: id,
	}
}

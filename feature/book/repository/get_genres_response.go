package repository

type GetAllGenresResponse struct {
	Genres []string
}

func newGetAllGenresResponse(genres []string) *GetAllGenresResponse {
	return &GetAllGenresResponse{
		Genres: genres,
	}
}

type GetGenresByBookIDResponse struct {
	Genres []string
}

func newGetGenresByBookIDResponse(genres []string) *GetGenresByBookIDResponse {
	return &GetGenresByBookIDResponse{
		Genres: genres,
	}
}

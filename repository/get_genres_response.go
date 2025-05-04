package repository

type GetGenresByBookIDResponse struct {
	Genres []string
}

func newGetGenresByBookIDResponse(genres []string) *GetGenresByBookIDResponse {
	return &GetGenresByBookIDResponse{
		Genres: genres,
	}
}

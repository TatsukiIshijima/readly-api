package entity

type Book struct {
	ID            int64         `json:"id"`
	Title         string        `json:"title"`
	Genres        []string      `json:"genres"`
	Description   *string       `json:"description"`
	CoverImageURL *string       `json:"cover_image_url"`
	URL           *string       `json:"url"`
	AuthorName    *string       `json:"author_name"`
	PublisherName *string       `json:"publisher_name"`
	PublishDate   *Date         `json:"publish_date"`
	ISBN          *string       `json:"isbn"`
	Status        ReadingStatus `json:"status"`
	StartDate     *Date         `json:"start_date"`
	EndDate       *Date         `json:"end_date"`
}

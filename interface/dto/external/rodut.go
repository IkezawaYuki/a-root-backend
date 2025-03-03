package external

type CreatePostRequest struct {
	ApiKey        string `json:"api_key"`
	Email         string `json:"email"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	FeaturedMedia int    `json:"featured_media"`
}

type CreatePostResponse struct {
	PostId  string `json:"post_id"`
	PostUrl string `json:"post_url"`
}

type UploadMediaResponse struct {
	ID        int    `json:"id"`
	SourceUrl string `json:"source_url"`
	MimeType  string `json:"mime_type"`
}

type TitleResponse struct {
	Title string `json:"title"`
}

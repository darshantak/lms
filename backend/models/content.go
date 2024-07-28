package models

type Content struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	FileType   string `json:"file_type"`
	UploadedBy string `json:"uploaded_by"`
	URL        string `json:"url"`
	ExpiresAt  string `json:"expires_at"`
	CreatedAt  string `json:"created_at"`
}

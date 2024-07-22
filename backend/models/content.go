package models

type Content struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	FileType string `json:"file_type"`
	URL      string `json:"url"`
}

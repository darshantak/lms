package models

type Assessment struct {
	ID        int    `json:"id"`
	ContentID string `json:"content_id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
}

package models

type Training struct {
	Title         string   `json:"title"`
	ScopedTeams   []string `json:"scoped_teams"`
	ScopedContent []string `json:"scoped_content"`
}

package github

import (
	"time"
)

// IssuesURL issues's url
const IssuesURL = "http://api.github.com/search/issues"

// IssuesSearchResult 数量和issue结构
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue 结构体
type Issue struct {
	Number   int
	HTMLURL  string `json:"html_url"`
	Title    string
	State    string
	User     *User
	CreateAt time.Time `json:"created_at"`
	Body     string
}

// User 用户
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

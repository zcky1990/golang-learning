package models

// Article - Our struct for all articles
type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

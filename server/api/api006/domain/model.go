package domain

import "time"

// RequestParam リクエストパラメータs
type RequestParam struct {
	UserID   int    `json:"user_id"`
	ImageURL string `json:"image_url"`
}

// Pictures is for picture table
type Pictures struct {
	PictureID   int       `json:"picture_id"`
	UserID      int       `json:"user_id"`
	ImageURL    string    `json:"image_url"`
	PublishedAt time.Time `json:"published_at"`
}

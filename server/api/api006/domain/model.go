package domain

import "time"

// RequestParam リクエストパラメータs
type RequestParam struct {
	UserID               int    `json:"user_id"`
	ImageURL             string `json:"image_url"`
	ViewCategoryCd       string `json:"view_category_cd"`
	PrefectureCategoryCd string `json:"prefecture_category_cd"`
	ImageNote            string `json:"image_note"`
}

// Picture is for picture table
type Picture struct {
	PictureID            int       `json:"picture_id"`
	UserID               int       `json:"user_id"`
	ImageURL             string    `json:"image_url"`
	PrefectureCategoryCd string    `json:"prefecture_category_cd"`
	ViewCategoryCd       string    `json:"view_category_cd"`
	ImageNote            string    `json:"image_note"`
	PublishedAt          time.Time `json:"published_at"`
}

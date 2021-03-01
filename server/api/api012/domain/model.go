package domain

// RequestParam リスエストパラメータ
type RequestParam struct {
	PictureID *int `validate:"required" json:"picture_id"`
}

// Picture contains each picture
type Picture struct {
	PictureID   int    `json:"picture_id"`
	UserID      int    `json:"user_id"`
	ImageURL    string `json:"image_url"`
	PublichedAt string `json:"published_at"`
}

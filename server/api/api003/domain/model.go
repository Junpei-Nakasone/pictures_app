package domain

// RequestParam リスエストパラメータ
type RequestParam struct {
	PrefectureCd *string `validate:"required" json:"prefecture_cd"`
}

// Picture contains each picture
type Picture struct {
	PictureID   int    `json:"picture_id"`
	UserID      int    `json:"user_id"`
	ImageURL    string `json:"image_url"`
	PublichedAt string `json:"published_at"`
}

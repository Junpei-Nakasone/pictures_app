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
	PublishedAt string `json:"published_at"`
	UserName    string `json:"user_name"`
	IconImage   string `json:"icon_image"`
	ImageNote   string `json:"image_note"`
}

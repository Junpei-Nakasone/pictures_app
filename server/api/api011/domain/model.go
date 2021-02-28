package domain

// RequestParam リスエストパラメータ
type RequestParam struct {
	UserID *int `validate:"required" json:"user_id"`
}

// UserData ユーザー情報
type UserData struct {
	UserID         int    `json:"user_id"`
	UserName       string `json:"user_name"`
	Note           string `json:"note"`
	IconImage      string `json:"icon_image"`
	PostedPictures []Picture
}

// Picture contains each picture
type Picture struct {
	PictureID   int    `json:"picture_id"`
	UserID      int    `json:"user_id"`
	ImageURL    string `json:"image_url"`
	PublichedAt string `json:"published_at"`
}

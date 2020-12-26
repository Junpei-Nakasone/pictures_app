package domain

// Picture contains each picture
type Picture struct {
	PictureID int    `json:"picture_id"`
	UserID    int    `json:"user_id"`
	ImageURL  string `json:"image_url"`
	// timestamp型のデータの取得しかたわからん
	PublichedAt string `json:"published_at"`
}

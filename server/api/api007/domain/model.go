package domain

// RequestParam リクエストパラメータ
type RequestParam struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// ResponseParam レスポンスデータ
type ResponseParam struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Note      string `json:"note"`
	IconImage string `json:"icond_image"`
}

package domain

// RequestParam リクエストパラメータ
type RequestParam struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

// UserData ユーザーデータ取得用
type UserData struct {
	UserID       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	Note         string `json:"note"`
	IconImage    string `json:"icond_image"`
}

// ResponseParam レスポンスデータ
type ResponseParam struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Note      string `json:"note"`
	IconImage string `json:"icond_image"`
}

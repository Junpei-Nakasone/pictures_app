package domain

// RequestParam リクエストパラメータ
type RequestParam struct {
	UserName     *string `json:"user_name"`
	Password     *string `json:"password"`
	EmailAddress *string `json:"email_address"`
	Note         *string `json:"note"`
}

// User ユーザー情報格納用の構造体
type User struct {
	UserID       *int    `gorm:"column:user_id;primary_key" json:"user_id";primary_key"`
	UserName     *string `json:"user_name"`
	UserPassword *string `json:"user_password"`
	EmailAddress *string `json:"mail_address"`
	Note         *string `json:"note"`
	IconImage    *string `json:"icon_image"`
}

// ResponseParam レスポンスデータ
type ResponseParam struct {
	UserID    *int    `json:"user_id"`
	UserName  *string `json:"user_name"`
	Note      *string `json:"note"`
	IconImage *string `json:"icon_image"`
}

package domain

// ResponseParam レスポンスデータ
type ResponseParam struct {
	ViewCategoryCd string `json:"view_category_cd"`
	ViewName       string `json:"view_name"`
	SortNo         int    `json:"sort_no"`
}

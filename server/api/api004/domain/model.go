package domain

// ResponseParam レスポンスデータ
type ResponseParam struct {
	PrefectureCategoryCd string `json:"prefecture_category_cd"`
	PrefectureName       string `json:"PrefectureCategoryCd"`
	SortNo               int    `json:"sort_no"`
}

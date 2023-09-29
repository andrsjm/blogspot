package entity

type Tag struct {
	ID  int    `json:"id" db:"id"`
	Tag string `json:"tag" db:"tag"`
}

type TagFilter struct {
	BlogID int `json:"blog_id" db:"blog_id"`
}

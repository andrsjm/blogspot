package entity

type Category struct {
	ID       int    `json:"id" db:"id"`
	Category string `json:"category" db:"category"`
}

type CategoryFilter struct {
	BlogID int `json:"blog_id" db:"blog_id"`
}

package entity

type Blog struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Content  string `json:"content" db:"content"`
	UserID   int    `json:"user_id" db:"user_id"`
	Hidden   int    `json:"hidden" db:"hidden"`
	Tag      []int  `json:"tags"`
	Category []int  `json:"categories"`
}

type BlogResponse struct {
	ID       int        `json:"id" db:"id"`
	Title    string     `json:"title" db:"title"`
	Content  string     `json:"content" db:"content"`
	Author   string     `json:"author" db:"name"`
	Tag      []Tag      `json:"tags"`
	Category []Category `json:"categories"`
}

type BlogFilter struct {
	ID       int    `json:"id"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
	Tag      string `json:"tag"`
	Category string `json:"category"`
	Hidden   int    `json:"hidden"`
	UserID   int    `json:"user_id"`
}

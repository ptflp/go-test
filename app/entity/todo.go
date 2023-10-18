package entity

type Todo struct {
	ID          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	IsCompleted bool   `json:"is_completed" db:"is_completed"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	DueDate     string `json:"due_date" db:"due_date"`
}

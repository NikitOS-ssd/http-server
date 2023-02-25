package todo

type TodoList struct {
	ID          int    `json:"-"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UserList struct {
	ID     int
	UserID int
	ListID int
}

type TodoItem struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	ID     int
	ListID int
	ItemID int
}

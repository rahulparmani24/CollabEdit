package ws

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	Type     string `json:"type"`     // e.g., "insert", "delete"
	Position int    `json:"position"` // Position for the operation
}

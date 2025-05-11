package entities

type User struct {
	ID        int    `json:"id"`
	ChatID    int    `json:"chat_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

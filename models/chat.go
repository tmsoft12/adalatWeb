package models

type ChatModel struct {
	ID        uint   `json:"id"`
	User_Id   string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Rule      string `json:"rule"`
}

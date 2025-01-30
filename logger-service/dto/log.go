package dto

type LogRequest struct {
	Content string `json:"content"`
}

type LogResponse struct {
	ID        any    `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

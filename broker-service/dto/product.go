package dto

type ProductSaveRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
}

type ProductUpdateRequest struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
}

type ProductResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Stock       uint   `json:"stock"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

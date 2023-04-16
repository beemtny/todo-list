package dto

type CreateTaskRequest struct {
	Description string `json:"description"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

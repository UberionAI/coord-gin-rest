package model

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DataResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Source  string      `json:"source"`
}

type MockItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Timestamp   string `json:"timestamp"`
}

package model

import "time"

// HealthResponse represents health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

// DataItem represents a generic data item from Elasticsearch (mock for now)
type DataItem struct {
	ID        string                 `json:"id"`
	Name      string                 `json:"name"`
	Value     string                 `json:"value"`
	Timestamp time.Time              `json:"timestamp"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// DataResponse represents API response with data items
type DataResponse struct {
	Success bool       `json:"success"`
	Data    []DataItem `json:"data"`
	Count   int        `json:"count"`
	Mock    bool       `json:"mock"` // Indicates mock data
}

// ErrorResponse represents error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    string `json:"code,omitempty"`
}

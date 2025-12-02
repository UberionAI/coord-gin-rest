package db

import (
	"time"

	"github.com/UberionAI/coord-gin-rest/internal/logger"
	"github.com/UberionAI/coord-gin-rest/internal/model"
)

// Client represents database client (Elasticsearch in future)
type Client struct {
	log  *logger.Logger
	mock bool // Indicates if using mock data
}

// NewClient creates a new database client
// For now, returns a mock client. Will be replaced with real ES client later.
func NewClient(log *logger.Logger) (*Client, error) {
	log.Warn("Database client running in MOCK mode (Elasticsearch not configured)")

	return &Client{
		log:  log,
		mock: true,
	}, nil
}

// Close closes database connection
func (c *Client) Close() error {
	if c.mock {
		c.log.Info("Mock database client closed")
		return nil
	}
	// Future: close real ES connection
	return nil
}

// GetData retrieves data from database (mock implementation)
func (c *Client) GetData() ([]model.DataItem, error) {
	if c.mock {
		return c.getMockData(), nil
	}

	// Future: implement real Elasticsearch query
	// Example:
	// res, err := c.esClient.Search(
	//     c.esClient.Search.WithIndex("your-index"),
	//     c.esClient.Search.WithBody(strings.NewReader(query)),
	// )

	return nil, nil
}

// getMockData returns mock data for testing
func (c *Client) getMockData() []model.DataItem {
	return []model.DataItem{
		{
			ID:        "mock-1",
			Name:      "Sample Item 1",
			Value:     "Mock value 1",
			Timestamp: time.Now(),
			Metadata: map[string]interface{}{
				"source": "mock",
				"type":   "test",
			},
		},
		{
			ID:        "mock-2",
			Name:      "Sample Item 2",
			Value:     "Mock value 2",
			Timestamp: time.Now().Add(-1 * time.Hour),
			Metadata: map[string]interface{}{
				"source": "mock",
				"type":   "test",
			},
		},
		{
			ID:        "mock-3",
			Name:      "Sample Item 3",
			Value:     "Mock value 3",
			Timestamp: time.Now().Add(-2 * time.Hour),
			Metadata: map[string]interface{}{
				"source": "mock",
				"type":   "test",
			},
		},
	}
}

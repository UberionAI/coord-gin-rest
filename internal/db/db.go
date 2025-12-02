package db

import (
	"time"

	"github.com/UberionAI/coord-gin-rest/internal/model"
	"github.com/rs/zerolog/log"
)

type Client struct {
	isMock bool
}

func NewClient(esHost string) *Client {
	if esHost == "" {
		log.Warn().Msg("Elasticsearch not configured, using mock data")
		return &Client{isMock: true}
	}

	// TODO: Реальное подключение к ES когда будут реквизиты
	log.Warn().Msg("Elasticsearch client initialization skipped, using mock data")
	return &Client{isMock: true}
}

func (c *Client) FetchData() ([]model.MockItem, error) {
	if c.isMock {
		return c.getMockData(), nil
	}

	// TODO: Реальный запрос к ES
	return nil, nil
}

func (c *Client) getMockData() []model.MockItem {
	return []model.MockItem{
		{
			ID:          "mock-001",
			Name:        "Sample Item 1",
			Description: "This is mock data from stub",
			Timestamp:   time.Now().Format(time.RFC3339),
		},
		{
			ID:          "mock-002",
			Name:        "Sample Item 2",
			Description: "Another mock entry",
			Timestamp:   time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
		},
	}
}

func (c *Client) IsMock() bool {
	return c.isMock
}

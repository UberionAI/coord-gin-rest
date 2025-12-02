package service

import (
	"github.com/UberionAI/coord-gin-rest/internal/db"
	"github.com/UberionAI/coord-gin-rest/internal/model"
)

type Service struct {
	dbClient *db.Client
}

func New(dbClient *db.Client) *Service {
	return &Service{
		dbClient: dbClient,
	}
}

func (s *Service) GetData() (*model.DataResponse, error) {
	items, err := s.dbClient.FetchData()
	if err != nil {
		return nil, err
	}

	source := "elasticsearch"
	if s.dbClient.IsMock() {
		source = "mock"
	}

	return &model.DataResponse{
		Success: true,
		Data:    items,
		Source:  source,
	}, nil
}

package handler

import (
	"net/http"

	"github.com/UberionAI/coord-gin-rest/internal/model"
	"github.com/UberionAI/coord-gin-rest/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, model.HealthResponse{
		Status:  "healthy",
		Message: "Service is running",
	})
}

func (h *Handler) GetData(c *gin.Context) {
	data, err := h.service.GetData()
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch data")
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/models"
	"github.com/KozlovNikolai/crud-cors-midlw-zap-gin/repository"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmployerHandler struct {
	logger *zap.Logger
	repo   repository.EmployerRepository
}

func NewEmployerHandler(logger *zap.Logger, repo repository.EmployerRepository) *EmployerHandler {
	return &EmployerHandler{logger: logger, repo: repo}
}

func (h *EmployerHandler) CreateEmployer(c *gin.Context) {
	var employer models.Employer
	if err := c.ShouldBindJSON(&employer); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.repo.CreateEmployer(context.Background(), employer)
	if err != nil {
		h.logger.Error("Error creating employer", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	employer.ID = id
	c.JSON(http.StatusCreated, employer)
}

func (h *EmployerHandler) GetEmployer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	employer, err := h.repo.GetEmployerByID(context.Background(), id)
	if err != nil {
		h.logger.Error("Error getting employer", zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Employer not found"})
		return
	}
	c.JSON(http.StatusOK, employer)
}

func (h *EmployerHandler) GetAllEmployers(c *gin.Context) {
	employers, err := h.repo.GetAllEmployers(context.Background())
	if err != nil {
		h.logger.Error("Error getting all employers", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employers)
}

func (h *EmployerHandler) UpdateEmployer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employer models.Employer
	if err := c.ShouldBindJSON(&employer); err != nil {
		h.logger.Error("Error binding JSON", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.repo.UpdateEmployer(context.Background(), id, employer)
	if err != nil {
		h.logger.Error("Error updating employer", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Updated successfully"})
}

func (h *EmployerHandler) DeleteEmployer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.repo.DeleteEmployer(context.Background(), id)
	if err != nil {
		h.logger.Error("Error deleting employer", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Deleted successfully"})
}

package controllers

import (
	"net/http"

	"github.com/example/amenity-poc/internal/models"
	"github.com/example/amenity-poc/internal/services"
	"github.com/gin-gonic/gin"
)

type AmenityRequestController struct {
	svc *services.AmenityRequestService
}

func NewAmenityRequestController(s *services.AmenityRequestService) *AmenityRequestController {
	return &AmenityRequestController{svc: s}
}

func (c *AmenityRequestController) Create(ctx *gin.Context) {
	reservationID := ctx.Param("id")
	var req models.AmenityRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.ReservationID = reservationID
	if err := c.svc.Create(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, req)
}

func (c *AmenityRequestController) ListForReservation(ctx *gin.Context) {
	reservationID := ctx.Param("id")
	list, err := c.svc.GetByReservation(reservationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// ListReservations returns reservation ids that have amenity requests
func (c *AmenityRequestController) ListReservations(ctx *gin.Context) {
	ids, err := c.svc.GetAllReservations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"reservations": ids})
}

func (c *AmenityRequestController) UpdateStatus(ctx *gin.Context) {
	id := ctx.Param("requestId")
	var payload struct {
		Status models.AmenityStatus `json:"status"`
	}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.UpdateStatus(id, payload.Status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *AmenityRequestController) GetByID(ctx *gin.Context) {
	id := ctx.Param("requestId")
	req, err := c.svc.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, req)
}

package router

import (
	"github.com/example/amenity-poc/internal/controllers"
	"github.com/example/amenity-poc/internal/db"
	"github.com/example/amenity-poc/internal/repositories"
	"github.com/example/amenity-poc/internal/services"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine, d *db.DB) {
	arRepo := repositories.NewAmenityRequestRepository(d)
	arSvc := services.NewAmenityRequestService(d, arRepo)
	arCtrl := controllers.NewAmenityRequestController(arSvc)

	v1 := r.Group("/api/v1")
	{
		// List reservations that have amenity requests
		v1.GET("/reservations", arCtrl.ListReservations)
		v1.POST("/reservations/:id/amenity-requests", arCtrl.Create)
		v1.GET("/reservations/:id/amenity-requests", arCtrl.ListForReservation)
		v1.GET("/amenity-requests/:requestId", arCtrl.GetByID)
		v1.PATCH("/amenity-requests/:requestId/status", arCtrl.UpdateStatus)
	}

	// Health check
	r.GET("/api/v1/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}

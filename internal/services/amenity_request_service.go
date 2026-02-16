package services

import (
	"errors"
	"time"

	"github.com/example/amenity-poc/internal/db"
	"github.com/example/amenity-poc/internal/models"
	"github.com/example/amenity-poc/internal/repositories"
	"github.com/google/uuid"
)

var (
	ErrInvalidQuantity     = errors.New("quantity must be greater than 0")
	ErrInvalidAmenityType  = errors.New("invalid amenity type")
	ErrReservationNotFound = errors.New("reservation not found or not checked-in")
)

type AmenityRequestService struct {
	repo *repositories.AmenityRequestRepository
	db   *db.DB
}

func NewAmenityRequestService(d *db.DB, r *repositories.AmenityRequestRepository) *AmenityRequestService {
	return &AmenityRequestService{repo: r, db: d}
}

func isValidAmenityType(t models.AmenityType) bool {
	switch t {
	case models.TOWELS, models.PILLOWS, models.TOILETRIES, models.ROOM_SERVICE, models.OTHER:
		return true
	default:
		return false
	}
}

// Here we assume reservation exists and is checked-in; in a real app we'd query reservations table.
func (s *AmenityRequestService) Create(req *models.AmenityRequest) error {
	if req.Quantity <= 0 {
		return ErrInvalidQuantity
	}
	if !isValidAmenityType(req.AmenityType) {
		return ErrInvalidAmenityType
	}

	// TODO: validate reservation exists and is checked-in. For now assume pass.

	req.ID = uuid.New().String()
	now := time.Now()
	req.RequestedAt = now
	req.CreatedAt = now
	req.UpdatedAt = now
	if req.Status == "" {
		req.Status = models.PENDING
	}

	return s.repo.Create(req)
}

func (s *AmenityRequestService) GetByReservation(reservationID string) ([]models.AmenityRequest, error) {
	return s.repo.GetByReservationID(reservationID)
}

func (s *AmenityRequestService) GetByID(id string) (*models.AmenityRequest, error) {
	return s.repo.GetByID(id)
}

// GetAllReservations returns the list of reservation IDs which have amenity requests
func (s *AmenityRequestService) GetAllReservations() ([]string, error) {
	return s.repo.GetAllReservations()
}

func (s *AmenityRequestService) UpdateStatus(id string, status models.AmenityStatus) error {
	// basic validation
	if status != models.PENDING && status != models.IN_PROGRESS && status != models.COMPLETED && status != models.CANCELLED {
		return errors.New("invalid status")
	}
	if status == models.COMPLETED {
		// set completed_at in repository or separate query; repo only updates status/updated_at for simplicity
	}
	return s.repo.UpdateStatus(id, string(status))
}

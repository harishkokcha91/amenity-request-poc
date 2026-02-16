package repositories

import (
	"time"

	"github.com/example/amenity-poc/internal/db"
	"github.com/example/amenity-poc/internal/models"
)

type AmenityRequestRepository struct {
	db *db.DB
}

func NewAmenityRequestRepository(d *db.DB) *AmenityRequestRepository {
	return &AmenityRequestRepository{db: d}
}

func (r *AmenityRequestRepository) Create(req *models.AmenityRequest) error {
	query := `INSERT INTO amenity_requests (id, reservation_id, guest_id, property_id, amenity_type, description, quantity, status, requested_at, completed_at, created_at, updated_at, created_by)
VALUES (:id, :reservation_id, :guest_id, :property_id, :amenity_type, :description, :quantity, :status, :requested_at, :completed_at, :created_at, :updated_at, :created_by)`
	_, err := r.db.NamedExec(query, req)
	return err
}

func (r *AmenityRequestRepository) GetByID(id string) (*models.AmenityRequest, error) {
	var req models.AmenityRequest
	err := r.db.Get(&req, "SELECT * FROM amenity_requests WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func (r *AmenityRequestRepository) GetByReservationID(reservationID string) ([]models.AmenityRequest, error) {
	var list []models.AmenityRequest
	err := r.db.Select(&list, "SELECT * FROM amenity_requests WHERE reservation_id=$1 ORDER BY requested_at DESC", reservationID)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *AmenityRequestRepository) UpdateStatus(id string, status string) error {
	now := time.Now()
	_, err := r.db.Exec("UPDATE amenity_requests SET status=$1, updated_at=$2 WHERE id=$3", status, now, id)
	return err
}

// GetAllReservations returns distinct reservation IDs present in amenity_requests
func (r *AmenityRequestRepository) GetAllReservations() ([]string, error) {
	var ids []string
	err := r.db.Select(&ids, "SELECT DISTINCT reservation_id FROM amenity_requests ORDER BY reservation_id")
	if err != nil {
		return nil, err
	}
	return ids, nil
}

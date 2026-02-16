package models

import "time"

type AmenityType string

const (
	TOWELS       AmenityType = "TOWELS"
	PILLOWS      AmenityType = "PILLOWS"
	TOILETRIES   AmenityType = "TOILETRIES"
	ROOM_SERVICE AmenityType = "ROOM_SERVICE"
	OTHER        AmenityType = "OTHER"
)

type AmenityStatus string

const (
	PENDING     AmenityStatus = "PENDING"
	IN_PROGRESS AmenityStatus = "IN_PROGRESS"
	COMPLETED   AmenityStatus = "COMPLETED"
	CANCELLED   AmenityStatus = "CANCELLED"
)

type AmenityRequest struct {
	ID            string        `db:"id" json:"id"`
	ReservationID string        `db:"reservation_id" json:"reservation_id"`
	GuestID       string        `db:"guest_id" json:"guest_id"`
	PropertyID    string        `db:"property_id" json:"property_id"`
	AmenityType   AmenityType   `db:"amenity_type" json:"amenity_type"`
	Description   string        `db:"description" json:"description"`
	Quantity      int           `db:"quantity" json:"quantity"`
	Status        AmenityStatus `db:"status" json:"status"`
	RequestedAt   time.Time     `db:"requested_at" json:"requested_at"`
	CompletedAt   *time.Time    `db:"completed_at" json:"completed_at"`
	CreatedAt     time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at" json:"updated_at"`
	CreatedBy     string        `db:"created_by" json:"created_by"`
}

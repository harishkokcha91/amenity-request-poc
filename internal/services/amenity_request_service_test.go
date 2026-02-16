package services

import (
	"testing"

	"github.com/example/amenity-poc/internal/models"
)

func TestInvalidQuantity(t *testing.T) {
	s := &AmenityRequestService{}
	req := &models.AmenityRequest{Quantity: 0, AmenityType: models.TOWELS}
	err := s.Create(req)
	if err == nil {
		t.Fatalf("expected error for zero quantity")
	}
}

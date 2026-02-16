-- migration: create amenity_requests table
CREATE TABLE IF NOT EXISTS amenity_requests (
    id TEXT PRIMARY KEY,
    reservation_id TEXT NOT NULL,
    guest_id TEXT NOT NULL,
    property_id TEXT NOT NULL,
    amenity_type TEXT NOT NULL,
    description TEXT,
    quantity INTEGER NOT NULL CHECK (quantity > 0),
    status TEXT NOT NULL,
    requested_at TIMESTAMP WITH TIME ZONE NOT NULL,
    completed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by TEXT
);

CREATE INDEX IF NOT EXISTS idx_amenity_requests_reservation_id ON amenity_requests(reservation_id);
CREATE INDEX IF NOT EXISTS idx_amenity_requests_status ON amenity_requests(status);

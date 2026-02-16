### **POC Task: Amenity Request Feature**

#### **Objective**

The goal of this POC is to evaluate your ability to understand and follow existing backend patterns, write clean and maintainable code, apply validations, and build working APIs using Go.

---

### **Overview**

Build a simple **Amenity Request** feature that allows guests to request additional amenities (such as extra towels, pillows, toiletries, or room service items) for an existing reservation.

---

### **Tech Stack Requirements**

* **Language:** Go
* **Framework:** Gin
* **Database:** PostgreSQL

Please follow a layered architecture similar to a real production codebase:

* Controllers
* Services
* Repositories
* Models

---

### **Why This POC?**

* Tests understanding of standard backend patterns (controllers, services, repositories, models)
* Requires working with existing concepts like reservations and guests
* Demonstrates ability to follow conventions and structure
* Simple enough to complete in **2–4 hours**
* You are **allowed to use AI tools** to assist with development

---

### **Functional Requirements**

#### **1. Model**

Create a new model file:

`internal/models/amenity_request.go`

Fields:

* ID
* ReservationID
* GuestID
* PropertyID
* AmenityType (enum: `TOWELS`, `PILLOWS`, `TOILETRIES`, `ROOM_SERVICE`, `OTHER`)
* Description (text for custom requests)
* Quantity
* Status (`PENDING`, `IN_PROGRESS`, `COMPLETED`, `CANCELLED`)
* RequestedAt
* CompletedAt
* Standard audit fields (`CreatedAt`, `UpdatedAt`, `CreatedBy`)

---

#### **2. Repository**

Create:

`internal/repositories/amenity_request_repository.go`

Methods:

* `Create(request *models.AmenityRequest) error`
* `GetByID(id string) (*models.AmenityRequest, error)`
* `GetByReservationID(reservationID string) ([]models.AmenityRequest, error)`
* `UpdateStatus(id string, status string) error`

---

#### **3. Service**

Create:

`internal/services/amenity_request_service.go`

Responsibilities:

* Business logic for creating and managing amenity requests
* Validate that the reservation exists and is active (checked-in)
* Validate inputs:

  * Quantity must be greater than 0
  * Amenity type must be valid
* Handle status updates cleanly

---

#### **4. Controller**

Create:

`internal/controllers/amenity_request_controller.go`

Endpoints:

* `POST /api/v1/reservations/:id/amenity-requests`
  → Create a new amenity request
* `GET /api/v1/reservations/:id/amenity-requests`
  → List all requests for a reservation
* `PATCH /api/v1/amenity-requests/:requestId/status`
  → Update the status of a request

---

#### **5. Routes**

Register all endpoints in:

`internal/router/router.go`

---

#### **6. Database Migration**

* Create a PostgreSQL migration for the `amenity_requests` table
* Include appropriate indexes and constraints where relevant

---

### **Deliverables**

* All source code files
* PostgreSQL migration file
* A brief **README** explaining:

  * How to run the application
  * How to test the APIs (sample curl/Postman requests)
* *(Optional)* Basic unit test for the service layer

---

### **Evaluation Criteria**

| Criteria                                       | Weight |
| ---------------------------------------------- | ------ |
| Follows existing code patterns and conventions | 25%    |
| Proper error handling                          | 20%    |
| Clean, readable code                           | 20%    |
| Working API endpoints                          | 20%    |
| Proper validation                              | 15%    |

---

### **Time Expectation**

* Estimated effort: **2–4 hours**
* Quality and clarity matter more than over-engineering

Please feel free to reach out if you have any questions or need clarifications before starting.

---
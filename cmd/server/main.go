package main

import (
	"log"
	"os"

	"github.com/example/amenity-poc/internal/db"
	"github.com/example/amenity-poc/internal/router"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:password@localhost:5432/amenitydb?sslmode=disable"
	}

	dbConn, err := db.New(dsn)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer dbConn.Close()

	r := gin.Default()
	router.Setup(r, dbConn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}


package main

import (
	"log"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/buingoctai/book-chapters-summary/book"
	"github.com/buingoctai/book-chapters-summary/internal/rest/middleware"
	"github.com/buingoctai/book-chapters-summary/internal/rest"
	"github.com/joho/godotenv"
)


const (
	defaultAddress = ":8080"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}


func main() {
	e := echo.New()
	e.Use(middleware.CORS)

	// Build service Layer
	svc := book.NewService()
	rest.NewBookHandler(e, svc)

	// Start Server
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address))
}
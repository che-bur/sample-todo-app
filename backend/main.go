package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	// Define a JSON response route
	server.GET("/", func(context echo.Context) error {
		data := map[string]string{"status": "OK"}
		return context.JSON(http.StatusOK, data)
	})

	// Start the server on port 8080
	server.Logger.Fatal(server.Start(":8080"))
}

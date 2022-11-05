package main

import (
	"log"
	"net/http"
	"os"

	echo "github.com/labstack/echo/v4"
)

func main() {
	dat, err := os.ReadFile("/etc/hostname")
	if err != nil {
		return
	}
	dat = dat[:len(dat)-2]

	server := echo.New()
	server.GET("/", handler(string(dat)))

	err = server.Start("0.0.0.0:8800")
	if err != nil {
		log.Fatal(err)
	}
}
func handler(containerID string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, containerID)
	}
}

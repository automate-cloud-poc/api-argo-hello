package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

var GitCommit string

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/argo/hello", func(c echo.Context) error {
		log.Println("my hello log")
		fmt.Println("log backup")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, from argo deploy from commit %s", GitCommit))
	})
	e.GET("/argo/hello/auth", func(c echo.Context) error {
		log.Println("contain auth routes")
		return c.String(http.StatusOK, "hello, did you pass it!?")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
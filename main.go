package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

var GitCommit string

func main() {
	secret1 := os.Getenv("SECRET1")
	configMap1 := os.Getenv("CONFIGMAP1")

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.GET("/argo/hello", func(c echo.Context) error {
		log.Println("my hello log")
		fmt.Println("log backup")
		return c.String(http.StatusOK, fmt.Sprintf("Hello, from argo! commit %s", GitCommit))
	})
	e.GET("/hello/secret", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("this is a secret: %s", secret1))
	})
	e.GET("/hello/configMap", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("this is a config map: %s", configMap1))
	})
	e.GET("/argo/hello/auth", func(c echo.Context) error {
		log.Println("contain auth routes")
		return c.String(http.StatusOK, "hello, did you pass it!?")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
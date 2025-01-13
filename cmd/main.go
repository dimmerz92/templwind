package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dimmerz92/templwind/internal"
	"github.com/dimmerz92/templwind/pkg/lib"
	"github.com/labstack/echo/v4"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8000, "use a port between 1024 and 65535")
	flag.Parse()
	if port < 1024 || port > 65535 {
		log.Fatalf("port: %d out of range, pick between 1024 and 65535", port)
	}
}

func main() {
	e := echo.New()

	e.File("/output.css", "pkg/themes/output.css")

	e.GET("/", func(c echo.Context) error {
		return lib.Render(c, http.StatusOK, internal.Index())
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

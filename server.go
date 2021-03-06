package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	// Read form fields
	name := c.FormValue("name")
    email := c.FormValue("email")
    imgDir := "images/"
	//-----------
	// Read file
	//-----------
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	// Destination
	dst, err := os.Create(imgDir + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
    }
    successMsg := "<p>File %s uploaded successfully with fields name=%s and email=%s.</p>"

	return c.HTML(http.StatusOK, fmt.Sprintf(successMsg, file.Filename, name, email))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":5001"))
}

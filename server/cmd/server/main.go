package main

import "github.com/labstack/echo/v4"

func main() {
	app := echo.New()

	app.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	app.Logger.Fatal(app.Start(":8080"))
}

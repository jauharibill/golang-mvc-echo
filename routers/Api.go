package routers

import (
	"Person/app/Http/Controllers"
	"github.com/labstack/echo"
	"os"
)

func Api() {

	e := echo.New()

	e.GET("/index", Controllers.Index)
	e.GET("/show/:id", Controllers.Show)
	e.POST("/store", Controllers.Store)
	e.PUT("/update/:id", Controllers.Update)
	e.DELETE("/delete/:id", Controllers.Delete)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}

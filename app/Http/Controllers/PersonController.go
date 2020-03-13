package Controllers

import (
	"Person/app/Models"
	DB "Person/config"
	"github.com/labstack/echo"
	"net/http"
)

var persons []Models.PersonModel
var messages map[string]string

func Index(c echo.Context) error {

	rows := DB.Init().Find(&persons)

	return c.JSON(http.StatusOK, rows)
}

func Show(c echo.Context) error {

	id := c.Param("id")

	person := DB.Init().Find(&Models.PersonModel{}, id)

	return c.JSON(http.StatusOK, person)
}

func Store(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")
	full_name := c.FormValue("full_name")
	email := c.FormValue("email")

	person := DB.Init().Create(&Models.PersonModel{
		Username:  username,
		Full_name: full_name,
		Password:  password,
		Email:     email,
	})

	return c.JSON(http.StatusOK, person)
}

func Update(c echo.Context) error {

	person_id := c.Param("id")

	username := c.FormValue("username")
	full_name := c.FormValue("full_name")
	password := c.FormValue("password")
	email := c.FormValue("email")

	person := DB.Init().First(&Models.PersonModel{}, person_id).Update(Models.PersonModel{
		Username:  username,
		Full_name: full_name,
		Password:  password,
		Email:     email,
	})

	return c.JSON(http.StatusOK, person)
}

func Delete(c echo.Context) error {

	person_id := c.Param("id")

	person := DB.Init().Delete(&Models.PersonModel{}, person_id)

	return c.JSON(http.StatusOK, person)
}
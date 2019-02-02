package controllers

import (
  "net/http"
  // "fmt"
	"crud/models"
  "github.com/labstack/echo"
  "strconv"
)

func GetUser(c echo.Context) error{
  id, _ := strconv.Atoi(c.Param("id"))
  result := models.GetUser(id)
  // fmt.Printf("id: %d\nname: %s\n", result.Id, result.Name)
  return c.Render(http.StatusOK, "user.html", result)
}

func GetUsers(c echo.Context) error{
  result := models.GetUsers()
  // fmt.Printf("%+v\n",result)
  // return c.JSON(http.StatusOK, result)
  return c.Render(http.StatusOK, "users.html", result)
}

func Insert(c echo.Context) error{
  // id, _ := strconv.Atoi(c.FormValue("id"))
  name  := c.FormValue("name")
	email := c.FormValue("email")
  models.Insert(name, email)
  return GetUsers(c)
}

func Update(c echo.Context) error{
  id, _ := strconv.Atoi(c.FormValue("id"))
  name  := c.FormValue("name")
	email := c.FormValue("email")
  models.Update(id, name, email)
  return GetUsers(c)
}

func Delete(c echo.Context) error{
  id, _ := strconv.Atoi(c.Param("id"))
  models.Delete(id)
  return GetUsers(c)
}

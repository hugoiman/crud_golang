package main

import (
  "fmt"
  "html/template"
  "io"
  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
  "crud/controllers"
)

type Template struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
  e := echo.New()

  e.Static("/","assets")

  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"*"},
    AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
  }))

  renderer := &Template{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }

  e.Renderer = renderer

  e.GET("/user/:id", controllers.GetUser)
  e.GET("/users", controllers.GetUsers)
  e.POST("/insert", controllers.Insert)
  e.POST("/update", controllers.Update)
  e.DELETE("/delete/:id", controllers.Delete)

  fmt.Println("server started at :9000")
  e.Logger.Fatal(e.Start(":9000"))
}

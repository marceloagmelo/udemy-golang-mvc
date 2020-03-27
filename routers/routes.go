package routers

import (
	"github.com/labstack/echo"
	"github.com/marceloagmelo/udemy-golang-mvc/controllers"
)

//App é uma instância de echo
var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)
	App.GET("/atualizar/:id", controllers.Atualizar)

	api := App.Group("/v1")
	api.POST("/insert", controllers.Inserir)
	api.DELETE("/delete/:id", controllers.Deletar)
	api.PUT("/update/:id", controllers.Update)
}

package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-quickstart/web/controllers"
)

func main() {
	/**

	 */
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// Register controller
	mvc.New(app.Party("/hello")).Handle(new(controllers.MoveController))
	app.Run(
		iris.Addr("localhost:8080"),
		)
}

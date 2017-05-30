package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/typescript"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
	Paths()
}

// Paths returns paths
func Paths() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New()) // adapt a router, order doesn't matters but before Listen.
	tmpl := view.HTML("./frontend/react_renderings", ".html")

	app.Adapt(tmpl)
	ts := typescript.New()
	ts.Config.Dir = "./frontend/scripts/typescript"
	app.Adapt(ts) // adapt the typescript compiler adaptor

	app.StaticWeb("/static", "./frontend/react_renderings")
	app.StaticWeb("/public", "./frontend/scripts/typescript")

	app.Get("/", hi)

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.HTML(iris.StatusNotFound, "<h1>Custom not found handler </h1>")
	})

	app.Listen(":9000")
}

func hi(ctx *iris.Context) {
	ctx.Render("index.html", nil)
}

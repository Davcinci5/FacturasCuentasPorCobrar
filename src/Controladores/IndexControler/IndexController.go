package Index

import "gopkg.in/kataras/iris.v6"

//IndexGet renderiza pagina principal
func IndexGet(ctx *iris.Context) {
	ctx.Render("wizard.html", nil)
}

//IndexPost renderiza pagina principal
func IndexPost(ctx *iris.Context) {
	ctx.Render("wizard.html", nil)
}

package Index

import (
	"gopkg.in/kataras/iris.v6"

	"../../Modelos/CuentasPorCobrarModel"
)

//IndexGet renderiza pagina principal
func IndexGet(ctx *iris.Context) {
	var Send CuentasPorCobrarModel.SCuentasPorCobrar
	ctx.Render("CuentasPorCobrarIndex.html", Send)
	//ctx.Render("wizard.html", nil)
}

//IndexPost renderiza pagina principal
func IndexPost(ctx *iris.Context) {
	var Send CuentasPorCobrarModel.SCuentasPorCobrar
	ctx.Render("CuentasPorCobrarIndex.html", Send)
	//ctx.Render("wizard.html", nil)
}

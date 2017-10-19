package DatosFacturaControler

import (
	"fmt"

	"gopkg.in/kataras/iris.v6"
)

//IndexGet renderiza pagina principal
func IndexGet(ctx *iris.Context) {
	NUMCFD := ctx.Param("ID")
	fmt.Println(NUMCFD)
	ctx.Render("wizard.html", iris.Map{"NUMCFD": NUMCFD})
}

//IndexPost renderiza pagina principal
func IndexPost(ctx *iris.Context) {
	ctx.Render("wizard.html", nil)
}

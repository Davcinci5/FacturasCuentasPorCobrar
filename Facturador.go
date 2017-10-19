package main

import (
	"fmt"

	"./src/Controladores/CatalogoControler"
	"./src/Controladores/CuentasPorCobrarControler"
	"./src/Controladores/DatosFacturaControler"
	"./src/Controladores/DetalleCuentasPorCobrarVisorusControler"
	"./src/Controladores/EmpresaControler"
	"./src/Controladores/HTTPErrorsControler"
	"./src/Controladores/IndexControler"
	"./src/Modulos/Variables"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
	//###################### Start ####################################
	app := iris.New()
	app.Adapt(httprouter.New())
	sesiones := sessions.New(sessions.Config{Cookie: "cookiekore"})
	app.Adapt(sesiones)

	app.Adapt(view.HTML("./Public/Vistas", ".html").Reload(true))

	app.Set(iris.OptionCharset("UTF-8"))

	app.StaticWeb("/icono", "./Public/Recursos/Generales/img")

	app.StaticWeb("/css", "./Public/Recursos/Generales/css")
	app.StaticWeb("/js", "./Public/Recursos/Generales/js")
	app.StaticWeb("/Plugins", "./Public/Recursos/Generales/Plugins")
	app.StaticWeb("/scripts", "./Public/Recursos/Generales/scripts")
	app.StaticWeb("/img", "./Public/Recursos/Generales/img")
	app.StaticWeb("/Comprobantes", "./Public/Recursos/Locales/comprobantes")
	app.StaticWeb("/Locales", "./Public/Recursos/Locales")

	//###################### CFG ######################################

	var DataCfg = MoVar.CargaSeccionCFG(MoVar.SecDefault)
	//###################### Ruteo ####################################

	//###################### HTTP ERRORS ################################

	//Error 403
	app.OnError(iris.StatusForbidden, HTTPErrors.Error403)
	//Error 404
	app.OnError(iris.StatusNotFound, HTTPErrors.Error404)

	//Error 500
	app.OnError(iris.StatusInternalServerError, HTTPErrors.Error500)

	// PRUEBAS DE ERRORES
	app.Get("/500", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusInternalServerError)
	})
	app.Get("/404", func(ctx *iris.Context) {
		ctx.EmitError(iris.StatusNotFound)
	})
	//Index
	app.Get("/", Index.IndexGet)
	app.Post("/", Index.IndexGet)
	//###################### Catalogo ################################
	//Index (Búsqueda)
	app.Get("/Catalogos", CatalogoControler.IndexGet)
	app.Post("/Catalogos", CatalogoControler.IndexPost)
	app.Post("/Catalogos/search", CatalogoControler.BuscaPagina)
	app.Post("/Catalogos/agrupa", CatalogoControler.MuestraIndexPorGrupo)

	//Alta
	app.Get("/Catalogos/alta", CatalogoControler.AltaGet)
	app.Post("/Catalogos/alta", CatalogoControler.AltaPost)

	//Edicion
	app.Get("/Catalogos/edita", CatalogoControler.EditaGet)
	app.Post("/Catalogos/edita", CatalogoControler.EditaPost)
	app.Get("/Catalogos/edita/:ID", CatalogoControler.EditaGet)
	app.Post("/Catalogos/edita/:ID", CatalogoControler.EditaPost)

	//Detalle
	app.Get("/Catalogos/detalle", CatalogoControler.DetalleGet)
	app.Post("/Catalogos/detalle", CatalogoControler.DetallePost)
	app.Get("/Catalogos/detalle/:ID", CatalogoControler.DetalleGet)
	app.Post("/Catalogos/detalle/:ID", CatalogoControler.DetallePost)

	//Rutinas adicionales

	//###################### Empresa ################################
	//Index (Búsqueda)

	app.Get("/Empresas", EmpresaControler.EditaGet)
	app.Post("/Empresas", EmpresaControler.EditaPost)

	//Alta
	app.Get("/Empresas/alta", EmpresaControler.AltaGet)
	app.Post("/Empresas/alta", EmpresaControler.AltaPost)

	//Edicion
	app.Get("/Empresas/edita", EmpresaControler.EditaGet)
	app.Post("/Empresas/edita", EmpresaControler.EditaPost)
	app.Get("/Empresas/edita/:ID", EmpresaControler.EditaGet)
	app.Post("/Empresas/edita/:ID", EmpresaControler.EditaPost)

	//Detalle
	app.Get("/Empresas/detalle", EmpresaControler.DetalleGet)
	app.Post("/Empresas/detalle", EmpresaControler.DetallePost)
	app.Get("/Empresas/detalle/:ID", EmpresaControler.DetalleGet)
	app.Post("/Empresas/detalle/:ID", EmpresaControler.DetallePost)

	//Rutinas adicionales

	//###################### CuentasPorCobrar ################################
	//Index (Búsqueda)
	app.Get("/CuentasPorCobrars", CuentasPorCobrarControler.IndexGet)
	app.Post("/CuentasPorCobrar/CargaIndex", CuentasPorCobrarControler.CargaIndex)
	app.Post("/CuentasPorCobrar/CargaIndexEspecifico", CuentasPorCobrarControler.CargaIndexEspecifico)
	//	app.Post("/CuentasPorCobrars", CuentasPorCobrarControler.IndexPost)
	//	app.Post("/CuentasPorCobrars/search", CuentasPorCobrarControler.BuscaPagina)
	//	app.Post("/CuentasPorCobrars/agrupa", CuentasPorCobrarControler.MuestraIndexPorGrupo)
	//Detalle
	app.Get("/CuentasPorCobrars/detalle", CuentasPorCobrarControler.DetalleGet)
	app.Post("/CuentasPorCobrars/detalle", CuentasPorCobrarControler.DetallePost)
	app.Get("/CuentasPorCobrars/detalle/:ID", CuentasPorCobrarControler.DetalleGet)
	app.Post("/CuentasPorCobrars/detalle/:ID", CuentasPorCobrarControler.DetallePost)

	//###################### DetalleCuentasPorCobrarVisorus ################################
	//Index (Búsqueda)
	app.Get("/DetalleCuentasPorCobrarVisoruss", DetalleCuentasPorCobrarVisorusControler.IndexGet)
	app.Post("/DetalleCuentasPorCobrarVisorus/CargaIndex/:ID", DetalleCuentasPorCobrarVisorusControler.CargaIndex)
	app.Post("/DetalleCuentasPorCobrarVisorus/CargaIndexEspecifico", DetalleCuentasPorCobrarVisorusControler.CargaIndexEspecifico)
	//app.Post("/DetalleCuentasPorCobrarVisoruss", DetalleCuentasPorCobrarVisorusControler.IndexPost)
	//app.Post("/DetalleCuentasPorCobrarVisoruss/search", DetalleCuentasPorCobrarVisorusControler.BuscaPagina)
	//app.Post("/DetalleCuentasPorCobrarVisoruss/agrupa", DetalleCuentasPorCobrarVisorusControler.MuestraIndexPorGrupo)

	//Detalle
	app.Get("/DetalleCuentasPorCobrarVisoruss/detalle", DetalleCuentasPorCobrarVisorusControler.DetalleGet)
	app.Post("/DetalleCuentasPorCobrarVisoruss/detalle", DetalleCuentasPorCobrarVisorusControler.DetallePost)
	app.Get("/DetalleCuentasPorCobrarVisoruss/detalle/:ID", DetalleCuentasPorCobrarVisorusControler.DetalleGet)
	app.Post("/DetalleCuentasPorCobrarVisoruss/detalle/:ID", DetalleCuentasPorCobrarVisorusControler.DetallePost)

	//###################### DatosFactura #####################################

	app.Get("/DatosFactura/alta/:ID", DatosFacturaControler.IndexGet)
	app.Post("/DatosFactura/alta/:ID", DatosFacturaControler.IndexPost)
	//                        ...

	//###################### Listen Server #############################

	//###################### Problemas de Sistema #############################
	if DataCfg.Puerto != "" {
		fmt.Println("Ejecutandose en el puerto: ", DataCfg.Puerto)
		fmt.Println("Acceder a la siguiente url: ", DataCfg.BaseURL)
		app.Listen(":" + DataCfg.Puerto)
	} else {
		fmt.Println("Ejecutandose en el puerto: 8080")
		fmt.Println("Acceder a la siguiente url: localhost")
		app.Listen(":8080")
	}

}

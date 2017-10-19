package DetalleCuentasPorCobrarVisorusControler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	//"net/http"

	"../../Modulos/Session"
	"github.com/tealeg/xlsx"

	"../../Modelos/DetalleCuentasPorCobrarVisorusModel"
	//"../../Modulos/CargaCombos"
	//"../../Modulos/Conexiones"
	"../../Modulos/General"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/mgo.v2/bson"
)

//##########< Variables Generales > ############

var cadenaBusqueda string
var numeroRegistros int
var paginasTotales int

//NumPagina especifica el numero de página en la que se cargarán los registros
var NumPagina int

//limitePorPagina limite de registros a mostrar en la pagina
var limitePorPagina = 10

//IDElastic id obtenido de Elastic
var IDElastic bson.ObjectId
var arrIDMgo []bson.ObjectId
var arrIDElastic []bson.ObjectId
var arrToMongo []bson.ObjectId

//####################< INDEX (BUSQUEDA) >###########################

//IndexGet renderea al index de DetalleCuentasPorCobrarVisorus
func IndexGet(ctx *iris.Context) {

	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus
	ctx.Render("DetalleCuentasPorCobrarVisorusIndex.html", Send)

}

//CargaIndex regresa los datos necesarios para cargar el Index de cada modelo
func CargaIndex(ctx *iris.Context) {
	id := ctx.Param("ID")
	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	Send.SIndex.STituloTabla = "SE CARGARON TODAS LAS UNIDADES"
	Modelo := DetalleCuentasPorCobrarVisorusModel.MapaModelo()
	Send.SIndex.SNombresDeColumnas = Modelo["name"].([]string)
	Send.SIndex.SModeloDeColumnas = MoGeneral.GeneraModeloColumnas(Modelo)
	Send.SIndex.SRenglones = DetalleCuentasPorCobrarVisorusModel.GeneraDataRowDeIndex(DetalleCuentasPorCobrarVisorusModel.GetOne(id))
	Send.SEstado = true
	jData, err := json.Marshal(Send)
	if err != nil {
		fmt.Println(err)
	}
	ctx.Header().Set("Content-Type", "application/json")
	ctx.Write(jData)
	return
}

//CargaIndexEspecifico regresa los datos necesarios para cargar el Index de cada modelo
func CargaIndexEspecifico(ctx *iris.Context) {
	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus
	var TextoABuscar string
	TextoABuscar = ctx.FormValue("Texto")
	fmt.Println("esto es lo que tiene Texto a Buscar: ", TextoABuscar)

	if TextoABuscar != "" {

		Send.SIndex.STituloTabla = "SE CARGARON LAS UNIDADES FILTRADAS POR: " + TextoABuscar
		Modelo := DetalleCuentasPorCobrarVisorusModel.MapaModelo()
		Send.SIndex.SNombresDeColumnas = Modelo["name"].([]string)
		Send.SIndex.SModeloDeColumnas = MoGeneral.GeneraModeloColumnas(Modelo)

		//Consultamos en Elastic y obtenemos los datos para mostrar en la nueva colección

		docs := DetalleCuentasPorCobrarVisorusModel.BuscarEnElastic(TextoABuscar)

		if docs.Hits.TotalHits > 0 {

			arrIDElastic = []bson.ObjectId{}
			for _, item := range docs.Hits.Hits {
				IDElastic = bson.ObjectIdHex(item.Id)
				arrIDElastic = append(arrIDElastic, IDElastic)
			}

		} else {

			Send.SEstado = false
			Send.SMsj += "No se encontraron resultados para: " + TextoABuscar

			jData, err := json.Marshal(Send)
			if err != nil {
				fmt.Println(err)
			}

			ctx.Header().Set("Content-Type", "application/json")
			ctx.Write(jData)
			return
		}

		fmt.Println("Se encontraron ", len(arrIDElastic), " registros.")
		Send.SIndex.SRenglones = DetalleCuentasPorCobrarVisorusModel.GeneraDataRowDeIndex(DetalleCuentasPorCobrarVisorusModel.GetEspecifics(arrIDElastic))
		fmt.Println(Send.SIndex.SRenglones)
		Send.SEstado = true
		jData, err := json.Marshal(Send)
		if err != nil {
			fmt.Println(err)
		}
		ctx.Header().Set("Content-Type", "application/json")
		ctx.Write(jData)
		return
	}

	jData, err := json.Marshal(Send)
	if err != nil {
		fmt.Println(err)
	}

	Send.SEstado = false
	Send.SMsj += " Se recibió una cadena de texto vacía, ingrese un texto a buscar e intente de nuevo."
	ctx.Header().Set("Content-Type", "application/json")
	ctx.Write(jData)
	return

}

func UploadFileGet(ctx *iris.Context) {

	ctx.Render("DetalleCuentasPorCobrarVisorusUploadFile.html", nil)
}

func UploadFilePost(ctx *iris.Context) {

	file, _, err := ctx.FormFile("file")

	nombrefile := "CatalogosCFDI.xlsx"
	dirpathTemp := "./Public/Recursos/Temporales"

	if _, err := os.Stat(dirpathTemp); os.IsNotExist(err) {
		fmt.Println("el directorio no	 existe")
		os.MkdirAll(dirpathTemp, 0777)
	}

	//subir imagen al servidor local
	out, err := os.Create(dirpathTemp + "/" + nombrefile)
	if err != nil {
		fmt.Println("No es posible crear el archivo en el directorio, compruebe los permisos", err)
		//return "No es posible crear el archivo en el directorio, compruebe los permisos", err
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		//return "Error al escribir la imagen al directorio", err
		fmt.Println("Error al escribir la imagen al directorio", err)
	}

	excelFileName := "./Public/Recursos/Temporales/CatalogosCFDI.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}

	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus
	var NombresSheets []string

	for _, sheet := range xlFile.Sheets {
		NombresSheets = append(NombresSheets, sheet.Name)
		fmt.Println(sheet.Name)
	}

	//fmt.Println("FILE: ", err, nombrefile)
	ctx.Render("DetalleCuentasPorCobrarVisorusSelectSheet.html", Send)
}

//###########################< ALTA >################################

//AltaGet renderea al alta de DetalleCuentasPorCobrarVisorus
func AltaGet(ctx *iris.Context) {

	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	NameUsrLoged, MenuPrincipal, MenuUsr, errSes := Session.GetDataSession(ctx) //Retorna los datos de la session
	Send.SSesion.Name = NameUsrLoged
	Send.SSesion.MenuPrincipal = template.HTML(MenuPrincipal)
	Send.SSesion.MenuUsr = template.HTML(MenuUsr)
	if errSes != nil {
		Send.SEstado = false
		Send.SMsj = errSes.Error()
		ctx.Render("ZError.html", Send)
		return
	}

	//####   TÚ CÓDIGO PARA CARGAR DATOS A LA VISTA DE ALTA----> PROGRAMADOR

	ctx.Render("DetalleCuentasPorCobrarVisorusAlta.html", Send)

}

//AltaPost regresa la petición post que se hizo desde el alta de DetalleCuentasPorCobrarVisorus
func AltaPost(ctx *iris.Context) {

	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	NameUsrLoged, MenuPrincipal, MenuUsr, errSes := Session.GetDataSession(ctx) //Retorna los datos de la session
	Send.SSesion.Name = NameUsrLoged
	Send.SSesion.MenuPrincipal = template.HTML(MenuPrincipal)
	Send.SSesion.MenuUsr = template.HTML(MenuUsr)
	if errSes != nil {
		Send.SEstado = false
		Send.SMsj = errSes.Error()
		ctx.Render("ZError.html", Send)
		return
	}

	//####   TÚ CÓDIGO PARA PROCESAR DATOS DE LA VISTA DE ALTA Y GUARDARLOS O REGRESARLOS----> PROGRAMADOR

	ctx.Render("DetalleCuentasPorCobrarVisorusAlta.html", Send)

}

//###########################< EDICION >###############################

//EditaGet renderea a la edición de DetalleCuentasPorCobrarVisorus
func EditaGet(ctx *iris.Context) {

	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	NameUsrLoged, MenuPrincipal, MenuUsr, errSes := Session.GetDataSession(ctx) //Retorna los datos de la session
	Send.SSesion.Name = NameUsrLoged
	Send.SSesion.MenuPrincipal = template.HTML(MenuPrincipal)
	Send.SSesion.MenuUsr = template.HTML(MenuUsr)
	if errSes != nil {
		Send.SEstado = false
		Send.SMsj = errSes.Error()
		ctx.Render("ZError.html", Send)
		return
	}

	ctx.Render("DetalleCuentasPorCobrarVisorusEdita.html", Send)

}

//EditaPost regresa el resultado de la petición post generada desde la edición de DetalleCuentasPorCobrarVisorus
func EditaPost(ctx *iris.Context) {

	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	NameUsrLoged, MenuPrincipal, MenuUsr, errSes := Session.GetDataSession(ctx) //Retorna los datos de la session
	Send.SSesion.Name = NameUsrLoged
	Send.SSesion.MenuPrincipal = template.HTML(MenuPrincipal)
	Send.SSesion.MenuUsr = template.HTML(MenuUsr)
	if errSes != nil {
		Send.SEstado = false
		Send.SMsj = errSes.Error()
		ctx.Render("ZError.html", Send)
		return
	}

	//####   TÚ CÓDIGO PARA PROCESAR DATOS DE LA VISTA DE ALTA Y GUARDARLOS O REGRESARLOS----> PROGRAMADOR

	ctx.Render("DetalleCuentasPorCobrarVisorusEdita.html", Send)

}

//#################< DETALLE >####################################

//DetalleGet renderea al index.html
func DetalleGet(ctx *iris.Context) {
	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	NameUsrLoged, MenuPrincipal, MenuUsr, errSes := Session.GetDataSession(ctx) //Retorna los datos de la session
	Send.SSesion.Name = NameUsrLoged
	Send.SSesion.MenuPrincipal = template.HTML(MenuPrincipal)
	Send.SSesion.MenuUsr = template.HTML(MenuUsr)
	if errSes != nil {
		Send.SEstado = false
		Send.SMsj = errSes.Error()
		ctx.Render("ZError.html", Send)
		return
	}

	//###### TU CÓDIGO AQUÍ PROGRAMADOR

	ctx.Render("DetalleCuentasPorCobrarVisorusDetalle.html", Send)
}

//DetallePost renderea al index.html
func DetallePost(ctx *iris.Context) {
	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	NameUsrLoged, MenuPrincipal, MenuUsr, errSes := Session.GetDataSession(ctx) //Retorna los datos de la session
	Send.SSesion.Name = NameUsrLoged
	Send.SSesion.MenuPrincipal = template.HTML(MenuPrincipal)
	Send.SSesion.MenuUsr = template.HTML(MenuUsr)
	if errSes != nil {
		Send.SEstado = false
		Send.SMsj = errSes.Error()
		ctx.Render("ZError.html", Send)
		return
	}

	//###### TU CÓDIGO AQUÍ PROGRAMADOR

	ctx.Render("DetalleCuentasPorCobrarVisorusDetalle.html", Send)
}

/*
//####################< RUTINAS ADICIONALES >##########################

//BuscaPagina regresa la tabla de busqueda y su paginacion en el momento de especificar página
func BuscaPagina(ctx *iris.Context) {
	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus

	Pagina := ctx.FormValue("Pag")
	if Pagina != "" {
		num, _ := strconv.Atoi(Pagina)
		if num == 0 {
			num = 1
		}
		NumPagina = num
		skip := limitePorPagina * (NumPagina - 1)
		limite := skip + limitePorPagina

		arrToMongo = []bson.ObjectId{}
		if NumPagina == paginasTotales {
			final := int(numeroRegistros) % limitePorPagina
			if final == 0 {
				for _, v := range arrIDElastic[skip:limite] {
					arrToMongo = append(arrToMongo, v)
				}
			} else {
				for _, v := range arrIDElastic[skip : skip+final] {
					arrToMongo = append(arrToMongo, v)
				}
			}

		} else {
			for _, v := range arrIDElastic[skip:limite] {
				arrToMongo = append(arrToMongo, v)
			}
		}

		Cabecera, Cuerpo := DetalleCuentasPorCobrarVisorusModel.GeneraTemplatesBusqueda(DetalleCuentasPorCobrarVisorusModel.GetEspecifics(arrToMongo))
		Send.SIndex.SCabecera = template.HTML(Cabecera)
		Send.SIndex.SBody = template.HTML(Cuerpo)

		Paginacion := MoGeneral.ConstruirPaginacion(paginasTotales, NumPagina)
		Send.SIndex.SPaginacion = template.HTML(Paginacion)

	} else {
		Send.SMsj = "No se recibió una cadena de consulta, favor de escribirla."

	}

	Send.SIndex.SGrupo = template.HTML(CargaCombos.CargaComboMostrarEnIndex(limitePorPagina))
	Send.SEstado = true

	jData, _ := json.Marshal(Send)
	ctx.Header().Set("Content-Type", "application/json")
	ctx.Write(jData)
	return
}

//MuestraIndexPorGrupo regresa template de busqueda y paginacion de acuerdo a la agrupacion solicitada
func MuestraIndexPorGrupo(ctx *iris.Context) {
	var Send DetalleCuentasPorCobrarVisorusModel.SDetalleCuentasPorCobrarVisorus
	var Cabecera, Cuerpo string

	grupo := ctx.FormValue("Grupox")
	if grupo != "" {
		gru, _ := strconv.Atoi(grupo)
		limitePorPagina = gru
	}

	cadenaBusqueda = ctx.FormValue("searchbox")
	//Send.DetalleCuentasPorCobrarVisorus.ENombreDetalleCuentasPorCobrarVisorus.Nombre = cadenaBusqueda

	if cadenaBusqueda != "" {

		docs := DetalleCuentasPorCobrarVisorusModel.BuscarEnElastic(cadenaBusqueda)

		if docs.Hits.TotalHits > 0 {
			arrIDElastic = []bson.ObjectId{}

			for _, item := range docs.Hits.Hits {
				IDElastic = bson.ObjectIdHex(item.Id)
				arrIDElastic = append(arrIDElastic, IDElastic)
			}

			numeroRegistros = len(arrIDElastic)

			arrToMongo = []bson.ObjectId{}
			if numeroRegistros <= limitePorPagina {
				for _, v := range arrIDElastic[0:numeroRegistros] {
					arrToMongo = append(arrToMongo, v)
				}
			} else if numeroRegistros >= limitePorPagina {
				for _, v := range arrIDElastic[0:limitePorPagina] {
					arrToMongo = append(arrToMongo, v)
				}
			}

			Cabecera, Cuerpo = DetalleCuentasPorCobrarVisorusModel.GeneraTemplatesBusqueda(DetalleCuentasPorCobrarVisorusModel.GetEspecifics(arrToMongo))
			Send.SIndex.SCabecera = template.HTML(Cabecera)
			Send.SIndex.SBody = template.HTML(Cuerpo)
			MoConexion.FlushElastic()

			paginasTotales = MoGeneral.Totalpaginas(numeroRegistros, limitePorPagina)
			Paginacion := MoGeneral.ConstruirPaginacion(paginasTotales, 1)
			Send.SIndex.SPaginacion = template.HTML(Paginacion)

		} else {

			if numeroRegistros <= limitePorPagina {
				Cabecera, Cuerpo = DetalleCuentasPorCobrarVisorusModel.GeneraTemplatesBusqueda(DetalleCuentasPorCobrarVisorusModel.GetEspecifics(arrIDMgo[0:numeroRegistros]))
			} else if numeroRegistros >= limitePorPagina {
				Cabecera, Cuerpo = DetalleCuentasPorCobrarVisorusModel.GeneraTemplatesBusqueda(DetalleCuentasPorCobrarVisorusModel.GetEspecifics(arrIDMgo[0:limitePorPagina]))
			}

			Send.SIndex.SCabecera = template.HTML(Cabecera)
			Send.SIndex.SBody = template.HTML(Cuerpo)

			paginasTotales = MoGeneral.Totalpaginas(numeroRegistros, limitePorPagina)
			Paginacion := MoGeneral.ConstruirPaginacion(paginasTotales, 1)
			Send.SIndex.SPaginacion = template.HTML(Paginacion)

			Send.SIndex.SRMsj = "No se encontraron resultados para: " + cadenaBusqueda + " ."
		}

	} else {

		if numeroRegistros <= limitePorPagina {
			Cabecera, Cuerpo = DetalleCuentasPorCobrarVisorusModel.GeneraTemplatesBusqueda(DetalleCuentasPorCobrarVisorusModel.GetEspecifics(arrIDMgo[0:numeroRegistros]))
		} else if numeroRegistros >= limitePorPagina {
			Cabecera, Cuerpo = DetalleCuentasPorCobrarVisorusModel.GeneraTemplatesBusqueda(DetalleCuentasPorCobrarVisorusModel.GetEspecifics(arrIDMgo[0:limitePorPagina]))
		}

		Send.SIndex.SCabecera = template.HTML(Cabecera)
		Send.SIndex.SBody = template.HTML(Cuerpo)

		paginasTotales = MoGeneral.Totalpaginas(numeroRegistros, limitePorPagina)
		Paginacion := MoGeneral.ConstruirPaginacion(paginasTotales, 1)
		Send.SIndex.SPaginacion = template.HTML(Paginacion)

		Send.SIndex.SRMsj = "No se encontraron resultados para: " + cadenaBusqueda + " ."
	}
	Send.SIndex.SGrupo = template.HTML(CargaCombos.CargaComboMostrarEnIndex(limitePorPagina))
	Send.SEstado = true

	jData, _ := json.Marshal(Send)
	ctx.Header().Set("Content-Type", "application/json")
	ctx.Write(jData)
	return
}
*/

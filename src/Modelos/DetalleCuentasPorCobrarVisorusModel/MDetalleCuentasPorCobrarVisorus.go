package DetalleCuentasPorCobrarVisorusModel

import (
	"fmt"
	"strconv"

	//"../../Modelos/CatalogoModel"

	"../../Modulos/Conexiones"
	"../../Modulos/General"

	"../../Modulos/Variables"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/olivere/elastic.v5"
)

//#########################< ESTRUCTURAS >##############################

//DetalleCuentasPorCobrarVisorusMgo estructura de DetalleCuentasPorCobrarVisoruss mongo
type DetalleCuentasPorCobrarVisorusMgo struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	NUMCFD      int64         `bson:"NUMCFD_1"`
	CODIGO      string        `bson:"CODIGO"`
	DESCRIPCION string        `bson:"DESCRIPCION"`
	UNIDAD      string        `bson:"UNIDAD"`
	CANTIDAD    float64       `bson:"CANTIDAD"`
	PRECIO      float64       `bson:"PRECIO"`
	PRECIOIVAIN float64       `bson:"PRECIOIVAIN"`
	DESCUENTO   float64       `bson:"DESCUENTO"`
	IMPORTE     float64       `bson:"IMPORTE"`
	TASA        string        `bson:"TASA"`
	IEPS        string        `bson:"IEPS"`
	VIEPS       float64       `bson:"VIEPS"`
	NUMMOVCFD   int64         `bson:"NUMMOVCFD"`
	NUMDTMCFD   int64         `bson:"NUMDTMCFD"`
	CANTIDADM   float64       `bson:"CANTIDADM"`
	COSTO       float64       `bson:"COSTO"`
}

//DetalleCuentasPorCobrarVisorusElastic estructura de DetalleCuentasPorCobrarVisoruss para insertar en Elastic
type DetalleCuentasPorCobrarVisorusElastic struct {
	NUMCFD      int64   `json:"NUMCFD"`
	CODIGO      string  `json:"CODIGO"`
	DESCRIPCION string  `json:"DESCRIPCION"`
	UNIDAD      string  `json:"UNIDAD"`
	CANTIDAD    float64 `json:"CANTIDAD"`
	PRECIO      float64 `json:"PRECIO"`
	PRECIOIVAIN float64 `json:"PRECIOIVAIN"`
	DESCUENTO   float64 `json:"DESCUENTO"`
	IMPORTE     float64 `json:"IMPORTE"`
	TASA        string  `json:"TASA"`
	IEPS        string  `json:"IEPS"`
	VIEPS       float64 `json:"VIEPS"`
	NUMMOVCFD   int64   `json:"NUMMOVCFD"`
	NUMDTMCFD   int64   `json:"NUMDTMCFD"`
	CANTIDADM   float64 `json:"CANTIDADM"`
	COSTO       float64 `json:"COSTO"`
}

//#########################< FUNCIONES GENERALES MGO >###############################

//GetAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func GetOne(id string) []DetalleCuentasPorCobrarVisorusMgo {
	ID, _ := strconv.Atoi(id)
	var result []DetalleCuentasPorCobrarVisorusMgo
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}

	err = DetalleCuentasPorCobrarVisoruss.Find(bson.M{"NUMCFD": ID}).All(&result) //agregar campo a buscar
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func GetAll() []DetalleCuentasPorCobrarVisorusMgo {
	var result []DetalleCuentasPorCobrarVisorusMgo
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}

	err = DetalleCuentasPorCobrarVisoruss.Find(nil).All(&result) //agregar campo a buscar
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//CountAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func CountAll() int {
	var result int
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)

	if err != nil {
		fmt.Println(err)
	}
	result, err = DetalleCuentasPorCobrarVisoruss.Find(nil).Count()
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

/*/GetOne Regresa un documento específico de Mongo (Por Coleccion) <--------original
func GetOne(ID bson.ObjectId) DetalleCuentasPorCobrarVisorusMgo {
	var result DetalleCuentasPorCobrarVisorusMgo
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}
	err = DetalleCuentasPorCobrarVisoruss.Find(bson.M{"_id": ID}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}*/

//GetEspecifics rsegresa un conjunto de documentos específicos de Mongo (Por Coleccion)
func GetEspecifics(Ides []bson.ObjectId) []DetalleCuentasPorCobrarVisorusMgo {
	var result []DetalleCuentasPorCobrarVisorusMgo
	var aux DetalleCuentasPorCobrarVisorusMgo
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range Ides {
		aux = DetalleCuentasPorCobrarVisorusMgo{}
		DetalleCuentasPorCobrarVisoruss.Find(bson.M{"_id": value}).One(&aux)
		result = append(result, aux)
	}
	s.Close()
	return result
}

//GetEspecificByFields regresa un documento de Mongo especificando un campo y un determinado valor
func GetEspecificByFields(field string, valor interface{}) DetalleCuentasPorCobrarVisorusMgo {
	var result DetalleCuentasPorCobrarVisorusMgo
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)

	if err != nil {
		fmt.Println(err)
	}
	err = DetalleCuentasPorCobrarVisoruss.Find(bson.M{field: valor}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetIDByField regresa un documento específico de Mongo (Por Coleccion)
func GetIDByField(field string, valor interface{}) bson.ObjectId {
	var result DetalleCuentasPorCobrarVisorusMgo
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}
	err = DetalleCuentasPorCobrarVisoruss.Find(bson.M{field: valor}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result.ID
}

//CargaComboDetalleCuentasPorCobrarVisoruss regresa un combo de DetalleCuentasPorCobrarVisorus de mongo
func CargaComboDetalleCuentasPorCobrarVisoruss(ID string) string {
	DetalleCuentasPorCobrarVisoruss := GetAll()

	templ := ``

	if ID != "" {
		templ = `<option value="">--SELECCIONE--</option> `
	} else {
		templ = `<option value="" selected>--SELECCIONE--</option> `
	}

	for _, v := range DetalleCuentasPorCobrarVisoruss {
		if ID == v.ID.Hex() {
			templ += `<option value="` + v.ID.Hex() + `" selected>  ` + v.CODIGO + ` </option> `
		} else {
			templ += `<option value="` + v.ID.Hex() + `">  ` + v.CODIGO + ` </option> `
		}

	}
	return templ
}

//GeneraRenglonesIndex crea un mapa para crear el json de renglones del index
func GeneraRenglonesIndex(DetalleCuentasPorCobrarVisorus []DetalleCuentasPorCobrarVisorusMgo) []map[string]interface{} {
	var Mapas []map[string]interface{}

	Sesion, err := MoConexion.GetConexionMgo()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer Sesion.Close()
	//floats := accounting.Accounting{Symbol: "", Precision: 2}
	for _, v := range DetalleCuentasPorCobrarVisorus {
		var Mapa map[string]interface{}
		Mapa = make(map[string]interface{})

		Mapa["NUMCFD"] = strconv.Itoa(int(v.NUMCFD))
		Mapa["CODIGO"] = v.CODIGO
		Mapa["DESCRIPCION"] = v.DESCRIPCION
		Mapa["UNIDAD"] = v.UNIDAD
		Mapa["CANTIDAD"] = strconv.FormatFloat(v.CANTIDAD, 'f', 2, 64)
		Mapa["PRECIO"] = strconv.FormatFloat(v.PRECIO, 'f', 2, 64)
		Mapa["PRECIOIVAIN"] = strconv.FormatFloat(v.PRECIOIVAIN, 'f', 2, 64)
		Mapa["DESCUENTO"] = strconv.FormatFloat(v.DESCUENTO, 'f', 2, 64)
		Mapa["IMPORTE"] = strconv.FormatFloat(v.IMPORTE, 'f', 2, 64)
		Mapa["TASA"] = v.TASA
		Mapa["IEPS"] = v.IEPS
		Mapa["VIEPS"] = strconv.FormatFloat(v.VIEPS, 'f', 2, 64)
		Mapa["NUMMOVCFD"] = strconv.Itoa(int(v.NUMMOVCFD))
		Mapa["NUMDTMCFD"] = strconv.Itoa(int(v.NUMDTMCFD))
		Mapa["CANTIDADM"] = strconv.FormatFloat(v.CANTIDADM, 'f', 2, 64)
		Mapa["COSTO"] = strconv.FormatFloat(v.COSTO, 'f', 2, 64)

		Mapas = append(Mapas, Mapa)

	}
	return Mapas

}

//GeneraDataRowDeIndex crea templates de tabla de búsqueda
func GeneraDataRowDeIndex(DetalleCuentasPorCobrarVisorus []DetalleCuentasPorCobrarVisorusMgo) map[string]interface{} {
	var Pagina = 1
	var totalRegistros = len(DetalleCuentasPorCobrarVisorus)
	var PaginasTotales = totalRegistros / 10

	if totalRegistros%10 != 0 {
		PaginasTotales++
	}

	Datos := GeneraRenglonesIndex(DetalleCuentasPorCobrarVisorus)
	var Mapa map[string]interface{}
	Mapa = make(map[string]interface{})

	Mapa["total"] = PaginasTotales
	Mapa["page"] = Pagina
	Mapa["records"] = totalRegistros
	Mapa["rows"] = Datos

	return Mapa
}

//MapaModelo regresa el mapa del modelo para jqgrid
func MapaModelo() map[string]interface{} {
	var ModelOptions map[string]interface{}
	ModelOptions = make(map[string]interface{})

	ModelOptions["name"] = []string{"NUMCFD", "CODIGO", "DESCRIPCION", "UNIDAD", "CANTIDAD", "PRECIO", "PRECIOIVAIN", "DESCUENTO", "IMPORTE", "TASA", "IEPS", "VIEPS", "NUMMOVCFD", "NUMDTMCFD", "CANTIDADM", "COSTO"}
	ModelOptions["index"] = []string{"NUMCFD", "CODIGO", "DESCRIPCION", "UNIDAD", "CANTIDAD", "PRECIO", "PRECIOIVAIN", "DESCUENTO", "IMPORTE", "TASA", "IEPS", "VIEPS", "NUMMOVCFD", "NUMDTMCFD", "CANTIDADM", "COSTO"}
	ModelOptions["align"] = []string{"left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left"}
	ModelOptions["editable"] = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	ModelOptions["hidden"] = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	ModelOptions["search"] = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	ModelOptions["stype"] = []string{"text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text"}
	ModelOptions["sorteable"] = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	ModelOptions["sorttype"] = []string{"numeric", "text", "text", "text", "float", "float", "float", "float", "float", "text", "text", "float", "numeric", "numeric", "float", "float"}
	var Textos map[string]interface{}
	Textos = make(map[string]interface{})
	Textos["sopt"] = []string{"eq", "ne", "cn"}

	var Numeros map[string]interface{}
	Numeros = make(map[string]interface{})
	Numeros["sopt"] = []string{"eq", "ne", "lt", "le", "gt", "ge"}

	var Fechas map[string]interface{}
	Fechas = make(map[string]interface{})

	Fechas["sopt"] = []string{"eq", "ne", "lt", "le", "gt", "ge"}

	Fechas["dataInit"] = `datePick`

	var Formatos map[string]interface{}
	Formatos = make(map[string]interface{})
	Formatos["srcformat"] = "d m y"
	Formatos["newformat"] = "d m y"

	Fechas["formatoptions"] = Formatos

	var Atributos map[string]interface{}
	Atributos = make(map[string]interface{})
	Atributos["title"] = "Seleccionar Fecha"

	Fechas["attr"] = Atributos

	/*
		var ComboEstatus map[string]interface{}
		ComboEstatus = make(map[string]interface{})
		ComboEstatus["sopt"] = []string{"eq", "ne", "cn"}
		ComboEstatus["value"] = CargaCombos.CargaComboCatalogoGrid(183)
	*/
	ModelOptions["searchoptions"] = []interface{}{Numeros, Textos, Textos, Textos, Numeros, Numeros, Numeros, Numeros, Numeros, Textos, Textos, Numeros, Numeros, Numeros, Numeros, Numeros}
	return ModelOptions
}

/*
//GeneraTemplatesBusqueda crea templates de tabla de búsqueda
func GeneraTemplatesBusqueda(DetalleCuentasPorCobrarVisoruss []DetalleCuentasPorCobrarVisorusMgo) (string, string) {
	floats := accounting.Accounting{Symbol: "", Precision: 2}
	cuerpo := ``

	cabecera := `<tr>
			<th>#</th>

				<th>NUMCFD </th>

				<th>SERIE </th>

				<th>FOLIO</th>

				<th>FECHA</th>

				<th>NUMCTE</th>

				<th>RFC</th>

				<th>SUBTOTAL</th>

				<th>DESCUENTO</th>

				<th>IMPUESTO</th>

				<th>IEPS</th>

				<th>TIPO_COMP</th>

				<th>STATUS</th>

				<th>FORPAG</th>

				<th>HORA</th>

				<th>NUMALM</th>

				<th>STATUS2</th>

				<th>TYPECFD</th>

				<th>NOMCTE</th>

				<th>DIRCTE</th>

				<th>NLICTE</th>

				<th>NLECTE</th>

				<th>COLCTE</th>

				<th>POBCTE</th>

				<th>EDOCTE</th>

				<th>PAISCTE</th>

				<th>CPCTE</th>

				<th>NUMMON</th>

				<th>FORPAG2</th>

				<th>NUMEYS</th>

				<th>STSCON</th>

				<th>NUMCFD_1</th>

				<th>CODIGO</th>

				<th>DESCRIPCION</th>

				<th>UNIDAD</th>

				<th>CANTIDAD</th>

				<th>PRECIO</th>

				<th>PRECIOIVAIN</th>

				<th>DESCUENTO</th>

				<th>IMPORTE</th>

				<th>TASA</th>

				<th>IEPS</th>

				<th>VIEPS</th>

				<th>NUMMOVCFD</th>

				<th>NUMDTMCFD</th>

				<th>CANTIDADM</th>

				<th>COSTO</th>
				</tr>`

	for k, v := range DetalleCuentasPorCobrarVisoruss {
		cuerpo += `<tr id = "` + v.ID.Hex() + `" onclick="window.location.href = '/DetalleCuentasPorCobrarVisoruss/detalle/` + v.ID.Hex() + `';">`
		cuerpo += `<td>` + strconv.Itoa(k+1) + `</td>`
		cuerpo += `<td>` + strconv.Itoa(int(v.NUMCFD)) + `</td>`

		cuerpo += `<td>` + v.SERIE + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.FOLIO)) + `</td>`

		cuerpo += `<td>` + v.FECHA.Format(time.RFC1123) + `</td>`

		cuerpo += `<td>` + v.NUMCTE + `</td>`

		cuerpo += `<td>` + v.RFC + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.SUBTOTAL) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.DESCUENTO) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.IMPUESTO) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.IEPS) + `</td>`

		cuerpo += `<td>` + v.TIPO_COMP + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.STATUS)) + `</td>`

		cuerpo += `<td>` + v.FORPAG + `</td>`

		cuerpo += `<td>` + v.HORA.Format(time.RFC1123) + `</td>`

		cuerpo += `<td>` + v.NUMALM + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.STATUS2)) + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.TYPECFD)) + `</td>`

		cuerpo += `<td>` + v.NOMCTE + `</td>`

		cuerpo += `<td>` + v.DIRCTE + `</td>`

		cuerpo += `<td>` + v.NLICTE + `</td>`

		cuerpo += `<td>` + v.NLECTE + `</td>`

		cuerpo += `<td>` + v.COLCTE + `</td>`

		cuerpo += `<td>` + v.POBCTE + `</td>`

		cuerpo += `<td>` + v.EDOCTE + `</td>`

		cuerpo += `<td>` + v.PAISCTE + `</td>`

		cuerpo += `<td>` + v.CPCTE + `</td>`

		cuerpo += `<td>` + v.NUMMON + `</td>`

		cuerpo += `<td>` + v.FORPAG2 + `</td>`

		cuerpo += `<td>` + v.NUMEYS + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.STSCON)) + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.NUMCFD_1)) + `</td>`

		cuerpo += `<td>` + v.CODIGO + `</td>`

		cuerpo += `<td>` + v.DESCRIPCION + `</td>`

		cuerpo += `<td>` + v.UNIDAD + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.CANTIDAD) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.PRECIO) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.PRECIOIVAIN) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.DESCUENTO) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.IMPORTE) + `</td>`

		cuerpo += `<td>` + v.TASA + `</td>`

		cuerpo += `<td>` + v.IEPS + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.VIEPS) + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.NUMMOVCFD)) + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.NUMDTMCFD)) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.CANTIDADM) + `</td>`

		cuerpo += `<td>` + floats.FormatMoney(v.COSTO) + `</td>`

		cuerpo += `</tr>`
	}

	return cabecera, cuerpo
}*/

//########## GET NAMES ####################################

//GetNameDetalleCuentasPorCobrarVisorus regresa el CODIGO del DetalleCuentasPorCobrarVisorus con el ID especificado
func GetNameDetalleCuentasPorCobrarVisorus(ID bson.ObjectId) string {
	var result DetalleCuentasPorCobrarVisorusMgo
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}
	DetalleCuentasPorCobrarVisoruss.Find(bson.M{"_id": ID}).One(&result)

	s.Close()
	return result.CODIGO
}

//########################< FUNCIONES GENERALES PSQL >#############################

//######################< FUNCIONES GENERALES ELASTIC >############################

//BuscarEnElastic busca el texto solicitado en los campos solicitados
func BuscarEnElastic(texto string) *elastic.SearchResult {
	textoTilde, textoQuotes := MoGeneral.ConstruirCadenas(texto)

	queryTilde := elastic.NewQueryStringQuery(textoTilde)
	queryQuotes := elastic.NewQueryStringQuery(textoQuotes)

	queryTilde = queryTilde.Field("NUMCFD ")
	queryQuotes = queryQuotes.Field("NUMCFD ")

	queryTilde = queryTilde.Field("CODIGO")
	queryQuotes = queryQuotes.Field("CODIGO")

	queryTilde = queryTilde.Field("DESCRIPCION")
	queryQuotes = queryQuotes.Field("DESCRIPCION")

	queryTilde = queryTilde.Field("UNIDAD")
	queryQuotes = queryQuotes.Field("UNIDAD")

	queryTilde = queryTilde.Field("CANTIDAD")
	queryQuotes = queryQuotes.Field("CANTIDAD")

	queryTilde = queryTilde.Field("PRECIO")
	queryQuotes = queryQuotes.Field("PRECIO")

	queryTilde = queryTilde.Field("PRECIOIVAIN")
	queryQuotes = queryQuotes.Field("PRECIOIVAIN")

	queryTilde = queryTilde.Field("DESCUENTO")
	queryQuotes = queryQuotes.Field("DESCUENTO")

	queryTilde = queryTilde.Field("IMPORTE")
	queryQuotes = queryQuotes.Field("IMPORTE")

	queryTilde = queryTilde.Field("TASA")
	queryQuotes = queryQuotes.Field("TASA")

	queryTilde = queryTilde.Field("IEPS")
	queryQuotes = queryQuotes.Field("IEPS")

	queryTilde = queryTilde.Field("VIEPS")
	queryQuotes = queryQuotes.Field("VIEPS")

	queryTilde = queryTilde.Field("NUMMOVCFD")
	queryQuotes = queryQuotes.Field("NUMMOVCFD")

	queryTilde = queryTilde.Field("NUMDTMCFD")
	queryQuotes = queryQuotes.Field("NUMDTMCFD")

	queryTilde = queryTilde.Field("CANTIDADM")
	queryQuotes = queryQuotes.Field("CANTIDADM")

	queryTilde = queryTilde.Field("COSTO")
	queryQuotes = queryQuotes.Field("COSTO")

	var docs *elastic.SearchResult
	var err bool

	docs, err = MoConexion.BuscaElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, queryTilde)
	if err {
		fmt.Println("No Match 1st Try")
	}

	if docs.Hits.TotalHits == 0 {
		docs, err = MoConexion.BuscaElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, queryQuotes)
		if err {
			fmt.Println("No Match 2nd Try")
		}
	}

	return docs
}

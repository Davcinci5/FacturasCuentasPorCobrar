package CuentasPorCobrarModel

import (
	"fmt"
	"strconv"
	"time"

	//"../../Modelos/CatalogoModel"

	"../../Modulos/Conexiones"
	"../../Modulos/General"

	"../../Modulos/Variables"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/olivere/elastic.v5"
)

//#########################< ESTRUCTURAS >##############################

//CuentasPorCobrarMgo estructura de CuentasPorCobrars mongo
type CuentasPorCobrarMgo struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	NUMCFD    int64         `bson:"NUMCFD"`
	SERIE     string        `bson:"SERIE"`
	FOLIO     int64         `bson:"FOLIO"`
	FECHA     time.Time     `bson:"FECHA"`
	NUMCTE    string        `bson:"NUMCTE"`
	RFC       string        `bson:"RFC"`
	SUBTOTAL  float64       `bson:"SUBTOTAL"`
	DESCUENTO float64       `bson:"DESCUENTO"`
	IMPUESTO  float64       `bson:"IMPUESTO"`
	IEPS      float64       `bson:"IEPS"`
	TIPOCOMP  string        `bson:"TIPOCOMP"`
	STATUS    int64         `bson:"STATUS"`
	FORPAG    string        `bson:"FORPAG"`
	HORA      time.Time     `bson:"HORA"`
	NUMALM    string        `bson:"NUMALM"`
	STATUS2   int64         `bson:"STATUS2"`
	TYPECFD   int64         `bson:"TYPECFD"`
	NOMCTE    string        `bson:"NOMCTE"`
	DIRCTE    string        `bson:"DIRCTE"`
	NLICTE    string        `bson:"NLICTE"`
	NLECTE    string        `bson:"NLECTE"`
	COLCTE    string        `bson:"COLCTE"`
	POBCTE    string        `bson:"POBCTE"`
	EDOCTE    string        `bson:"EDOCTE"`
	MUNCTE    string        `bson:"MUNCTE"`
	PAISCTE   string        `bson:"PAISCTE"`
	CPCTE     string        `bson:"CPCTE"`
	NUMMON    string        `bson:"NUMMON"`
	FORPAG2   string        `bson:"FORPAG2"`
	NUMEYS    string        `bson:"NUMEYS"`
	STSCON    int64         `bson:"STSCON"`
}

//CuentasPorCobrarElastic estructura de CuentasPorCobrars para insertar en Elastic
type CuentasPorCobrarElastic struct {
	NUMCFD    int64     `json:"NUMCFD"`
	SERIE     string    `json:"SERIE"`
	FOLIO     int64     `json:"FOLIO"`
	FECHA     time.Time `json:"FECHA"`
	NUMCTE    string    `json:"NUMCTE"`
	RFC       string    `json:"RFC"`
	SUBTOTAL  float64   `json:"SUBTOTAL"`
	DESCUENTO float64   `json:"DESCUENTO"`
	IMPUESTO  float64   `json:"IMPUESTO"`
	IEPS      float64   `json:"IEPS"`
	TIPOCOMP  string    `json:"TIPOCOMP"`
	STATUS    int64     `json:"STATUS"`
	FORPAG    string    `json:"FORPAG"`
	HORA      time.Time `json:"HORA"`
	NUMALM    string    `json:"NUMALM"`
	STATUS2   int64     `json:"STATUS2"`
	TYPECFD   int64     `json:"TYPECFD"`
	NOMCTE    string    `json:"NOMCTE"`
	DIRCTE    string    `json:"DIRCTE"`
	NLICTE    string    `json:"NLICTE"`
	NLECTE    string    `json:"NLECTE"`
	COLCTE    string    `json:"COLCTE"`
	POBCTE    string    `json:"POBCTE"`
	EDOCTE    string    `json:"EDOCTE"`
	MUNCTE    string    `json:"MUNCTE"`
	PAISCTE   string    `json:"PAISCTE"`
	CPCTE     string    `json:"CPCTE"`
	NUMMON    string    `json:"NUMMON"`
	FORPAG2   string    `json:"FORPAG2"`
	NUMEYS    string    `json:"NUMEYS"`
	STSCON    int64     `json:"STSCON"`
}

//#########################< FUNCIONES GENERALES MGO >###############################

//GetAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func GetAll() []CuentasPorCobrarMgo {
	var result []CuentasPorCobrarMgo
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	err = CuentasPorCobrars.Find(nil).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//CountAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func CountAll() int {
	var result int
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)

	if err != nil {
		fmt.Println(err)
	}
	result, err = CuentasPorCobrars.Find(nil).Count()
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetOne Regresa un documento específico de Mongo (Por Coleccion)
func GetOne(ID bson.ObjectId) CuentasPorCobrarMgo {
	var result CuentasPorCobrarMgo
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	err = CuentasPorCobrars.Find(bson.M{"_id": ID}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetEspecifics rsegresa un conjunto de documentos específicos de Mongo (Por Coleccion)
func GetEspecifics(Ides []bson.ObjectId) []CuentasPorCobrarMgo {
	var result []CuentasPorCobrarMgo
	var aux CuentasPorCobrarMgo
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range Ides {
		aux = CuentasPorCobrarMgo{}
		CuentasPorCobrars.Find(bson.M{"_id": value}).One(&aux)
		result = append(result, aux)
	}
	s.Close()
	return result
}

//GetEspecificByFields regresa un documento de Mongo especificando un campo y un determinado valor
func GetEspecificByFields(field string, valor interface{}) CuentasPorCobrarMgo {
	var result CuentasPorCobrarMgo
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)

	if err != nil {
		fmt.Println(err)
	}
	err = CuentasPorCobrars.Find(bson.M{field: valor}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//GetIDByField regresa un documento específico de Mongo (Por Coleccion)
func GetIDByField(field string, valor interface{}) bson.ObjectId {
	var result CuentasPorCobrarMgo
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	err = CuentasPorCobrars.Find(bson.M{field: valor}).One(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result.ID
}

//CargaComboCuentasPorCobrars regresa un combo de CuentasPorCobrar de mongo
func CargaComboCuentasPorCobrars(ID string) string {
	CuentasPorCobrars := GetAll()

	templ := ``

	if ID != "" {
		templ = `<option value="">--SELECCIONE--</option> `
	} else {
		templ = `<option value="" selected>--SELECCIONE--</option> `
	}

	for _, v := range CuentasPorCobrars {
		if ID == v.ID.Hex() {
			templ += `<option value="` + v.ID.Hex() + `" selected>  ` + v.SERIE + ` </option> `
		} else {
			templ += `<option value="` + v.ID.Hex() + `">  ` + v.SERIE + ` </option> `
		}

	}
	return templ
}

//GeneraRenglonesIndex crea un mapa para crear el json de renglones del index
func GeneraRenglonesIndex(CuentasPorCobrar []CuentasPorCobrarMgo) []map[string]interface{} {
	var Mapas []map[string]interface{}

	Sesion, err := MoConexion.GetConexionMgo()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer Sesion.Close()
	//floats := accounting.Accounting{Symbol: "", Precision: 2}
	for _, v := range CuentasPorCobrar {
		var Mapa map[string]interface{}
		Mapa = make(map[string]interface{})

		Mapa["NUMCFD"] = strconv.Itoa(int(v.NUMCFD))
		Mapa["SERIE"] = v.SERIE
		Mapa["FOLIO"] = strconv.Itoa(int(v.FOLIO))
		Mapa["FECHA"] = v.FECHA.Format("2006/1/2")
		Mapa["NUMCTE"] = v.NUMCTE
		Mapa["RFC"] = v.RFC
		Mapa["SUBTOTAL"] = strconv.FormatFloat(v.SUBTOTAL, 'f', 2, 64)
		Mapa["DESCUENTO"] = strconv.FormatFloat(v.DESCUENTO, 'f', 2, 64)
		Mapa["IMPUESTO"] = strconv.FormatFloat(v.IMPUESTO, 'f', 2, 64)
		Mapa["IEPS"] = strconv.FormatFloat(v.IEPS, 'f', 2, 64)
		Mapa["TIPOCOMP"] = v.TIPOCOMP
		Mapa["STATUS"] = strconv.Itoa(int(v.STATUS))
		Mapa["FORPAG"] = v.FORPAG
		Mapa["HORA"] = v.HORA.Format("2006/1/2")
		Mapa["NUMALM"] = v.NUMALM
		Mapa["STATUS2"] = strconv.Itoa(int(v.STATUS2))
		Mapa["TYPECFD"] = strconv.Itoa(int(v.TYPECFD))
		Mapa["NOMCTE"] = v.NOMCTE
		Mapa["DIRCTE"] = v.DIRCTE
		Mapa["NLICTE"] = v.NLICTE
		Mapa["NLECTE"] = v.NLECTE
		Mapa["COLCTE"] = v.COLCTE
		Mapa["POBCTE"] = v.POBCTE
		Mapa["EDOCTE"] = v.EDOCTE
		Mapa["MUNCTE"] = v.MUNCTE
		Mapa["PAISCTE"] = v.PAISCTE
		Mapa["CPCTE"] = v.CPCTE
		Mapa["NUMMON"] = v.NUMMON
		Mapa["FORPAG2"] = v.FORPAG2
		Mapa["NUMEYS"] = v.NUMEYS
		Mapa["STSCON"] = strconv.Itoa(int(v.STSCON))

		//Mapa["VersionKore"] = strconv.Itoa(v.VersionKore)
		//Mapa["FechaInicioVigencia"] = v.FechaInicioVigencia.Format("2006/1/2")

		Mapas = append(Mapas, Mapa)

	}
	return Mapas

}

//GeneraDataRowDeIndex crea templates de tabla de búsqueda
func GeneraDataRowDeIndex(CuentasPorCobrar []CuentasPorCobrarMgo) map[string]interface{} {
	var Pagina = 1
	var totalRegistros = len(CuentasPorCobrar)
	var PaginasTotales = totalRegistros / 10

	if totalRegistros%10 != 0 {
		PaginasTotales++
	}

	Datos := GeneraRenglonesIndex(CuentasPorCobrar)
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
	ModelOptions["name"] = []string{"NUMCFD", "SERIE", "FOLIO", "FECHA", "NUMCTE", "RFC", "SUBTOTAL", "DESCUENTO", "IMPUESTO", "IEPS", "TIPOCOMP", "STATUS", "FORPAG", "HORA", "NUMALM", "STATUS2", "TYPECFD", "NOMCTE", "DIRCTE", "NLICTE", "NLECTE", "COLCTE", "POBCTE", "EDOCTE", "MUNCTE", "PAISCTE", "CPCTE", "NUMMON", "FORPAG2", "NUMEYS", "STSCON"}
	ModelOptions["index"] = []string{"NUMCFD", "SERIE", "FOLIO", "FECHA", "NUMCTE", "RFC", "SUBTOTAL", "DESCUENTO", "IMPUESTO", "IEPS", "TIPOCOMP", "STATUS", "FORPAG", "HORA", "NUMALM", "STATUS2", "TYPECFD", "NOMCTE", "DIRCTE", "NLICTE", "NLECTE", "COLCTE", "POBCTE", "EDOCTE", "MUNCTE", "PAISCTE", "CPCTE", "NUMMON", "FORPAG2", "NUMEYS", "STSCON"}
	ModelOptions["align"] = []string{"left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left", "left"}
	ModelOptions["editable"] = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	ModelOptions["hidden"] = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	ModelOptions["search"] = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	ModelOptions["stype"] = []string{"text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text"}
	ModelOptions["sorteable"] = []bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	ModelOptions["sorttype"] = []string{"numeric", "text", "numeric", "date", "text", "text", "float", "float", "float", "float", "text", "numeric", "text", "date", "text", "numeric", "numeric", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "text", "numeric"}

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
	ModelOptions["searchoptions"] = []interface{}{Numeros, Textos, Numeros, Fechas, Textos, Textos, Numeros, Numeros, Numeros, Numeros, Textos, Numeros, Textos, Fechas, Textos, Numeros, Numeros, Textos, Textos, Textos, Textos, Textos, Textos, Textos, Textos, Textos, Textos, Textos, Textos, Textos, Numeros}
	return ModelOptions
}

//GeneraTemplatesBusqueda crea templates de tabla de búsqueda
/*func GeneraTemplatesBusqueda(CuentasPorCobrars []CuentasPorCobrarMgo) (string, string) {
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

				<th>TIPOCOMP</th>

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

				<th>MUNCTE</th>

				<th>PAISCTE</th>

				<th>CPCTE</th>

				<th>NUMMON</th>

				<th>FORPAG2</th>

				<th>NUMEYS</th>

				<th>STSCON</th>
				</tr>`

	for k, v := range CuentasPorCobrars {
		cuerpo += `<tr id = "` + v.ID.Hex() + `" onclick="window.location.href = '/CuentasPorCobrars/detalle/` + v.ID.Hex() + `';">`
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

		cuerpo += `<td>` + v.TIPOCOMP + `</td>`

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

		cuerpo += `<td>` + v.MUNCTE + `</td>`

		cuerpo += `<td>` + v.PAISCTE + `</td>`

		cuerpo += `<td>` + v.CPCTE + `</td>`

		cuerpo += `<td>` + v.NUMMON + `</td>`

		cuerpo += `<td>` + v.FORPAG2 + `</td>`

		cuerpo += `<td>` + v.NUMEYS + `</td>`

		cuerpo += `<td>` + strconv.Itoa(int(v.STSCON)) + `</td>`

		cuerpo += `</tr>`
	}

	return cabecera, cuerpo
}*/

//########## GET NAMES ####################################

//GetNameCuentasPorCobrar regresa el nombre del CuentasPorCobrar con el ID especificado
func GetNameCuentasPorCobrar(ID bson.ObjectId) string {
	var result CuentasPorCobrarMgo
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	CuentasPorCobrars.Find(bson.M{"_id": ID}).One(&result)

	s.Close()
	return result.SERIE
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

	queryTilde = queryTilde.Field("SERIE ")
	queryQuotes = queryQuotes.Field("SERIE ")

	queryTilde = queryTilde.Field("FOLIO")
	queryQuotes = queryQuotes.Field("FOLIO")

	queryTilde = queryTilde.Field("FECHA")
	queryQuotes = queryQuotes.Field("FECHA")

	queryTilde = queryTilde.Field("NUMCTE")
	queryQuotes = queryQuotes.Field("NUMCTE")

	queryTilde = queryTilde.Field("RFC")
	queryQuotes = queryQuotes.Field("RFC")

	queryTilde = queryTilde.Field("SUBTOTAL")
	queryQuotes = queryQuotes.Field("SUBTOTAL")

	queryTilde = queryTilde.Field("DESCUENTO")
	queryQuotes = queryQuotes.Field("DESCUENTO")

	queryTilde = queryTilde.Field("IMPUESTO")
	queryQuotes = queryQuotes.Field("IMPUESTO")

	queryTilde = queryTilde.Field("IEPS")
	queryQuotes = queryQuotes.Field("IEPS")

	queryTilde = queryTilde.Field("TIPOCOMP")
	queryQuotes = queryQuotes.Field("TIPOCOMP")

	queryTilde = queryTilde.Field("STATUS")
	queryQuotes = queryQuotes.Field("STATUS")

	queryTilde = queryTilde.Field("FORPAG")
	queryQuotes = queryQuotes.Field("FORPAG")

	queryTilde = queryTilde.Field("HORA")
	queryQuotes = queryQuotes.Field("HORA")

	queryTilde = queryTilde.Field("NUMALM")
	queryQuotes = queryQuotes.Field("NUMALM")

	queryTilde = queryTilde.Field("STATUS2")
	queryQuotes = queryQuotes.Field("STATUS2")

	queryTilde = queryTilde.Field("TYPECFD")
	queryQuotes = queryQuotes.Field("TYPECFD")

	queryTilde = queryTilde.Field("NOMCTE")
	queryQuotes = queryQuotes.Field("NOMCTE")

	queryTilde = queryTilde.Field("DIRCTE")
	queryQuotes = queryQuotes.Field("DIRCTE")

	queryTilde = queryTilde.Field("NLICTE")
	queryQuotes = queryQuotes.Field("NLICTE")

	queryTilde = queryTilde.Field("NLECTE")
	queryQuotes = queryQuotes.Field("NLECTE")

	queryTilde = queryTilde.Field("COLCTE")
	queryQuotes = queryQuotes.Field("COLCTE")

	queryTilde = queryTilde.Field("POBCTE")
	queryQuotes = queryQuotes.Field("POBCTE")

	queryTilde = queryTilde.Field("EDOCTE")
	queryQuotes = queryQuotes.Field("EDOCTE")

	queryTilde = queryTilde.Field("MUNCTE")
	queryQuotes = queryQuotes.Field("MUNCTE")

	queryTilde = queryTilde.Field("PAISCTE")
	queryQuotes = queryQuotes.Field("PAISCTE")

	queryTilde = queryTilde.Field("CPCTE")
	queryQuotes = queryQuotes.Field("CPCTE")

	queryTilde = queryTilde.Field("NUMMON")
	queryQuotes = queryQuotes.Field("NUMMON")

	queryTilde = queryTilde.Field("FORPAG2")
	queryQuotes = queryQuotes.Field("FORPAG2")

	queryTilde = queryTilde.Field("NUMEYS")
	queryQuotes = queryQuotes.Field("NUMEYS")

	queryTilde = queryTilde.Field("STSCON")
	queryQuotes = queryQuotes.Field("STSCON")

	var docs *elastic.SearchResult
	var err bool

	docs, err = MoConexion.BuscaElastic(MoVar.TipoCuentasPorCobrar, queryTilde)
	if err {
		fmt.Println("No Match 1st Try")
	}

	if docs.Hits.TotalHits == 0 {
		docs, err = MoConexion.BuscaElastic(MoVar.TipoCuentasPorCobrar, queryQuotes)
		if err {
			fmt.Println("No Match 2nd Try")
		}
	}

	return docs
}

package MoVar

import (
	config "github.com/robfig/config"
)

const (

	//############# ARCHIVOS LOCALES ######################################

	//FileConfigName contiene el nombre del archivo CFG
	FileConfigName = "Acfg.cfg"

	//################ SECCIONES CFG  ######################################

	//SecDefault nombre de la seccion default del servidor en CFG
	SecDefault = "DEFAULT"
	//SecMongo nombre de la seccion de mongo en CFG
	SecMongo = "CONFIG_DB_MONGO"
	//SecPsql nombre de la seccion de postgresql en cfg
	SecPsql = "CONFIG_DB_POSTGRES"
	//SecElastic nombre de la seccion de postgresql en cfg
	SecElastic = "CONFIG_DB_ELASTIC"

	//############# COLUMNAS DE ALMACEN PSQL ######################################

	//############# COLECCIONES MONGO ######################################

	//Mongo---------------> Catalogo

	//ColeccionCatalogo nombre de la coleccion de Catalogo en mongo
	ColeccionCatalogo = "Catalogo"

	//Mongo---------------> CatalogoNuevo

	//ColeccionCatalogoNuevo nombre de la coleccion de CatalogoNuevo en mongo
	ColeccionCatalogoNuevo = "CatalogoNuevo"

	//Mongo---------------> Empresa

	//ColeccionEmpresa nombre de la coleccion de Empresa en mongo
	ColeccionEmpresa = "Empresa"

	//Mongo---------------> Facturacion

	//ColeccionFacturacion nombre de la coleccion de Facturacion en mongo
	ColeccionFacturacion = "Facturacion"

	//Mongo---------------> Conexion

	//ColeccionConexion nombre de la coleccion de Conexion en mongo
	ColeccionConexion = "Conexion"

	//Mongo---------------> Bug

	//ColeccionBug nombre de la coleccion de Bug en mongo
	ColeccionBug = "Bug"

	////////////////////////////////////////////////////////////////////////77
	//FUNCIONES AGREGADAS PARA CATALOGOS SAT
	////////////////////////////////////////////////////////////////////////
	//ColeccionCodigosPostales nombre de la coleccion de CodigosPostales en mongo
	ColeccionCodigosPostales = "CodigosPostales"

	//ColeccionUnidades nombre de la coleccion de Unidades en mongo
	ColeccionUnidades = "Unidades"

	//ColeccionProductoServicios nombre de la coleccion de ProductoServicios en mongo
	ColeccionProductoServicios = "ProductoServicios"

	//ColeccionCuentasPorCobrar nombre de la coleccion de CuentasPorCobrar en mongo
	ColeccionCuentasPorCobrar = "CuentasPorCobrar"

	//##########################<CATÁLOGOS DEL SISTEMA>######################

	//CatSysTipoImpuesto nombre de la colección de tipos de impuestos del sistema
	CatSysTipoImpuesto = "CatSysTiposDeImpuestos"

	//CatSysTipoFactor nombre de la colección de tipos de factores de impuestos del sistema
	CatSysTipoFactor = "CatSysTipoDeFactor"

	//CatSysClasificacionDeimpuestos nombre de la colección de tipos de factores de impuestos del sistema
	CatSysClasificacionDeimpuestos = "CatSysClasificacionDeimpuestos"

	//CatSysSubClasificacionDeimpuestos nombre de la colección de tipos de factores de impuestos del sistema
	CatSysSubClasificacionDeimpuestos = "CatSysSubClasificacionDeimpuestos"

	//################# DATOS ELASTIC ######################################

	//Elastic---------------> Catalogo

	//TipoCatalogo tipo a manejar en elastic
	TipoCatalogo = "Catalogo"

	//Elastic---------------> CatalogoNuevo

	//TipoCatalogoNuevo tipo a manejar en elastic
	TipoCatalogoNuevo = "CatalogoNuevo"

	//Elastic---------------> Empresa

	//TipoEmpresa tipo a manejar en elastic
	TipoEmpresa = "Empresa"

	//Elastic---------------> Caja

	//TipoCaja tipo a manejar en elastic
	TipoCaja = "Caja"

	//Elastic---------------> Conexion

	//TipoConexion tipo a manejar en elastic
	TipoConexion = "Conexion"

	//Elastic---------------> Bug

	//TipoBug tipo a manejar en elastic
	TipoBug = "Bug"

	//IndexElastic nombre del index a usar en elastic
	IndexElastic = "minisuperampliado"

	//TipoCuentasPorCobrar tipo a manejar en elastic
	TipoCuentasPorCobrar = "CuentasPorCobrar"

	// * -----------------------CLAVES DE CATALOGO ------------------- * //

	// CatalogoDeEstatusDeCliente numero de catalogo en la coleccion Catalogos
	CatalogoDeEstatusDeCliente = 137
)

//DataCfg estructura de datos del entorno
type DataCfg struct {
	BaseURL    string
	Servidor   string
	Puerto     string
	Usuario    string
	Pass       string
	Protocolo  string
	NombreBase string
}

//#################<Funciones Generales>#######################################

//CargaSeccionCFG rellena los datos de la seccion a utilizar
func CargaSeccionCFG(seccion string) DataCfg {
	var d DataCfg
	var FileConfig, err = config.ReadDefault(FileConfigName)
	if err == nil {
		if FileConfig.HasOption(seccion, "baseurl") {
			d.BaseURL, _ = FileConfig.String(seccion, "baseurl")
		}
		if FileConfig.HasOption(seccion, "servidor") {
			d.Servidor, _ = FileConfig.String(seccion, "servidor")
		}
		if FileConfig.HasOption(seccion, "puerto") {
			d.Puerto, _ = FileConfig.String(seccion, "puerto")
		}
		if FileConfig.HasOption(seccion, "usuario") {
			d.Usuario, _ = FileConfig.String(seccion, "usuario")
		}
		if FileConfig.HasOption(seccion, "pass") {
			d.Pass, _ = FileConfig.String(seccion, "pass")
		}
		if FileConfig.HasOption(seccion, "protocolo") {
			d.Protocolo, _ = FileConfig.String(seccion, "protocolo")
		}
		if FileConfig.HasOption(seccion, "base") {
			d.NombreBase, _ = FileConfig.String(seccion, "base")
		}
	}
	return d
}

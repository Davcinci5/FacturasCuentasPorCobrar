package CatalogosSATModel

import (
	"fmt"
	"time"

	"../../Modulos/Conexiones"
	"../../Modulos/Variables"
	"gopkg.in/mgo.v2/bson"
)

//CatalogoProductosServiciosMgo para insertar datos en mongo
type ProductoServiciosMgo struct {
	ID                    bson.ObjectId `bson:"_id,omitempty"`
	VersionSat            float32       `bson:"VersionSat"`
	VersionKore           int           `bson:"VersionKore"`
	Clave                 string        `bson:"Clave"`
	Descripcion           string        `bson:"Descripcion"`
	FechaInicioVigencia   time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia      time.Time     `bson:"FechaFinVigencia"`
	IncluirIVATrasladado  string        `bson:"IncluirIVATrasladado"`
	IncluirIEPSTrasladado string        `bson:"IncluirIEPSTrasladado"`
	Complemento           string        `bson:"Complemento"`
	Estatus               string        `bson:"Estatus"`
}

//GetAllProductoServicios Regresa todos los documentos existentes de Mongo (Por Coleccion)
func GetAllProductoServicios() []ProductoServiciosMgo {
	var result []ProductoServiciosMgo
	s, ProductoServicioss, err := MoConexion.GetColectionMgo(MoVar.ColeccionProductoServicios)
	if err != nil {
		fmt.Println(err)
	}
	err = ProductoServicioss.Find(nil).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//-------------------------------------------------------------------------------------------------------------------------------------------------------------
//CatalogoUnidadesMgo para insertar datos en mongo
type UnidadesMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Nombre              string        `bson:"Nombre"`
	Descripcion         string        `bson:"Descripcion"`
	Nota                string        `bson:"Nota"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Simbolo             string        `bson:"Simbolo"`
	Estatus             string        `bson:"Estatus"`
}

//GetAll Regresa todos los documentos existentes de Mongo (Por Coleccion)
func GetAllUnidades() []UnidadesMgo {

	var result []UnidadesMgo
	s, Unidadess, err := MoConexion.GetColectionMgo(MoVar.ColeccionUnidades)
	if err != nil {
		fmt.Println(err)
	}
	err = Unidadess.Find(nil).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	s.Close()
	return result
}

//------------------------------------------------------------------------------------------------------------------------------------------------------------
//CodigosPostalesMgo para insertar datos en mongo
type CodigosPostalesMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	CodigoPostal        string        `bson:"CodigoPostal"`
	Estado              string        `bson:"Estado"`
	Municipio           string        `bson:"Municipio"`
	Localidad           string        `bson:"Localidad"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------
//FormasPagoMgo para insertar datos en mongo
type FormasPagoMgo struct {
	ID                                        bson.ObjectId `bson:"_id,omitempty"`
	VersionSat                                float32       `bson:"VersionSat"`
	VersionKore                               int           `bson:"VersionKore"`
	Clave                                     string        `bson:"Clave"`
	Descripcion                               string        `bson:"Descripcion"`
	Bancarizado                               string        `bson:"Bancarizado"`
	NumeroDEOperación                         string        `bson:"NumeroDEOperación"`
	RFCDelEmisorDeLaCuentaOrdenante           string        `bson:"RFCDelEmisorDeLaCuentaOrdenante"`
	CuentaOrdenante                           string        `bson:"CuentaOrdenante"`
	PatronParaCuentaOrdenante                 string        `bson:"PatronParaCuentaOrdenante"`
	RFCDelEmisorCuentaDeBeneficiario          string        `bson:"RFCDelEmisorCuentaDeBeneficiario"`
	CuentaDEBenenficiario                     string        `bson:"CuentaDEBenenficiario"`
	PatronParaCuentaBeneficiaria              string        `bson:"PatronParaCuentaBeneficiaria"`
	TipoCadenaPago                            string        `bson:"TipoCadenaPago"`
	NombreBancoEmisorCuentaOrdenantExtranjero string        `bson:"NombreBancoEmisorCuentaOrdenantExtranjero"`
	FechaInicioVigencia                       time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia                          time.Time     `bson:"FechaFinVigencia"`
	Estatus                                   string        `bson:"Estatus"`
}

//ImpuestosMgo para insertar datos en mongo
type ImpuestosMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Descripcion         string        `bson:"Descripcion"`
	Retencion           string        `bson:"Retencion"`
	Traslado            string        `bson:"Traslado"`
	LocalOFederal       string        `bson:"LocalOFederal"`
	EntidadAplica       string        `bson:"EntidadAplica"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//MetodoDEPagoMgo para insertar datos en mongo
type MetodoDEPagoMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Descripcion         string        `bson:"Descripcion"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//MonedaMgo para insertar datos en mongo
type MonedaMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Descripcion         string        `bson:"Descripcion"`
	Decimales           int8          `bson:"Decimales"`
	PorcentajeVariacion float32       `bson:"PorcentajeVariacion"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//RegimenFiscalMgo para insertar datos en mongo
type RegimenFiscalMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Descripcion         string        `bson:"Descripcion"`
	Fisica              string        `bson:"Fisica"`
	Moral               string        `bson:"Moral"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//TiposDEComprobanteMgo para insertar datos en mongo
type TiposDEComprobanteMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Descripcion         string        `bson:"Descripcion"`
	ValorMaximo         string        `bson:"ValorMaximo"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//CatalogoTipoFactorMgo para insertar datos en mongo
type TipoFactorMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//TiposDERelacuionEntreCFDIMgo para insertar datos en mongo
type TiposDERelacuionEntreCFDIMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Descripcion         string        `bson:"Descripcion"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

//UsoDEComprobantesMgo para insertar datos en mongo
type UsoDEComprobantesMgo struct {
	ID                  bson.ObjectId `bson:"_id,omitempty"`
	VersionSat          float32       `bson:"VersionSat"`
	VersionKore         int           `bson:"VersionKore"`
	Clave               string        `bson:"Clave"`
	Descripcion         string        `bson:"Descripcion"`
	AplicaPersonaFisica string        `bson:"AplicaPersonaFisica"`
	AplicaPersonaMoral  string        `bson:"AplicaPersonaMoral"`
	FechaInicioVigencia time.Time     `bson:"FechaInicioVigencia"`
	FechaFinVigencia    time.Time     `bson:"FechaFinVigencia"`
	Estatus             string        `bson:"Estatus"`
}

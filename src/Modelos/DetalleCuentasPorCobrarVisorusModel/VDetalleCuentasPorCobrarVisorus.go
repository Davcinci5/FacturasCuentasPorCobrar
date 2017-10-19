package DetalleCuentasPorCobrarVisorusModel

import (
	"html/template"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//#########################< ESTRUCTURAS >##############################

//ENUMCFD DetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMCFDDetalleCuentasPorCobrarVisorus struct {
	NUMCFD   int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ESERIE DetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ESERIEDetalleCuentasPorCobrarVisorus struct {
	SERIE    string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EFOLIODetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EFOLIODetalleCuentasPorCobrarVisorus struct {
	FOLIO    int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EFECHADetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EFECHADetalleCuentasPorCobrarVisorus struct {
	FECHA    time.Time
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENUMCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMCTEDetalleCuentasPorCobrarVisorus struct {
	NUMCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ERFCDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ERFCDetalleCuentasPorCobrarVisorus struct {
	RFC      string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ESUBTOTALDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ESUBTOTALDetalleCuentasPorCobrarVisorus struct {
	SUBTOTAL float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EDESCUENTODetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EDESCUENTODetalleCuentasPorCobrarVisorus struct {
	DESCUENTO float64
	IEstatus  bool
	IMsj      string
	Ihtml     template.HTML
}

//EIMPUESTODetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EIMPUESTODetalleCuentasPorCobrarVisorus struct {
	IMPUESTO float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EIEPSDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EIEPSDetalleCuentasPorCobrarVisorus struct {
	IEPS     float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ETIPO_COMPDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ETIPO_COMPDetalleCuentasPorCobrarVisorus struct {
	TIPO_COMP string
	IEstatus  bool
	IMsj      string
	Ihtml     template.HTML
}

//ESTATUSDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ESTATUSDetalleCuentasPorCobrarVisorus struct {
	STATUS   int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EFORPAGDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EFORPAGDetalleCuentasPorCobrarVisorus struct {
	FORPAG   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EHORADetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EHORADetalleCuentasPorCobrarVisorus struct {
	HORA     time.Time
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENUMALMDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMALMDetalleCuentasPorCobrarVisorus struct {
	NUMALM   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ESTATUS2DetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ESTATUS2DetalleCuentasPorCobrarVisorus struct {
	STATUS2  int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ETYPECFDDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ETYPECFDDetalleCuentasPorCobrarVisorus struct {
	TYPECFD  int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENOMCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENOMCTEDetalleCuentasPorCobrarVisorus struct {
	NOMCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EDIRCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EDIRCTEDetalleCuentasPorCobrarVisorus struct {
	DIRCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENLICTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENLICTEDetalleCuentasPorCobrarVisorus struct {
	NLICTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENLECTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENLECTEDetalleCuentasPorCobrarVisorus struct {
	NLECTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ECOLCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ECOLCTEDetalleCuentasPorCobrarVisorus struct {
	COLCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EPOBCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EPOBCTEDetalleCuentasPorCobrarVisorus struct {
	POBCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EEDOCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EEDOCTEDetalleCuentasPorCobrarVisorus struct {
	EDOCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EPAISCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EPAISCTEDetalleCuentasPorCobrarVisorus struct {
	PAISCTE  string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ECPCTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ECPCTEDetalleCuentasPorCobrarVisorus struct {
	CPCTE    string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENUMMONDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMMONDetalleCuentasPorCobrarVisorus struct {
	NUMMON   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EFORPAG2DetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EFORPAG2DetalleCuentasPorCobrarVisorus struct {
	FORPAG2  string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENUMEYSDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMEYSDetalleCuentasPorCobrarVisorus struct {
	NUMEYS   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ESTSCONDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ESTSCONDetalleCuentasPorCobrarVisorus struct {
	STSCON   int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENUMCFD_1DetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMCFD_1DetalleCuentasPorCobrarVisorus struct {
	NUMCFD_1 int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ECODIGODetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ECODIGODetalleCuentasPorCobrarVisorus struct {
	CODIGO   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EDESCRIPCIONDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EDESCRIPCIONDetalleCuentasPorCobrarVisorus struct {
	DESCRIPCION string
	IEstatus    bool
	IMsj        string
	Ihtml       template.HTML
}

//EUNIDADDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EUNIDADDetalleCuentasPorCobrarVisorus struct {
	UNIDAD   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ECANTIDADDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ECANTIDADDetalleCuentasPorCobrarVisorus struct {
	CANTIDAD float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EPRECIODetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EPRECIODetalleCuentasPorCobrarVisorus struct {
	PRECIO   float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EPRECIOIVAINDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EPRECIOIVAINDetalleCuentasPorCobrarVisorus struct {
	PRECIOIVAIN float64
	IEstatus    bool
	IMsj        string
	Ihtml       template.HTML
}

//EDESCUENTO_1DetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EDESCUENTO_1DetalleCuentasPorCobrarVisorus struct {
	DESCUENTO_1 float64
	IEstatus    bool
	IMsj        string
	Ihtml       template.HTML
}

//EIMPORTEDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EIMPORTEDetalleCuentasPorCobrarVisorus struct {
	IMPORTE  float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ETASADetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ETASADetalleCuentasPorCobrarVisorus struct {
	TASA     string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EIEPS_1DetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EIEPS_1DetalleCuentasPorCobrarVisorus struct {
	IEPS_1   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//EVIEPSDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type EVIEPSDetalleCuentasPorCobrarVisorus struct {
	VIEPS    float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//ENUMMOVCFDDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMMOVCFDDetalleCuentasPorCobrarVisorus struct {
	NUMMOVCFD int64
	IEstatus  bool
	IMsj      string
	Ihtml     template.HTML
}

//ENUMDTMCFDDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ENUMDTMCFDDetalleCuentasPorCobrarVisorus struct {
	NUMDTMCFD int64
	IEstatus  bool
	IMsj      string
	Ihtml     template.HTML
}

//ECANTIDADMDetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ECANTIDADMDetalleCuentasPorCobrarVisorus struct {
	CANTIDADM float64
	IEstatus  bool
	IMsj      string
	Ihtml     template.HTML
}

//ECOSTODetalleCuentasPorCobrarVisorus Estructura de campo de DetalleCuentasPorCobrarVisorus
type ECOSTODetalleCuentasPorCobrarVisorus struct {
	COSTO    float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//DetalleCuentasPorCobrarVisorus estructura de DetalleCuentasPorCobrarVisoruss mongo
type DetalleCuentasPorCobrarVisorus struct {
	ID bson.ObjectId
	ENUMCFDDetalleCuentasPorCobrarVisorus
	ESERIEDetalleCuentasPorCobrarVisorus
	EFOLIODetalleCuentasPorCobrarVisorus
	EFECHADetalleCuentasPorCobrarVisorus
	ENUMCTEDetalleCuentasPorCobrarVisorus
	ERFCDetalleCuentasPorCobrarVisorus
	ESUBTOTALDetalleCuentasPorCobrarVisorus
	EDESCUENTODetalleCuentasPorCobrarVisorus
	EIMPUESTODetalleCuentasPorCobrarVisorus
	EIEPSDetalleCuentasPorCobrarVisorus
	ETIPO_COMPDetalleCuentasPorCobrarVisorus
	ESTATUSDetalleCuentasPorCobrarVisorus
	EFORPAGDetalleCuentasPorCobrarVisorus
	EHORADetalleCuentasPorCobrarVisorus
	ENUMALMDetalleCuentasPorCobrarVisorus
	ESTATUS2DetalleCuentasPorCobrarVisorus
	ETYPECFDDetalleCuentasPorCobrarVisorus
	ENOMCTEDetalleCuentasPorCobrarVisorus
	EDIRCTEDetalleCuentasPorCobrarVisorus
	ENLICTEDetalleCuentasPorCobrarVisorus
	ENLECTEDetalleCuentasPorCobrarVisorus
	ECOLCTEDetalleCuentasPorCobrarVisorus
	EPOBCTEDetalleCuentasPorCobrarVisorus
	EEDOCTEDetalleCuentasPorCobrarVisorus
	EPAISCTEDetalleCuentasPorCobrarVisorus
	ECPCTEDetalleCuentasPorCobrarVisorus
	ENUMMONDetalleCuentasPorCobrarVisorus
	EFORPAG2DetalleCuentasPorCobrarVisorus
	ENUMEYSDetalleCuentasPorCobrarVisorus
	ESTSCONDetalleCuentasPorCobrarVisorus
	ENUMCFD_1DetalleCuentasPorCobrarVisorus
	ECODIGODetalleCuentasPorCobrarVisorus
	EDESCRIPCIONDetalleCuentasPorCobrarVisorus
	EUNIDADDetalleCuentasPorCobrarVisorus
	ECANTIDADDetalleCuentasPorCobrarVisorus
	EPRECIODetalleCuentasPorCobrarVisorus
	EPRECIOIVAINDetalleCuentasPorCobrarVisorus
	EDESCUENTO_1DetalleCuentasPorCobrarVisorus
	EIMPORTEDetalleCuentasPorCobrarVisorus
	ETASADetalleCuentasPorCobrarVisorus
	EIEPS_1DetalleCuentasPorCobrarVisorus
	EVIEPSDetalleCuentasPorCobrarVisorus
	ENUMMOVCFDDetalleCuentasPorCobrarVisorus
	ENUMDTMCFDDetalleCuentasPorCobrarVisorus
	ECANTIDADMDetalleCuentasPorCobrarVisorus
	ECOSTODetalleCuentasPorCobrarVisorus
}

//SSesion estructura de variables de sesion de Usuarios del sistema
type SSesion struct {
	Name          string
	MenuPrincipal template.HTML
	MenuUsr       template.HTML
}

//SIndex estructura de variables de index
type SIndex struct {
	SResultados        bool
	SRMsj              string
	STituloTabla       string
	SUrlDeDatos        string
	SNombresDeColumnas []string
	SModeloDeColumnas  []map[string]interface{}
	SRenglones         map[string]interface{}
}

//SDetalleCuentasPorCobrarVisorus estructura de  para la vista
type SDetalleCuentasPorCobrarVisorus struct {
	SEstado bool
	SMsj    string
	DetalleCuentasPorCobrarVisorus
	SIndex
	SSesion
}

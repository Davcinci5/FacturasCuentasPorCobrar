package CuentasPorCobrarModel

import (
	"html/template"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//#########################< ESTRUCTURAS >##############################

//E    NUMCFD CuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENUMCFDCuentasPorCobrar struct {
	NUMCFD   int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    SERIE CuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ESERIECuentasPorCobrar struct {
	SERIE    string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    FOLIOCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EFOLIOCuentasPorCobrar struct {
	FOLIO    int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    FECHACuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EFECHACuentasPorCobrar struct {
	FECHA    time.Time
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    NUMCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENUMCTECuentasPorCobrar struct {
	NUMCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    RFCCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ERFCCuentasPorCobrar struct {
	RFC      string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    SUBTOTALCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ESUBTOTALCuentasPorCobrar struct {
	SUBTOTAL float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    DESCUENTOCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EDESCUENTOCuentasPorCobrar struct {
	DESCUENTO float64
	IEstatus  bool
	IMsj      string
	Ihtml     template.HTML
}

//E    IMPUESTOCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EIMPUESTOCuentasPorCobrar struct {
	IMPUESTO float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    IEPSCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EIEPSCuentasPorCobrar struct {
	IEPS     float64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    TIPOCOMPCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ETIPOCOMPCuentasPorCobrar struct {
	TIPOCOMP string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    STATUSCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ESTATUSCuentasPorCobrar struct {
	STATUS   int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    FORPAGCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EFORPAGCuentasPorCobrar struct {
	FORPAG   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    HORACuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EHORACuentasPorCobrar struct {
	HORA     time.Time
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    NUMALMCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENUMALMCuentasPorCobrar struct {
	NUMALM   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    STATUS2CuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ESTATUS2CuentasPorCobrar struct {
	STATUS2  int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    TYPECFDCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ETYPECFDCuentasPorCobrar struct {
	TYPECFD  int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    NOMCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENOMCTECuentasPorCobrar struct {
	NOMCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    DIRCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EDIRCTECuentasPorCobrar struct {
	DIRCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    NLICTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENLICTECuentasPorCobrar struct {
	NLICTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    NLECTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENLECTECuentasPorCobrar struct {
	NLECTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    COLCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ECOLCTECuentasPorCobrar struct {
	COLCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    POBCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EPOBCTECuentasPorCobrar struct {
	POBCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    EDOCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EEDOCTECuentasPorCobrar struct {
	EDOCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    MUNCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EMUNCTECuentasPorCobrar struct {
	MUNCTE   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    PAISCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EPAISCTECuentasPorCobrar struct {
	PAISCTE  string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    CPCTECuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ECPCTECuentasPorCobrar struct {
	CPCTE    string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    NUMMONCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENUMMONCuentasPorCobrar struct {
	NUMMON   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    FORPAG2CuentasPorCobrar Estructura de campo de CuentasPorCobrar
type EFORPAG2CuentasPorCobrar struct {
	FORPAG2  string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    NUMEYSCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ENUMEYSCuentasPorCobrar struct {
	NUMEYS   string
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//E    STSCONCuentasPorCobrar Estructura de campo de CuentasPorCobrar
type ESTSCONCuentasPorCobrar struct {
	STSCON   int64
	IEstatus bool
	IMsj     string
	Ihtml    template.HTML
}

//CuentasPorCobrar estructura de CuentasPorCobrars mongo
type CuentasPorCobrar struct {
	ID bson.ObjectId
	ENUMCFDCuentasPorCobrar
	ESERIECuentasPorCobrar
	EFOLIOCuentasPorCobrar
	EFECHACuentasPorCobrar
	ENUMCTECuentasPorCobrar
	ERFCCuentasPorCobrar
	ESUBTOTALCuentasPorCobrar
	EDESCUENTOCuentasPorCobrar
	EIMPUESTOCuentasPorCobrar
	EIEPSCuentasPorCobrar
	ETIPOCOMPCuentasPorCobrar
	ESTATUSCuentasPorCobrar
	EFORPAGCuentasPorCobrar
	EHORACuentasPorCobrar
	ENUMALMCuentasPorCobrar
	ESTATUS2CuentasPorCobrar
	ETYPECFDCuentasPorCobrar
	ENOMCTECuentasPorCobrar
	EDIRCTECuentasPorCobrar
	ENLICTECuentasPorCobrar
	ENLECTECuentasPorCobrar
	ECOLCTECuentasPorCobrar
	EPOBCTECuentasPorCobrar
	EEDOCTECuentasPorCobrar
	EMUNCTECuentasPorCobrar
	EPAISCTECuentasPorCobrar
	ECPCTECuentasPorCobrar
	ENUMMONCuentasPorCobrar
	EFORPAG2CuentasPorCobrar
	ENUMEYSCuentasPorCobrar
	ESTSCONCuentasPorCobrar
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

//SCuentasPorCobrar estructura de  para la vista
type SCuentasPorCobrar struct {
	SEstado bool
	SMsj    string
	CuentasPorCobrar
	SIndex
	SSesion
}

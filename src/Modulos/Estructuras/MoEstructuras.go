package MoEstructuras

import (
	"html/template"
	"time"
)

//#############################################################################

//Send estructura que contiene documento y mac
type Send struct {
	X1X2X3 []byte
	Y1Y2Y3 []byte
}

//#############################################################################

//#############################################################################

//DocAsk estructura de documento para operaciones que requieren sólo datos para consultar algo.
type DocAsk struct {
	Header Head
	Body   BodyAsk
}

//DocAnswer estructura de documento para operaciones que requieren sólo contestar.
type DocAnswer struct {
	Header Head
	Body   BodyAnswer
}

//DocOperation estructura de documento para operaciones que requieren validar los campos.
type DocOperation struct {
	Header Head
	Body   BodyOpe
}

//DocQuery estructura de documento para consulta de datos simples sin Validación.
type DocQuery struct {
	Header Head
	Body   BodyQuery
}

//#############################################################################

//Head estructura que contiene los datos de cabecera de los documentos.
type Head struct {
	From    string
	Process string
	Date    time.Time
	Status  bool
	Msg     string
}

//#############################################################################

//BodyAsk estructura para el cuerpo del documento para datos de envío de Preguntas.
type BodyAsk struct {
	When    time.Time
	Content map[string]interface{}
}

//BodyAnswer estructura para el cuerpo del documento para datos de respuesta de Preguntas.
type BodyAnswer struct {
	When   time.Time
	Answer bool
	Msg    string
	Aux    template.HTML
}

//BodyOpe estructura para el cuerpo del documento para datos con Validación.
type BodyOpe struct {
	When    time.Time
	Content []Type
}

//BodyQuery estructura para el cuerpo del documento para datos sin Validación.
type BodyQuery struct {
	When    time.Time
	Content interface{}
}

//#############################################################################

//Type estructura para los campos que llevarán validación.
type Type struct {
	ID     string
	Fields []Field
}

//Field estructura para los atributos de los campos que llevarán validación.
type Field struct {
	FieldData interface{}
	Name      string
	Type      string
	Status    string
	Msg       string
	Aux       template.HTML
}

//#############################################################################

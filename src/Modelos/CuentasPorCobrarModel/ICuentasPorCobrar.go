package CuentasPorCobrarModel

import (
	"fmt"

	"../../Modulos/Conexiones"
	"../../Modulos/Variables"
	"gopkg.in/mgo.v2/bson"
)

//ICuentasPorCobrar interface con los métodos de la clase
type ICuentasPorCobrar interface {
	InsertaMgo() bool
	InsertaElastic() bool

	ActualizaMgo(campos []string, valores []interface{}) bool
	ActualizaElastic(campos []string, valores []interface{}) bool //Reemplaza No Actualiza

	ReemplazaMgo() bool
	ReemplazaElastic() bool

	ConsultaExistenciaByFieldMgo(field string, valor string)

	ConsultaExistenciaByIDMgo() bool
	ConsultaExistenciaByIDElastic() bool

	EliminaByIDMgo() bool
	EliminaByIDElastic() bool
}

//################################################<<METODOS DE GESTION >>################################################################

//##################################<< INSERTAR >>###################################

//InsertaMgo es un método que crea un registro en Mongo
func (p CuentasPorCobrarMgo) InsertaMgo() bool {
	result := false
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}

	err = CuentasPorCobrars.Insert(p)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}

	s.Close()
	return result
}

//InsertaElastic es un método que crea un registro en Mongo
func (p CuentasPorCobrarMgo) InsertaElastic() bool {
	var CuentasPorCobrarE CuentasPorCobrarElastic

	CuentasPorCobrarE.NUMCFD = p.NUMCFD
	CuentasPorCobrarE.SERIE = p.SERIE
	CuentasPorCobrarE.FOLIO = p.FOLIO
	CuentasPorCobrarE.FECHA = p.FECHA
	CuentasPorCobrarE.NUMCTE = p.NUMCTE
	CuentasPorCobrarE.RFC = p.RFC
	CuentasPorCobrarE.SUBTOTAL = p.SUBTOTAL
	CuentasPorCobrarE.DESCUENTO = p.DESCUENTO
	CuentasPorCobrarE.IMPUESTO = p.IMPUESTO
	CuentasPorCobrarE.IEPS = p.IEPS
	CuentasPorCobrarE.TIPOCOMP = p.TIPOCOMP
	CuentasPorCobrarE.STATUS = p.STATUS
	CuentasPorCobrarE.FORPAG = p.FORPAG
	CuentasPorCobrarE.HORA = p.HORA
	CuentasPorCobrarE.NUMALM = p.NUMALM
	CuentasPorCobrarE.STATUS2 = p.STATUS2
	CuentasPorCobrarE.TYPECFD = p.TYPECFD
	CuentasPorCobrarE.NOMCTE = p.NOMCTE
	CuentasPorCobrarE.DIRCTE = p.DIRCTE
	CuentasPorCobrarE.NLICTE = p.NLICTE
	CuentasPorCobrarE.NLECTE = p.NLECTE
	CuentasPorCobrarE.COLCTE = p.COLCTE
	CuentasPorCobrarE.POBCTE = p.POBCTE
	CuentasPorCobrarE.EDOCTE = p.EDOCTE
	CuentasPorCobrarE.MUNCTE = p.MUNCTE
	CuentasPorCobrarE.PAISCTE = p.PAISCTE
	CuentasPorCobrarE.CPCTE = p.CPCTE
	CuentasPorCobrarE.NUMMON = p.NUMMON
	CuentasPorCobrarE.FORPAG2 = p.FORPAG2
	CuentasPorCobrarE.NUMEYS = p.NUMEYS
	CuentasPorCobrarE.STSCON = p.STSCON
	insert := MoConexion.InsertaElastic(MoVar.TipoCuentasPorCobrar, p.ID.Hex(), CuentasPorCobrarE)
	if !insert {
		fmt.Println("Error al insertar CuentasPorCobrar en Elastic")
		return false
	}
	return true
}

//##########################<< UPDATE >>############################################

//ActualizaMgo es un método que encuentra y Actualiza un registro en Mongo
//IMPORTANTE --> Debe coincidir el número y orden de campos con el de valores
func (p CuentasPorCobrarMgo) ActualizaMgo(campos []string, valores []interface{}) bool {
	result := false
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	var Abson bson.M
	Abson = make(map[string]interface{})
	for k, v := range campos {
		Abson[v] = valores[k]
	}
	change := bson.M{"$set": Abson}
	err = CuentasPorCobrars.Update(bson.M{"_id": p.ID}, change)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}
	s.Close()
	return result
}

//ActualizaElastic es un método que encuentra y Actualiza un registro en Mongo
func (p CuentasPorCobrarMgo) ActualizaElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoCuentasPorCobrar, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar CuentasPorCobrar en Elastic")
		return false
	}

	if !p.InsertaElastic() {
		fmt.Println("Error al actualizar CuentasPorCobrar en Elastic, se perdió Referencia.")
		return false
	}

	return true
}

//##########################<< REEMPLAZA >>############################################

//ReemplazaMgo es un método que encuentra y Actualiza un registro en Mongo
func (p CuentasPorCobrarMgo) ReemplazaMgo() bool {
	result := false
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	err = CuentasPorCobrars.Update(bson.M{"_id": p.ID}, p)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}
	s.Close()
	return result
}

//ReemplazaElastic es un método que encuentra y reemplaza un CuentasPorCobrar en elastic
func (p CuentasPorCobrarMgo) ReemplazaElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoCuentasPorCobrar, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar CuentasPorCobrar en Elastic")
		return false
	}
	insert := MoConexion.InsertaElastic(MoVar.TipoCuentasPorCobrar, p.ID.Hex(), p)
	if !insert {
		fmt.Println("Error al actualizar CuentasPorCobrar en Elastic")
		return false
	}
	return true
}

//###########################<< CONSULTA EXISTENCIAS >>###################################

//ConsultaExistenciaByFieldMgo es un método que verifica si un registro existe en Mongo indicando un campo y un valor string
func (p CuentasPorCobrarMgo) ConsultaExistenciaByFieldMgo(field string, valor string) bool {
	result := false
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	n, e := CuentasPorCobrars.Find(bson.M{field: valor}).Count()
	if e != nil {
		fmt.Println(e)
	}
	if n > 0 {
		result = true
	}
	s.Close()
	return result
}

//ConsultaExistenciaByIDMgo es un método que encuentra un registro en Mongo buscándolo por ID
func (p CuentasPorCobrarMgo) ConsultaExistenciaByIDMgo() bool {
	result := false
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	n, e := CuentasPorCobrars.Find(bson.M{"_id": p.ID}).Count()
	if e != nil {
		fmt.Println(e)
	}
	if n > 0 {
		result = true
	}
	s.Close()
	return result
}

//ConsultaExistenciaByIDElastic es un método que encuentra un registro en Mongo buscándolo por ID
func (p CuentasPorCobrarMgo) ConsultaExistenciaByIDElastic() bool {
	result := MoConexion.ConsultaElastic(MoVar.TipoCuentasPorCobrar, p.ID.Hex())
	return result
}

//##################################<< ELIMINACIONES >>#################################################

//EliminaByIDMgo es un método que elimina un registro en Mongo
func (p CuentasPorCobrarMgo) EliminaByIDMgo() bool {
	result := false
	s, CuentasPorCobrars, err := MoConexion.GetColectionMgo(MoVar.ColeccionCuentasPorCobrar)
	if err != nil {
		fmt.Println(err)
	}
	e := CuentasPorCobrars.RemoveId(bson.M{"_id": p.ID})
	if e != nil {
		result = true
	} else {
		fmt.Println(e)
	}
	s.Close()
	return result
}

//EliminaByIDElastic es un método que elimina un registro en Mongo
func (p CuentasPorCobrarMgo) EliminaByIDElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoCuentasPorCobrar, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar CuentasPorCobrar en Elastic")
		return false
	}
	return true
}

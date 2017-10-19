package DetalleCuentasPorCobrarVisorusModel

import (
	"fmt"

	"../../Modulos/Conexiones"
	"../../Modulos/Variables"
	"gopkg.in/mgo.v2/bson"
)

//IDetalleCuentasPorCobrarVisorus interface con los métodos de la clase
type IDetalleCuentasPorCobrarVisorus interface {
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
func (p DetalleCuentasPorCobrarVisorusMgo) InsertaMgo() bool {
	result := false
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}

	err = DetalleCuentasPorCobrarVisoruss.Insert(p)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}

	s.Close()
	return result
}

//InsertaElastic es un método que crea un registro en Mongo
func (p DetalleCuentasPorCobrarVisorusMgo) InsertaElastic() bool {
	var DetalleCuentasPorCobrarVisorusE DetalleCuentasPorCobrarVisorusElastic

	DetalleCuentasPorCobrarVisorusE.NUMCFD = p.NUMCFD
	DetalleCuentasPorCobrarVisorusE.CODIGO = p.CODIGO
	DetalleCuentasPorCobrarVisorusE.DESCRIPCION = p.DESCRIPCION
	DetalleCuentasPorCobrarVisorusE.UNIDAD = p.UNIDAD
	DetalleCuentasPorCobrarVisorusE.CANTIDAD = p.CANTIDAD
	DetalleCuentasPorCobrarVisorusE.PRECIO = p.PRECIO
	DetalleCuentasPorCobrarVisorusE.PRECIOIVAIN = p.PRECIOIVAIN
	DetalleCuentasPorCobrarVisorusE.DESCUENTO = p.DESCUENTO
	DetalleCuentasPorCobrarVisorusE.IMPORTE = p.IMPORTE
	DetalleCuentasPorCobrarVisorusE.TASA = p.TASA
	DetalleCuentasPorCobrarVisorusE.IEPS = p.IEPS
	DetalleCuentasPorCobrarVisorusE.VIEPS = p.VIEPS
	DetalleCuentasPorCobrarVisorusE.NUMMOVCFD = p.NUMMOVCFD
	DetalleCuentasPorCobrarVisorusE.NUMDTMCFD = p.NUMDTMCFD
	DetalleCuentasPorCobrarVisorusE.CANTIDADM = p.CANTIDADM
	DetalleCuentasPorCobrarVisorusE.COSTO = p.COSTO
	insert := MoConexion.InsertaElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, p.ID.Hex(), DetalleCuentasPorCobrarVisorusE)
	if !insert {
		fmt.Println("Error al insertar DetalleCuentasPorCobrarVisorus en Elastic")
		return false
	}
	return true
}

//##########################<< UPDATE >>############################################

//ActualizaMgo es un método que encuentra y Actualiza un registro en Mongo
//IMPORTANTE --> Debe coincidir el número y orden de campos con el de valores
func (p DetalleCuentasPorCobrarVisorusMgo) ActualizaMgo(campos []string, valores []interface{}) bool {
	result := false
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	var Abson bson.M
	Abson = make(map[string]interface{})
	for k, v := range campos {
		Abson[v] = valores[k]
	}
	change := bson.M{"$set": Abson}
	err = DetalleCuentasPorCobrarVisoruss.Update(bson.M{"_id": p.ID}, change)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}
	s.Close()
	return result
}

//ActualizaElastic es un método que encuentra y Actualiza un registro en Mongo
func (p DetalleCuentasPorCobrarVisorusMgo) ActualizaElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar DetalleCuentasPorCobrar en Elastic")
		return false
	}

	if !p.InsertaElastic() {
		fmt.Println("Error al actualizar DetalleCuentasPorCobrar en Elastic, se perdió Referencia.")
		return false
	}

	return true
}

//##########################<< REEMPLAZA >>############################################

//ReemplazaMgo es un método que encuentra y Actualiza un registro en Mongo
func (p DetalleCuentasPorCobrarVisorusMgo) ReemplazaMgo() bool {
	result := false
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	err = DetalleCuentasPorCobrarVisoruss.Update(bson.M{"_id": p.ID}, p)
	if err != nil {
		fmt.Println(err)
	} else {
		result = true
	}
	s.Close()
	return result
}

//ReemplazaElastic es un método que encuentra y reemplaza un DetalleCuentasPorCobrarVisorus en elastic
func (p DetalleCuentasPorCobrarVisorusMgo) ReemplazaElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar DetalleCuentasPorCobrar en Elastic")
		return false
	}
	insert := MoConexion.InsertaElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, p.ID.Hex(), p)
	if !insert {
		fmt.Println("Error al actualizar DetalleCuentasPorCobrar en Elastic")
		return false
	}
	return true
}

//###########################<< CONSULTA EXISTENCIAS >>###################################

//ConsultaExistenciaByFieldMgo es un método que verifica si un registro existe en Mongo indicando un campo y un valor string
func (p DetalleCuentasPorCobrarVisorusMgo) ConsultaExistenciaByFieldMgo(field string, valor string) bool {
	result := false
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}
	n, e := DetalleCuentasPorCobrarVisoruss.Find(bson.M{field: valor}).Count()
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
func (p DetalleCuentasPorCobrarVisorusMgo) ConsultaExistenciaByIDMgo() bool {
	result := false
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}
	n, e := DetalleCuentasPorCobrarVisoruss.Find(bson.M{"_id": p.ID}).Count()
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
func (p DetalleCuentasPorCobrarVisorusMgo) ConsultaExistenciaByIDElastic() bool {
	result := MoConexion.ConsultaElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, p.ID.Hex())
	return result
}

//##################################<< ELIMINACIONES >>#################################################

//EliminaByIDMgo es un método que elimina un registro en Mongo
func (p DetalleCuentasPorCobrarVisorusMgo) EliminaByIDMgo() bool {
	result := false
	s, DetalleCuentasPorCobrarVisoruss, err := MoConexion.GetColectionMgo(MoVar.ColeccionDetalleCuentasPorCobrarVisorus)
	if err != nil {
		fmt.Println(err)
	}
	e := DetalleCuentasPorCobrarVisoruss.RemoveId(bson.M{"_id": p.ID})
	if e != nil {
		result = true
	} else {
		fmt.Println(e)
	}
	s.Close()
	return result
}

//EliminaByIDElastic es un método que elimina un registro en Mongo
func (p DetalleCuentasPorCobrarVisorusMgo) EliminaByIDElastic() bool {
	delete := MoConexion.DeleteElastic(MoVar.TipoDetalleCuentasPorCobrarVisorus, p.ID.Hex())
	if !delete {
		fmt.Println("Error al actualizar DetalleCuentasPorCobrar en Elastic")
		return false
	}
	return true
}

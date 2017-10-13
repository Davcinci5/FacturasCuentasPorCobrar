package MoConexion

//Se realiza las conexiones a mongodb
//Tiene los metodos crear, obtener e inicializar las sesiones
//Los almacena en la variable session
//---Autor: Ramón Cruz Juárez
//---FechaCreacion: 13/09/2017

//###########################################################################################################
//##################################################< IMPORTACIONES >########################################
//###########################################################################################################

import (
	"log"

	"gopkg.in/mgo.v2"
)

//###########################################################################################################
//##################################################< VARIABLES >###########################################
//###########################################################################################################

//session Contiene la referencia a una sesion de mongodb
var session *mgo.Session

//###########################################################################################################
//##################################################< ESTRUCTURAS >##########################################
//###########################################################################################################

//###########################################################################################################
//##################################################< INTERFACES >###########################################
//###########################################################################################################

//###########################################################################################################
//##################################################< FUNCIONES PUBLICAS MONGO >#############################
//###########################################################################################################

/*
InitData Inicializa una sesion de MongoDb
Requiere ningun objeto
Regresa ningun objeto

Autor: Ramón Cruz Juárez
FechaCreacion: 13/09/2017
FechaModificacion: --- : ---
Modificacion : --
*/
func InitData() {
	createDBSession()
}

//###########################################################################################################
//##################################################< FUNCIONES PRIVADAS MONGO >#############################
//###########################################################################################################

/*
createDBSession Crea una sesion de mongodb y la asigna a la variable session
Requiere ningun objeto
Regresa ningun objeto

Autor: Ramón Cruz Juárez
FechaCreacion: 13/09/2017
FechaModificacion: --- : ---
Modificacion : --
*/
func createDBSession() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
}

/*
getSession Obtiene una sesión de MongoDb
Requiere ningun objeto
Regresa la sesión existente (conexion a Mongo)

Autor: Ramón Cruz Juárez
FechaCreacion: 13/09/2017
FechaModificacion: --- : ---
Modificacion : --
*/
func getSession() *mgo.Session {
	if session == nil {
		createDBSession()
	}
	return session
}

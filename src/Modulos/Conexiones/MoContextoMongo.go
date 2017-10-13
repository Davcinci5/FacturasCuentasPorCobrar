package MoConexion

//Establecece un contexto que tendrá la sesion de MongoDb a utilizar
//---Autor: Ramón Cruz Juárez
//---FechaCreacion: 13/09/2017

//###########################################################################################################
//##################################################< IMPORTACIONES >########################################
//###########################################################################################################

import (
	mgo "gopkg.in/mgo.v2"
)

//###########################################################################################################
//##################################################< VARIABLES >###########################################
//###########################################################################################################

//###########################################################################################################
//##################################################< ESTRUCTURAS >##########################################
//###########################################################################################################

//Contexto Contiene la referencia a una session de mongo
type Contexto struct {
	Sesion *mgo.Session
}

/*
AlmacenContexto estructura que contendrá una colección y un contexto a usar para una sesion

Autor: Ramón Cruz Juárez
FechaCreacion: 14/06/2017
FechaModificacion: --- : ---
Modificacion : --
*/
type AlmacenContexto struct {
	Coleccion *mgo.Collection
	Contexto  *Contexto
}

//###########################################################################################################
//##################################################< INTERFACES >###########################################
//###########################################################################################################

//###########################################################################################################
//##################################################< FUNCIONES PUBLICAS MONGO >#############################
//###########################################################################################################

/*
Cerrar Accede al contexto y cierra la sesion del mismo
Requiere ningún objeto
Regresa ningún objeto

Autor: Ramón Cruz Juárez
FechaCreacion: 13/06/2017
FechaModificacion: --- : ---
Modificacion : --
*/
func (c *Contexto) Cerrar() {
	c.Sesion.Close()
}

/*
ObtenerColeccion Regresa el contenido de una coleccion
Requiere el nombre de la base de datos y el nombre de la coleccion
Regresa el contenido de la colección solicitada

Autor: Ramón Cruz Juárez
FechaCreacion: 13/09/2017
FechaModificacion: --- : ---
Modificacion : --
*/
func (c *Contexto) ObtenerColeccion(nombreBD, nombreColeccion string) *mgo.Collection {
	return c.Sesion.DB(nombreBD).C(nombreColeccion)
}

/*
NuevoContexto Obtiene la copia de una sesion y lo regresa en un contexto.
Requiere ningun objeto.
Regresa la referencia a la sesión del contexto a utilizar.

Autor: Ramón Cruz Juárez
FechaCreacion: 13/09/2017
FechaModificacion: --- : ---
Modificacion : --
*/
func NuevoContexto() *Contexto {
	contexto := &Contexto{
		Sesion: getSession().Copy(),
	}
	return contexto
}

//###########################################################################################################
//##################################################< FUNCIONES PRIVADAS MONGO >#############################
//###########################################################################################################

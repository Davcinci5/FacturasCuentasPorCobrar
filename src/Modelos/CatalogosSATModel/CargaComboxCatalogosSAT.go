package CatalogosSATModel

//CargaComboUnidades recibe la clave del catálogo, el identificador opcional
//y regresa el template del combo del catálogo con el identificador seleccionado si así se desea
func CargaComboUnidades(Clave string) string {

	templ := ``

	if Clave != "" {
		templ = `<option value="">--SELECCIONE--</option>`
	} else {
		templ = `<option value="" selected>--SELECCIONE--</option>`
	}

	Unidades := GetAllUnidades()
	for _, v := range Unidades {
		if Clave == v.Clave {
			templ += `<option value="` + v.Clave + `" selected>` + v.Nombre + `</option>`
		} else {
			templ += `<option value="` + v.ID.Hex() + `">` + v.Nombre + `</option>`
		}
	}
	return templ
}

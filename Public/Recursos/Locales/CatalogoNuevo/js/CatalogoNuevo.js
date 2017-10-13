//##############################< SCRIPTS JS >##########################################
//################################< CatalogoNuevo.js >#####################################
//#########################< VALIDACIONES DE JEQUERY >##################################

$(document).ready(function() {
    if (document.getElementById("tbody_elementos_catalogo").children.length == 0) {
        $('#div_tabla_catalogo').hide();
    }
    var defaultActivo=false;
    var validator = valida();

$( "#Valor" ).focus(function() {
  var clv = $("#Clave").val();  
  if (numeroElementos==0){
      if (rebasado==0){
        $("#Subclave").val(clv+"A");
      }else{
        $("#Subclave").val(clv+"A"+rebasado);
      }
  }else{
      letraUsada = nextChar(abecedario[numeroElementos-1]);
      if (rebasado==0){
        $("#Subclave").val(clv+letraUsada);
      }else{
        $("#Subclave").val(clv+letraUsada+rebasado);          
      }
      if (letraUsada=="Z"){
            rebasado++;
            numeroElementos=-1;
      }
  }
});

var numeroElementos=0;
var rebasado = 0;
function nextChar(c) {    
    return String.fromCharCode(c.charCodeAt(0) + 1);    
}
var abecedario = ['A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z']
var letraUsada='';
    $('#AgregaCampo').click(function() {
        numeroElementos++;
        var errores = 0;
        var valor=$('#Valor').val();
        var subclave=$('#Subclave').val();
        var valorEditable = false;
        if($("#ValorEditable").prop('checked') == true){
           var valorEditable = true;
        }
        if (valor == ""){
            errores ++;
            validator.showErrors({
                "Valor": "No puede agregar valores vacíos"
            });
            $("#Valor").focus();            
        }
        if ( subclave== "" && errores==0){
            validator.showErrors({
                "Subclave": "No puede agregar valores vacíos"
            });
            $("#Subclave").focus();             
        }

        if (errores==0) {
            $('#div_tabla_catalogo').show();
            var html= '<td><input type="hidden" class="form-control" name="ValoresIds" value="">'+
                '<input type="text" class="form-control" onblur="ValidaCampo(this)"  id="inputTX" name="Valores" value="' + valor + '" readonly></td>'+
                '<td><input type="text" class="form-control" name="subclave" value="' + subclave + '" readonly></td>';
            if (valorEditable){
                html += '<td><input type="checkbox" class="form-control" name="ValorEditable" value="' + subclave + '" readonly checked ></td>';
            }else{
                html += '<td><input type="checkbox" class="form-control" name="ValorEditable" value="' + subclave + '" readonly></td>';                
            }
            if (defaultActivo){
                html+= '<td><input type="radio" class="form-control" name="ValorDefault" value="' + subclave + '" readonly></td>';
            }else{
                defaultActivo = true;
                html += '<td><input type="radio" class="form-control" name="ValorDefault" value="' + subclave + '" readonly checked></td>';
            }
            html += '<td><SELECT id="EstatusElemento'+subclave+'" class="form-control" name="EstatusElemento" value="EstatusElemento" readonly> </SELECT></td>';
            //html += '<td><SELECT id="EstatusElemento2.1" class="form-control" name="EstatusElemento" value="EstatusElemento" readonly> </SELECT></td>';
            //html += '<td id="EstatusElemento'+subclave+'"></td>';            
            html += '<td><button type="button" class="btn btn-primary editButton"><span class="glyphicon glyphicon-pencil btn-xs"></span></button></td>'+
                '<td><button type="button" class="btn btn-danger deleteButton"><span class="glyphicon glyphicon-trash btn-xs"></span></button></td></td>';                
            $("#tbody_elementos_catalogo").append(
                '<tr>'+html+'</tr>'
                );
            //$('#EstatusCatalogo').find('option').clone().appendTo('#EstatusElemento'+subclave);
            var subClaveArray = subclave.split(".");
            if (subClaveArray.length>1){
                $('#EstatusCatalogo').find('option').clone().appendTo('#EstatusElemento'+subClaveArray[0]+'\\.'+subClaveArray[1]);
                //$('select#EstatusCatalogo').clone().attr('id', 'newOptions').appendTo('#EstatusElemento'+subclave);            
            }else{
                $('#EstatusCatalogo').find('option').clone().appendTo('#EstatusElemento'+subclave);
            }
            $("#Valor").val("");
            $('#Subclave').val("");
            $("#Valor").focus(); 
        }
    });

	$('#Form_Alta_CatalogoNuevo').keydown(function(e) {
		if(e.which == 13 || e.keyCode == 13) {
			if ($('#Valor').val() != ""){
				e.preventDefault();
				$('#AgregaCampo').trigger("click");
				validator.element("#Valor");
			}else{
				if (document.getElementById("tbody_etiquetas_catalogo").children.length == 0){
					e.preventDefault();
					validator.showErrors({
						"Valor": "No puede dar de alta Catálogos sin valores"
					});
					
				}else if($("#Nombre")==''){
					validator.showErrors({
						"Nombre": "El Nombre es requerido"
					});
					$("#Nombre").focus();
				}
			}  
		}
	});    
});

function valida() {
    var validator = $("#Form_Alta_CatalogoNuevo").validate({
        rules: {
            Clave: {
                required: true, rangelength: [1, 4]
            },
            Nombre: {
                required: true, rangelength: [5, 50]
            },
            Descripcion: {
                required: true, rangelength: [15, 150]
            },
            CatalogoEditable: {
                required: false
            },
            SubClave: {
                required: true, rangelength: [1, 10]
            },
            Valor: {
                required: false, rangelength: [5, 20]
            },
            ValorEditable: {
                required: false
            },
        },
        messages: {
            Clave: {
                required: "El campo Clave es requerido.",
                range: "el valor del campo Clave debe estar entre  [1, 4]"
            },
            Nombre: {
                required: "El campo Nombre es requerido.",
                rangelength: "La longitud del campo Nombre debe estar entre  [5, 50]"
            },
            Descripcion: {
                required: "El campo Descripcion es requerido.",
                rangelength: "La longitud del campo Descripcion debe estar entre  [15, 150]"
            },
            CatalogoEditable: {
                required: "El campo CatalogoEditable es requerido."
            },
            SubClave: {
                required: "El campo SubClave es requerido.",
                range: "el valor del campo SubClave debe estar entre  [1, 4]"
            },
            Valor: {
                required: "El campo Valor es requerido.",
                rangelength: "La longitud del campo Valor debe estar entre  [5, 20]"
            },
            ValorEditable: {
                required: "El campo ValorEditable es requerido."
            },
        },
        errorElement: "em",
        errorPlacement: function(error, element) {
            error.addClass("help-block");
            element.parents(".col-sm-5").addClass("has-feedback");

            if (element.prop("type") === "checkbox") {
                error.insertAfter(element.parent("label"));
            } else {
                error.insertAfter(element);
            }

            if (!element.next("span")[0]) {
                $("<span class='glyphicon glyphicon-remove form-control-feedback'></span>").insertAfter(element);
            }
        },
        success: function(label, element) {
            if (!$(element).next("span")[0]) {
                $("<span class='glyphicon glyphicon-ok form-control-feedback'></span>").insertAfter($(element));
            }
        },
        highlight: function(element, errorClass, validClass) {
            $(element).parents(".col-sm-5").addClass("has-error").removeClass("has-success");
            $(element).next("span").addClass("glyphicon-remove").removeClass("glyphicon-ok");
        },
        unhighlight: function(element, errorClass, validClass) {
            $(element).parents(".col-sm-5").addClass("has-success").removeClass("has-error");
            $(element).next("span").addClass("glyphicon-ok").removeClass("glyphicon-remove");
        }
    });
    return validator;
}

function EditaCatalogoNuevo(vista) {
    if (vista == "Index" || vista == "") {
        if ($('#CatalogoNuevos').val() != "") {
            window.location = '/CatalogoNuevos/edita/' + $('#CatalogoNuevos').val();
        } else {
            alertify.error("Debe Seleccionar un CatalogoNuevo para editar");
        }
    } else if (vista == "Detalle") {
        if ($('#ID').val() != "") {
            window.location = '/CatalogoNuevos/edita/' + $('#ID').val();
        } else {
            alertify.error("No se puede editar debido a un error de referencias, favor de intentar en el index");
            window.location = '/CatalogoNuevos';
        }
    }

}


function DetalleCatalogoNuevo() {
    if ($('#CatalogoNuevos').val() != "") {
        window.location = '/CatalogoNuevos/detalle/' + $('#CatalogoNuevos').val();
    } else {
        alertify.error("Debe Seleccionar un CatalogoNuevo para editar");
    }
}



function BuscaPagina(num) {
    $('#Loading').show();

    $.ajax({
        url: "/CatalogoNuevos/search",
        type: 'POST',
        dataType: 'json',
        data: {
            Pag: num,
        },
        success: function(data) {
            if (data != null) {
                if (data.SEstado) {
                    $("#Cabecera").empty();
                    $("#Cabecera").append(data.SCabecera);
                    $("#Cuerpo").empty();
                    $("#Cuerpo").append(data.SBody);
                    $("#Paginacion").empty();
                    $("#Paginacion").append(data.SPaginacion);
                } else {
                    alertify.error(data.SMsj);
                }
            } else {
                alertify.error("Hubo un problema al recibir información del servidor, favor de volver a intentar.");
            }
            $('#Loading').hide();
        },
        error: function(data) {
            $('#Loading').hide();
        },
    });
}


function SubmitGroup() {
    $('#Loading').show();
    $.ajax({
        url: "/CatalogoNuevos/agrupa",
        type: 'POST',
        dataType: 'json',
        data: {
            Grupox: $('#Grupos').val(),
            searchbox: $('#searchbox').val()
        },
        success: function(data) {
            if (data != null) {
                if (data.SEstado) {
                    $("#Cabecera").empty();
                    $("#Cabecera").append(data.SCabecera);
                    $("#Cuerpo").empty();
                    $("#Cuerpo").append(data.SBody);
                    $("#Paginacion").empty();
                    $("#Paginacion").append(data.SPaginacion);
                } else {
                    alertify.error(data.SMsj);
                }
            } else {
                alertify.error("Hubo un problema al recibir información del servidor, favor de volver a intentar.");
            }
            $('#Loading').hide();
        },
        error: function(data) {
            $('#Loading').hide();
        },
    });
}


	//##############################< SCRIPTS JS >##########################################
	//################################< CuentasPorCobrar.js >#####################################
	//#########################< VALIDACIONES DE JEQUERY >##################################

	$( document ).ready( function () {
		var validator = valida();
	});

    function valida(){
	var validator = $("#Form_Alta_CuentasPorCobrar").validate({
		rules: {
			
			    NUMCFD  : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			    SERIE  : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			    FOLIO : {
						
					rangelength : [1, 100],
				
					required : true
				
					},
			    FECHA : {
						
					required : true
				
					},
			    NUMCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    RFC : {
						
					required : true,
				
					rangelength : [1, 20]
				
					},
			    SUBTOTAL : {
						
					required : true,
				
					range : [1, 20]
				
					},
			    DESCUENTO : {
						
					range : [1, 20],
				
					required : true
				
					},
			    IMPUESTO : {
						
					required : true,
				
					range : [1, 20]
				
					},
			    IEPS : {
						
					required : true,
				
					range : [1, 20]
				
					},
			    TIPOCOMP : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    STATUS : {
						
					required : true,
				
					rangelength : [1, 10]
				
					},
			    FORPAG : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			    HORA : {
						
					required : true
				
					},
			    NUMALM : {
						
					required : true,
				
					rangelength : [0, 50]
				
					},
			    STATUS2 : {
						
					required : true,
				
					rangelength : [1, 10]
				
					},
			    TYPECFD : {
						
					required : true,
				
					rangelength : [1, 20]
				
					},
			    NOMCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    DIRCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    NLICTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    NLECTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    COLCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    POBCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    EDOCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    MUNCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    PAISCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    CPCTE : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    NUMMON : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
			    FORPAG2 : {
						
					rangelength : [1, 50],
				
					required : true
				
					},
			    NUMEYS : {
						
					required : true,
				
					rangelength : [1, 50]
				
					},
				    STSCON : {
						
					required : true,
				
					rangelength : [1, 10]
				
					}
		},
		messages: {
			
			    NUMCFD  : {
						
					required : "El campo     NUMCFD  es requerido.",
					rangelength : "La longitud del campo     NUMCFD  debe estar entre  [1, 100]"
					},
			    SERIE  : {
						
					required : "El campo     SERIE  es requerido.",
					rangelength : "La longitud del campo     SERIE  debe estar entre  [1, 100]"
					},
			    FOLIO : {
						
					required : "El campo     FOLIO es requerido.",
					rangelength : "La longitud del campo     FOLIO debe estar entre  [1, 100]"
					},
			    FECHA : {
						
					required : "El campo     FECHA es requerido."
					},
			    NUMCTE : {
						
					required : "El campo     NUMCTE es requerido.",
					rangelength : "La longitud del campo     NUMCTE debe estar entre  [1, 50]"
					},
			    RFC : {
						
					required : "El campo     RFC es requerido.",
					rangelength : "La longitud del campo     RFC debe estar entre  [1, 20]"
					},
			    SUBTOTAL : {
						
					required : "El campo     SUBTOTAL es requerido.",
					range : "el valor del campo     SUBTOTAL debe estar entre  [1, 20]"
					},
			    DESCUENTO : {
						
					required : "El campo     DESCUENTO es requerido.",
					range : "el valor del campo     DESCUENTO debe estar entre  [1, 20]"
					},
			    IMPUESTO : {
						
					range : "el valor del campo     IMPUESTO debe estar entre  [1, 20]",
					required : "El campo     IMPUESTO es requerido."
					},
			    IEPS : {
						
					required : "El campo     IEPS es requerido.",
					range : "el valor del campo     IEPS debe estar entre  [1, 20]"
					},
			    TIPOCOMP : {
						
					required : "El campo     TIPOCOMP es requerido.",
					rangelength : "La longitud del campo     TIPOCOMP debe estar entre  [1, 50]"
					},
			    STATUS : {
						
					required : "El campo     STATUS es requerido.",
					rangelength : "La longitud del campo     STATUS debe estar entre  [1, 10]"
					},
			    FORPAG : {
						
					required : "El campo     FORPAG es requerido.",
					rangelength : "La longitud del campo     FORPAG debe estar entre  [1, 100]"
					},
			    HORA : {
						
					required : "El campo     HORA es requerido."
					},
			    NUMALM : {
						
					required : "El campo     NUMALM es requerido.",
					rangelength : "La longitud del campo     NUMALM debe estar entre  [0, 50]"
					},
			    STATUS2 : {
						
					required : "El campo     STATUS2 es requerido.",
					rangelength : "La longitud del campo     STATUS2 debe estar entre  [1, 10]"
					},
			    TYPECFD : {
						
					required : "El campo     TYPECFD es requerido.",
					rangelength : "La longitud del campo     TYPECFD debe estar entre  [1, 20]"
					},
			    NOMCTE : {
						
					required : "El campo     NOMCTE es requerido.",
					rangelength : "La longitud del campo     NOMCTE debe estar entre  [1, 50]"
					},
			    DIRCTE : {
						
					required : "El campo     DIRCTE es requerido.",
					rangelength : "La longitud del campo     DIRCTE debe estar entre  [1, 50]"
					},
			    NLICTE : {
						
					required : "El campo     NLICTE es requerido.",
					rangelength : "La longitud del campo     NLICTE debe estar entre  [1, 50]"
					},
			    NLECTE : {
						
					required : "El campo     NLECTE es requerido.",
					rangelength : "La longitud del campo     NLECTE debe estar entre  [1, 50]"
					},
			    COLCTE : {
						
					required : "El campo     COLCTE es requerido.",
					rangelength : "La longitud del campo     COLCTE debe estar entre  [1, 50]"
					},
			    POBCTE : {
						
					rangelength : "La longitud del campo     POBCTE debe estar entre  [1, 50]",
					required : "El campo     POBCTE es requerido."
					},
			    EDOCTE : {
						
					rangelength : "La longitud del campo     EDOCTE debe estar entre  [1, 50]",
					required : "El campo     EDOCTE es requerido."
					},
			    MUNCTE : {
						
					required : "El campo     MUNCTE es requerido.",
					rangelength : "La longitud del campo     MUNCTE debe estar entre  [1, 50]"
					},
			    PAISCTE : {
						
					required : "El campo     PAISCTE es requerido.",
					rangelength : "La longitud del campo     PAISCTE debe estar entre  [1, 50]"
					},
			    CPCTE : {
						
					required : "El campo     CPCTE es requerido.",
					rangelength : "La longitud del campo     CPCTE debe estar entre  [1, 50]"
					},
			    NUMMON : {
						
					required : "El campo     NUMMON es requerido.",
					rangelength : "La longitud del campo     NUMMON debe estar entre  [1, 50]"
					},
			    FORPAG2 : {
						
					required : "El campo     FORPAG2 es requerido.",
					rangelength : "La longitud del campo     FORPAG2 debe estar entre  [1, 50]"
					},
			    NUMEYS : {
						
					rangelength : "La longitud del campo     NUMEYS debe estar entre  [1, 50]",
					required : "El campo     NUMEYS es requerido."
					},
			    STSCON : {
						
					required : "El campo     STSCON es requerido.",
					rangelength : "La longitud del campo     STSCON debe estar entre  [1, 10]"
					}
		},
		errorElement: "em",
		errorPlacement: function ( error, element ) {
			error.addClass( "help-block" );
			element.parents( ".col-sm-5" ).addClass( "has-feedback" );

			if ( element.prop( "type" ) === "checkbox" ) {
				error.insertAfter( element.parent( "label" ) );
			} else {
				error.insertAfter( element );
			}

			if ( !element.next( "span" )[ 0 ] ) {
				$( "<span class='glyphicon glyphicon-remove form-control-feedback'></span>" ).insertAfter( element );
			}
		},
		success: function ( label, element ) {
			if ( !$( element ).next( "span" )[ 0 ] ) {
				$( "<span class='glyphicon glyphicon-ok form-control-feedback'></span>" ).insertAfter( $( element ) );
			}
		},
		highlight: function ( element, errorClass, validClass ) {
			$( element ).parents( ".col-sm-5" ).addClass( "has-error" ).removeClass( "has-success" );
			$( element ).next( "span" ).addClass( "glyphicon-remove" ).removeClass( "glyphicon-ok" );
		},
		unhighlight: function ( element, errorClass, validClass ) {
			$( element ).parents( ".col-sm-5" ).addClass( "has-success" ).removeClass( "has-error" );
			$( element ).next( "span" ).addClass( "glyphicon-ok" ).removeClass( "glyphicon-remove" );
		}
	});
	return validator;
}

function EditaCuentasPorCobrar(vista){
	if (vista == "Index" || vista ==""){
		if ($('#CuentasPorCobrars').val() != ""){
			window.location = '/CuentasPorCobrars/edita/' + $('#CuentasPorCobrars').val();
		}else{
			alertify.error("Debe Seleccionar un CuentasPorCobrar para editar");
		}
	}else if(vista == "Detalle"){
		if ($('#ID').val() != ""){
			window.location = '/CuentasPorCobrars/edita/' + $('#ID').val();
		}else{
			alertify.error("No se puede editar debido a un error de referencias, favor de intentar en el index");
			window.location = '/CuentasPorCobrars';
		}
	}

}


function DetalleCuentasPorCobrar(){
	if ($('#CuentasPorCobrars').val() != ""){
		window.location = '/CuentasPorCobrars/detalle/' + $('#CuentasPorCobrars').val();
	}else{
	alertify.error("Debe Seleccionar un CuentasPorCobrar para editar");
	}
}



function BuscaPagina(num){
			$('#Loading').show();

			$.ajax({
			url:"/CuentasPorCobrars/search",
			type: 'POST',
			dataType:'json',
			data:{
				Pag : num,
			},
			success: function(data){
				if (data != null){
					if (data.SEstado){
						$("#Cabecera").empty();
						$("#Cabecera").append(data.SCabecera);
						$("#Cuerpo").empty();
						$("#Cuerpo").append(data.SBody);
						$("#Paginacion").empty();
						$("#Paginacion").append(data.SPaginacion);
					}else{
						alertify.error(data.SMsj);
					}
				}else{
					alertify.error("Hubo un problema al recibir información del servidor, favor de volver a intentar.");
				}
				$('#Loading').hide();
			},
		  error: function(data){
				$('#Loading').hide();
		  },
		});
}


 function SubmitGroup(){
	 $('#Loading').show();
			$.ajax({
			url:"/CuentasPorCobrars/agrupa",
			type: 'POST',
			dataType:'json',
			data:{
				Grupox : $('#Grupos').val(),
				searchbox: $('#searchbox').val()
			},
			success: function(data){
				if (data != null){
					if (data.SEstado){
						$("#Cabecera").empty();
						$("#Cabecera").append(data.SCabecera);
						$("#Cuerpo").empty();
						$("#Cuerpo").append(data.SBody);
						$("#Paginacion").empty();
						$("#Paginacion").append(data.SPaginacion);
					}else{
						alertify.error(data.SMsj);
					}
				}else{
					alertify.error("Hubo un problema al recibir información del servidor, favor de volver a intentar.");
				}
				$('#Loading').hide();
			},
		  error: function(data){
			  $('#Loading').hide();
		  },
		});
}



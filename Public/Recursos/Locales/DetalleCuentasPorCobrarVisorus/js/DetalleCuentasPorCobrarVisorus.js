

	//##############################< SCRIPTS JS >##########################################
	//################################< DetalleCuentasPorCobrarVisorus.js >#####################################
	//#########################< VALIDACIONES DE JEQUERY >##################################

	$( document ).ready( function () {
		var validator = valida();
	});

    function valida(){
	var validator = $("#Form_Alta_DetalleCuentasPorCobrarVisorus").validate({
		rules: {
			
			NUMCFD  : {
						
					required : true,
				
					rangelength : [1, 1000]
				
					},
			SERIE  : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			FOLIO : {
						
					required : true,
				
					rangelength : [1, 1000]
				
					},
			FECHA : {
						
					required : true
				
					},
			NUMCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			RFC : {
						
					required : true,
				
					rangelength : [1, 20]
				
					},
			SUBTOTAL : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			DESCUENTO : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			IMPUESTO : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			IEPS : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			TIPO_COMP : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			STATUS : {
						
					required : true,
				
					rangelength : [1, 1000]
				
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
				
					rangelength : [0, 100]
				
					},
			STATUS2 : {
						
					required : true,
				
					rangelength : [1, 1000]
				
					},
			TYPECFD : {
						
					required : true,
				
					rangelength : [1, 1000]
				
					},
			NOMCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			DIRCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			NLICTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			NLECTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			COLCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			POBCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			EDOCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			PAISCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			CPCTE : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			NUMMON : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			FORPAG2 : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			NUMEYS : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			STSCON : {
						
					required : true,
				
					rangelength : [1, 1000]
				
					},
			NUMCFD_1 : {
						
					required : true,
				
					rangelength : [1, 1000]
				
					},
			CODIGO : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			DESCRIPCION : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			UNIDAD : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			CANTIDAD : {
						
					range : [1, 1000],
				
					required : true
				
					},
			PRECIO : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			PRECIOIVAIN : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			DESCUENTO_1 : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			IMPORTE : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			TASA : {
						
					rangelength : [1, 100],
				
					required : true
				
					},
			IEPS_1 : {
						
					required : true,
				
					rangelength : [1, 100]
				
					},
			VIEPS : {
						
					required : true,
				
					range : [1, 1000]
				
					},
			NUMMOVCFD : {
						
					required : true,
				
					rangelength : [1, 1000]
				
					},
			NUMDTMCFD : {
						
					rangelength : [1, 1000],
				
					required : true
				
					},
			CANTIDADM : {
						
					required : true,
				
					range : [1, 1000]
				
					},
				COSTO : {
						
					required : true,
				
					range : [1, 1000]
				
					}
		},
		messages: {
			
			NUMCFD  : {
						
					required : "El campo NUMCFD  es requerido.",
					rangelength : "La longitud del campo NUMCFD  debe estar entre  [1, 1000]"
					},
			SERIE  : {
						
					required : "El campo SERIE  es requerido.",
					rangelength : "La longitud del campo SERIE  debe estar entre  [1, 100]"
					},
			FOLIO : {
						
					required : "El campo FOLIO es requerido.",
					rangelength : "La longitud del campo FOLIO debe estar entre  [1, 1000]"
					},
			FECHA : {
						
					required : "El campo FECHA es requerido."
					},
			NUMCTE : {
						
					required : "El campo NUMCTE es requerido.",
					rangelength : "La longitud del campo NUMCTE debe estar entre  [1, 100]"
					},
			RFC : {
						
					required : "El campo RFC es requerido.",
					rangelength : "La longitud del campo RFC debe estar entre  [1, 20]"
					},
			SUBTOTAL : {
						
					required : "El campo SUBTOTAL es requerido.",
					range : "el valor del campo SUBTOTAL debe estar entre  [1, 1000]"
					},
			DESCUENTO : {
						
					required : "El campo DESCUENTO es requerido.",
					range : "el valor del campo DESCUENTO debe estar entre  [1, 1000]"
					},
			IMPUESTO : {
						
					required : "El campo IMPUESTO es requerido.",
					range : "el valor del campo IMPUESTO debe estar entre  [1, 1000]"
					},
			IEPS : {
						
					required : "El campo IEPS es requerido.",
					range : "el valor del campo IEPS debe estar entre  [1, 1000]"
					},
			TIPO_COMP : {
						
					required : "El campo TIPO_COMP es requerido.",
					rangelength : "La longitud del campo TIPO_COMP debe estar entre  [1, 100]"
					},
			STATUS : {
						
					required : "El campo STATUS es requerido.",
					rangelength : "La longitud del campo STATUS debe estar entre  [1, 1000]"
					},
			FORPAG : {
						
					required : "El campo FORPAG es requerido.",
					rangelength : "La longitud del campo FORPAG debe estar entre  [1, 100]"
					},
			HORA : {
						
					required : "El campo HORA es requerido."
					},
			NUMALM : {
						
					required : "El campo NUMALM es requerido.",
					rangelength : "La longitud del campo NUMALM debe estar entre  [0, 100]"
					},
			STATUS2 : {
						
					rangelength : "La longitud del campo STATUS2 debe estar entre  [1, 1000]",
					required : "El campo STATUS2 es requerido."
					},
			TYPECFD : {
						
					required : "El campo TYPECFD es requerido.",
					rangelength : "La longitud del campo TYPECFD debe estar entre  [1, 1000]"
					},
			NOMCTE : {
						
					required : "El campo NOMCTE es requerido.",
					rangelength : "La longitud del campo NOMCTE debe estar entre  [1, 100]"
					},
			DIRCTE : {
						
					required : "El campo DIRCTE es requerido.",
					rangelength : "La longitud del campo DIRCTE debe estar entre  [1, 100]"
					},
			NLICTE : {
						
					required : "El campo NLICTE es requerido.",
					rangelength : "La longitud del campo NLICTE debe estar entre  [1, 100]"
					},
			NLECTE : {
						
					required : "El campo NLECTE es requerido.",
					rangelength : "La longitud del campo NLECTE debe estar entre  [1, 100]"
					},
			COLCTE : {
						
					required : "El campo COLCTE es requerido.",
					rangelength : "La longitud del campo COLCTE debe estar entre  [1, 100]"
					},
			POBCTE : {
						
					required : "El campo POBCTE es requerido.",
					rangelength : "La longitud del campo POBCTE debe estar entre  [1, 100]"
					},
			EDOCTE : {
						
					rangelength : "La longitud del campo EDOCTE debe estar entre  [1, 100]",
					required : "El campo EDOCTE es requerido."
					},
			PAISCTE : {
						
					rangelength : "La longitud del campo PAISCTE debe estar entre  [1, 100]",
					required : "El campo PAISCTE es requerido."
					},
			CPCTE : {
						
					required : "El campo CPCTE es requerido.",
					rangelength : "La longitud del campo CPCTE debe estar entre  [1, 100]"
					},
			NUMMON : {
						
					required : "El campo NUMMON es requerido.",
					rangelength : "La longitud del campo NUMMON debe estar entre  [1, 100]"
					},
			FORPAG2 : {
						
					required : "El campo FORPAG2 es requerido.",
					rangelength : "La longitud del campo FORPAG2 debe estar entre  [1, 100]"
					},
			NUMEYS : {
						
					required : "El campo NUMEYS es requerido.",
					rangelength : "La longitud del campo NUMEYS debe estar entre  [1, 100]"
					},
			STSCON : {
						
					rangelength : "La longitud del campo STSCON debe estar entre  [1, 1000]",
					required : "El campo STSCON es requerido."
					},
			NUMCFD_1 : {
						
					required : "El campo NUMCFD_1 es requerido.",
					rangelength : "La longitud del campo NUMCFD_1 debe estar entre  [1, 1000]"
					},
			CODIGO : {
						
					required : "El campo CODIGO es requerido.",
					rangelength : "La longitud del campo CODIGO debe estar entre  [1, 100]"
					},
			DESCRIPCION : {
						
					required : "El campo DESCRIPCION es requerido.",
					rangelength : "La longitud del campo DESCRIPCION debe estar entre  [1, 100]"
					},
			UNIDAD : {
						
					required : "El campo UNIDAD es requerido.",
					rangelength : "La longitud del campo UNIDAD debe estar entre  [1, 100]"
					},
			CANTIDAD : {
						
					required : "El campo CANTIDAD es requerido.",
					range : "el valor del campo CANTIDAD debe estar entre  [1, 1000]"
					},
			PRECIO : {
						
					required : "El campo PRECIO es requerido.",
					range : "el valor del campo PRECIO debe estar entre  [1, 1000]"
					},
			PRECIOIVAIN : {
						
					required : "El campo PRECIOIVAIN es requerido.",
					range : "el valor del campo PRECIOIVAIN debe estar entre  [1, 1000]"
					},
			DESCUENTO_1 : {
						
					required : "El campo DESCUENTO_1 es requerido.",
					range : "el valor del campo DESCUENTO_1 debe estar entre  [1, 1000]"
					},
			IMPORTE : {
						
					required : "El campo IMPORTE es requerido.",
					range : "el valor del campo IMPORTE debe estar entre  [1, 1000]"
					},
			TASA : {
						
					required : "El campo TASA es requerido.",
					rangelength : "La longitud del campo TASA debe estar entre  [1, 100]"
					},
			IEPS_1 : {
						
					required : "El campo IEPS_1 es requerido.",
					rangelength : "La longitud del campo IEPS_1 debe estar entre  [1, 100]"
					},
			VIEPS : {
						
					required : "El campo VIEPS es requerido.",
					range : "el valor del campo VIEPS debe estar entre  [1, 1000]"
					},
			NUMMOVCFD : {
						
					required : "El campo NUMMOVCFD es requerido.",
					rangelength : "La longitud del campo NUMMOVCFD debe estar entre  [1, 1000]"
					},
			NUMDTMCFD : {
						
					required : "El campo NUMDTMCFD es requerido.",
					rangelength : "La longitud del campo NUMDTMCFD debe estar entre  [1, 1000]"
					},
			CANTIDADM : {
						
					required : "El campo CANTIDADM es requerido.",
					range : "el valor del campo CANTIDADM debe estar entre  [1, 1000]"
					},
			COSTO : {
						
					required : "El campo COSTO es requerido.",
					range : "el valor del campo COSTO debe estar entre  [1, 1000]"
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

function EditaDetalleCuentasPorCobrarVisorus(vista){
	if (vista == "Index" || vista ==""){
		if ($('#DetalleCuentasPorCobrarVisoruss').val() != ""){
			window.location = '/DetalleCuentasPorCobrarVisoruss/edita/' + $('#DetalleCuentasPorCobrarVisoruss').val();
		}else{
			alertify.error("Debe Seleccionar un DetalleCuentasPorCobrarVisorus para editar");
		}
	}else if(vista == "Detalle"){
		if ($('#ID').val() != ""){
			window.location = '/DetalleCuentasPorCobrarVisoruss/edita/' + $('#ID').val();
		}else{
			alertify.error("No se puede editar debido a un error de referencias, favor de intentar en el index");
			window.location = '/DetalleCuentasPorCobrarVisoruss';
		}
	}

}


function DetalleDetalleCuentasPorCobrarVisorus(){
	if ($('#DetalleCuentasPorCobrarVisoruss').val() != ""){
		window.location = '/DetalleCuentasPorCobrarVisoruss/detalle/' + $('#DetalleCuentasPorCobrarVisoruss').val();
	}else{
	alertify.error("Debe Seleccionar un DetalleCuentasPorCobrarVisorus para editar");
	}
}



function BuscaPagina(num){
			$('#Loading').show();

			$.ajax({
			url:"/DetalleCuentasPorCobrarVisoruss/search",
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
			url:"/DetalleCuentasPorCobrarVisoruss/agrupa",
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



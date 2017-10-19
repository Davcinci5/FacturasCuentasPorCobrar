$(document).ready(function() {
    $('#textoabuscar').keydown(function(e) {
        if (e.which == 13 || e.keyCode == 13) {
            e.preventDefault();
            GetEspecificToIndex('DetalleCuentasPorCobrarVisorus');
        }
    });
});

function GetEspecificToIndex(NModelo) {
    var TextoABuscar = $('#textoabuscar').val();
    if (TextoABuscar != '') {
        var url = '/' + NModelo + '/CargaIndexEspecifico'
        $('#Loading').show();
        $.ajax({
            url: url,
            type: 'POST',
            dataType: 'json',
            data: {
                Texto: TextoABuscar
            },
            success: function(data) {
                if (data != null) {
                    if (data.SEstado) {
                        function Modelo() {
                            data.SModeloDeColumnas[3].searchoptions["FECHA"] = datePick;
                            data.SModeloDeColumnas[13].searchoptions["HORA"] = datePick;
                            // data.SModeloDeColumnas[5].formatter = formatSelect;
                            return data.SModeloDeColumnas;
                        }
                        GeneraGrid(data.STituloTabla, data.SRenglones.rows, data.SNombresDeColumnas, Modelo, MetodoSelectRow, 'id');
                    } else {
                        alertify.error(data.SMsj);
                    }
                } else {
                    alertify.error("Hubo un problema al recibir información del servidor, favor de volver a intentar más tarde.");
                }
                $('#Loading').hide();
            },
            error: function(data) {
                $('#Loading').hide();
            },
        });
    } else {
        alertify.error("Debe ingresar un texto a buscar");
        $('#textoabuscar').focus();
    }
}

function GetAllToIndex(NModelo,val) {
    var url = '/' + NModelo + '/CargaIndex/'+val
    $('#Loading').show();
    $.ajax({
        url: url,
        type: 'POST',
        dataType: 'json',
        data: {},
        success: function(data) {

            if (data != null) {
                if (data.SEstado) {
                    function Modelo() {
                        data.SModeloDeColumnas[3].searchoptions["FECHA"] = datePick;
                        data.SModeloDeColumnas[13].searchoptions["HORA"] = datePick;
                        //data.SModeloDeColumnas[5].formatter = formatSelect;
                        return data.SModeloDeColumnas;
                    }
                    GeneraGrid(data.STituloTabla, data.SRenglones.rows, data.SNombresDeColumnas, Modelo, MetodoSelectRow, 'id');
                } else {
                    alertify.error(data.SMsj);
                }
            } else {
                alertify.error("Hubo un problema al recibir información del servidor, favor de volver a intentar más tarde.");
            }
            $('#Loading').hide();
        },
        error: function(data) {
            $('#Loading').hide();
        },
    });
}



var FiltrosDeGrid;

function GeneraGrid(EncabezadoDeGrid, CadenaURLDeMetodoParaObtenerDatos, ListadoDeEncabezadosDeGrid, ListadoDeColumnasDeModeloDeGrid, MetodoSelectRow, Surname, IdList, IdPager, MetodoAlFinalizarCarga, loadOnce, MetodoAfterInsertRow) {
    var FiltrosComplejos = "";
    var FiltrosToolbar = "";
    var Filtros = "";
    if (!IdList) {
        IdList = '#list'
    }

    if (!IdPager) {
        IdPager = '#pager'
    }
    if (!loadOnce) {
        loadOnce = false;
    }

    function resize_the_grid() {
        $(IdList).setGridWidth($("#SeccionCuerpo").width() * .95);
    }

    datePick = function(elem) {
        jQuery(elem).datepicker({
            dateFormat: 'yy/m/d',
            maxDate: new Date(2020, 0, 1),
            showOn: 'focus',
            onClose: function(dateText, inst) {
                $(elem).focus().trigger({ type: 'keypress', charCode: 13 });
                $(elem).focus();
            }
        });
    }
    $.jgrid.gridUnload(IdList)

    resize_the_grid();

    $(document).ready(function() {
        $(IdList).jqGrid({
            datatype: "local",
            data: CadenaURLDeMetodoParaObtenerDatos,

            // mtype: "GET",
            // datatype: "json",
            // url: CadenaURLDeMetodoParaObtenerDatos,

            styleUI: 'Bootstrap',
            colNames: ListadoDeEncabezadosDeGrid,
            colModel: ListadoDeColumnasDeModeloDeGrid(),
            pager: IdPager,
            rowNum: 10,
            rowList: [5, 10, 20, 50, 100],
            sortname: Surname,
            sortorder: 'desc',
            hidegrid: false,
            viewrecords: true,
            caption: EncabezadoDeGrid,
            scroll: false,
            loadonce: loadOnce, //dependiendo de este parámetro se harán peticiones a la base por cada acción de ordenamiento, busqueda, etc. o no.
            multiselect: false, //estos dos permiten seleccionar variso registros a la vez, en automático añade a la izquierda un control de tipo check box (los dos deben estar en true)
            multiboxonly: false,
            toolbar: [true, "top"], //establece una barra de búsqueda, este parámetro va ligado con las opciones de filter Toolbar mas abajo
            treegrid: true,
            colMenu: false,
            coloptions: { sorting: true, columns: true, filtering: false, seraching: false, grouping: true, freeze: false },
            shrinkToFit: true,
            altRows: true,
            onSelectRow: MetodoSelectRow,
            loadComplete: MetodoAlFinalizarCarga,
            afterInsertRow: MetodoAfterInsertRow

        });

        $(IdList).jqGrid('filterToolbar', {
            stringResult: true,
            searchOperators: true,
            searchOnEnter: true,
            enableClear: true
        });

        $(IdList).jqGrid('navGrid', IdPager, {
            search: true,
            add: false,
            edit: false,
            del: false,
            refresh: true,
            view: false,
            position: "left",
            cloneToTop: true
        }, {}, {}, {}, {
            multipleSearch: true,
            afterRedraw: function($p) {
                $("select.opsel").remove();
            },
            closeOnEscape: true,
            closeAfterSearch: true
        });


        $(IdList).navButtonAdd(IdPager, {
            buttonicon: "ui-icon-calculator",
            title: "Seleccionar Columnas",
            caption: "Ordenar y Seleccionar Columnas",
            position: "last",
            onClickButton: function() {
                $(IdList).jqGrid('columnChooser', {
                    done: function(perm) {
                        if (perm) {
                            this.jqGrid("remapColumns", perm, true);
                            resize_the_grid();
                        }
                    }
                });
            }
        });
        resize_the_grid();
    });
 
}


function MetodoSelectRow(id) {
    itemSeleccionado = $("#list").jqGrid('getGridParam', 'selrow');
    var query = ""

    query += $("#list").jqGrid('getCell', itemSeleccionado, 'Folio');
    window.location = '/DetalleCuentasPorCobrarVisorus/alta'; //'/CuentasPorCobrar/edita' + "/" + query;  
}

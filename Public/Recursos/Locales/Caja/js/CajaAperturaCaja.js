//##############################< SCRIPTS JS >##########################################
//################################< CajaApertura.js >#####################################
//#########################< VALIDACIONES DE JEQUERY >##################################

$(document).ready(function () {
    var validator = valida();
});

function valida() {
    var validator = $("#Form_Apertura_Caja").validate({
        rules: {
            Usuario: {
                required: true
            },
            Cargo: {
                required: true
            }

        },
        messages: {

            Usuario: {
                required: "El campo Usuario es requerido."
            },
            Cargo: {
                required: "El campo Cargo es requerido."
            },

        },
        errorElement: "em",
        errorPlacement: function (error, element) {
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
        success: function (label, element) {
            if (!$(element).next("span")[0]) {
                $("<span class='glyphicon glyphicon-ok form-control-feedback'></span>").insertAfter($(element));
            }
        },
        highlight: function (element, errorClass, validClass) {
            $(element).parents(".col-sm-5").addClass("has-error").removeClass("has-success");
            $(element).next("span").addClass("glyphicon-remove").removeClass("glyphicon-ok");
        },
        unhighlight: function (element, errorClass, validClass) {
            $(element).parents(".col-sm-5").addClass("has-success").removeClass("has-error");
            $(element).next("span").addClass("glyphicon-ok").removeClass("glyphicon-remove");
        }
    });
    return validator;
}
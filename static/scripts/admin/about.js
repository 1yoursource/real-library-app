
$(function() {
    $('#admCodeGenerate').on('click tap',function(){
        $.ajax({
            data: {
                adminId : window.userId,
            },
            url: '/ajax2/adm/auth/generate',
            type: 'POST',
            timeout: 15000,
            success: function(result) {
                if (result.error === null) {
                    alert("Код для реєстрації адміністратора");
                    alert(result.result);
                } else {
                    alert(result.error);
                }
            },
        });
    });

});

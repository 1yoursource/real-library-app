function getById(id) {
    return document.getElementById(id)
}

$(function() {
    $('#regSubmitlBtn').on('click tap',function(){
        $.ajax({
            data: {
                email : getById("reg_email").value,
                password : getById("reg_password").value,
                passwordSubmit : getById("reg_password-confirmation").value,
                firstName : getById("reg_first_name").value,
                lastName : getById("reg_last_name").value,
                surName : getById("reg_surname").value,
                faculty : getById("reg_faculty").value,
            },
            url: '/ajax/auth/registration',
            type: 'POST',
            timeout: 15000,
            success: function(result) {
                if (result.error === null) {
                    window.location = result.url;
                } else {
                       alert(result.error);
                    window.href="/";
                }
            },
        });
    });

    $('#authSubmitBtn').on('click tap',function(){
        $.ajax({
            data: {
                login : getById("auth_email").value,
                password : getById("auth_password").value,
            },
            url: '/ajax/auth/authorization',
            type: 'POST',
            timeout: 15000,
            success: function(result) {
            console.log("r: ",result)
                if (result.error === null) {
                    window.location = result.url;
                } else {
                       alert(result.error);
                    window.href="/"
                }
            },
        });
    });
});
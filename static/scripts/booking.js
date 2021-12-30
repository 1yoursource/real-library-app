function getById(id) {
    return document.getElementById(id)
}
$(function() {
    $('#getBook').on('click tap',function(){
        console.log("here 1 booking")
        $.ajax({
            url: '/ajax/booking/getBook',
            type: 'POST',
            timeout: 15000,
            error: function(result) {
                console.log("ftyguh")
            },
            success: function(result) {
                console.log("ftyguh success")
            },
        });
    });
});
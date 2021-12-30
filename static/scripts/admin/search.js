function getById(id) {
    return document.getElementById(id)
}

$(function() {

    $('#searchBtn').on('click tap',function(){
                console.log("here 1")
            $.ajax({
                url: '/ajax2/adm/book/search',
                type: 'POST',
                timeout: 15000,
                error: function(result) {
                    console.log(result)
                },
                success: function(result) {
                    if (result.error === null) {
                    window.res = result.result
                    for (var i = 0; i<result.result.length; i++) {
                        console.log(result.result[i]);
                        $("#adm_search_table").append('<tr>'+
                            '<td>'+result.result[i].Id+'</td>'+
                            '<td>'+result.result[i].Name+'</td>'+
                            '<td>'+result.result[i].Author+'</td>'+
                            '<td><a id="adm_book_delete"></a></td>'+
                        '</tr>');

                    }
                    } else {
                        console.log("eeeeeeeerr")
                    }
                },
            });
        });

});

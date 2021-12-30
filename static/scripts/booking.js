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

    $('#searchBookByName').on('click tap',function(){
        console.log("here 1 searchBookByName")
        $.ajax({
            url: '/ajax/booking/searchBookByName',
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

    $('#searchBookByAuthor').on('click tap',function(){
        console.log("here 1 searchBookByAuthor")
        $.ajax({
            url: '/ajax/booking/searchBookByAuthor',
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

    $('#getAllUsersBooks').on('click tap',function(){
        var userId = getCookie("lib-id");
        console.log("here 1 getAllUsersBooks window.userId", userId)

        $.ajax({
            url: '/ajax/booking/getAllTakenBooks',
            data: {userId: userId},
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
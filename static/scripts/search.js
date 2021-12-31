function getById(id) {
    return document.getElementById(id)
}

$(function() {

    $('#searchBtn').on('click tap',function(){
        console.log("f-",$('#search_filter').val())
        var filterId = $('#search_filter').val();
        var filterVal = $('#input_search').val();
        search(filterId, filterVal)
    });

});

   function book_take(id){
           $.ajax({
               data: {
                    bookId: id,
               },
               url: '/ajax2/usr/book/take',
               type: 'POST',
               timeout: 15000,
               error: function(result) {
                   console.log("err:",result)
               },
               success: function(result) {
                   if (result.error === null) {
                        $("#adm_search_table tr").remove('.table-row-'+id);
                   } else {
                       console.log("ERR")
                   }
               },
           });
       }

   function search(filterId, filterVal){
           $.ajax({
               data: {
                   filter: filterId,
                   value: filterVal,
               },
               url: '/ajax2/usr/book/search',
               type: 'POST',
               timeout: 15000,
               error: function(result) {
                   console.log(result)
               },
               success: function(result) {
                   if (result.error === null) {
                       if (result.result == null || result.result.length == 0) {
                           return
                       }
                   $("#user_search_table td").parent().remove();
                       switch (filterId) {
                           case "1":
                               makeTableResult(result.result);
                               break;
                           case "2": // by name
                               makeTableResult(result.result);
                               break;
                           case "3": // by author
                               makeTableResult(result.result);
                               break;
                       }
                   } else {
                       console.log("eeeeeeeerr")
                   }
               },
           });
       }

    function makeTableResult(result) {
        if (result == null || result.length == 0) {
            return
        }
        for (var i = 0; i<result.length; i++) {
            $("#user_search_table").append('<tr class="table-row-'+result[i].Id+'">'+
            '<td>'+result[i].Id+'</td>'+
            '<td>'+result[i].Name+'</td>'+
            '<td>'+result[i].Author+'</td>'+
            '<td><button onclick="book_take(this.id)" id="'+result[i].Id+'" class="book_take">Взяти</button></td>'+
            '</tr>');
        }
    }














//
//function getById(id) {
//    return document.getElementById(id)
//}
//
//$(function() {
//    $('#getBook').on('click tap',function(){
//        $.ajax({
//            url: '/ajax/booking/getBook',
//            type: 'POST',
//            timeout: 15000,
//            error: function(result) {
//                console.log("ftyguh")
//            },
//            success: function(result) {
//                console.log("ftyguh success")
//            },
//        });
//    });
//
//    $('#searchBookByName').on('click tap',function(){
//        console.log("here 1 searchBookByName")
//        var bookName = $('#searchBookByNameInput').val();
//
//        console.log("here 1 searchBookByName", bookName)
//        $.ajax({
//            url: '/ajax/booking/searchBookByName',
//            data: {bookName: bookName},
//            type: 'POST',
//            timeout: 15000,
//            error: function(result) {
//                console.log(result)
//            },
//            success: function(result) {
//                if (result.error === null) {
//                    window.res = result.result
//                    for (var i = 0; i<result.result.length; i++) {
//
//                        if (result.result[i].takenBy === getCookie("lib-id")) {
//                            var bookAction = '<a id="user_book_return" >Повернути книгу</a>'
//                        } else if (result.result[i].takenBy === ""){
//                            var bookAction =  '<input type="submit" title="Взяти книгу" id="searchBookByAuthorInput" value="result.result[i].id" class="form-control required" data-valid="email">'
//
//                              } else {
//                            var bookAction = '<a id="user_book_absent">Книги немає у наявності</a>'
//                        }
//                        console.log(result.result[i]);
//                        $("#user_search_table").append('<tr>'+
//                            '<td>'+result.result[i].id+'</td>'+
//                            '<td>'+result.result[i].name+'</td>'+
//                            '<td>'+result.result[i].author+'</td>'+
//                            '<td>'+bookAction+'</td>'+
//                            '</tr>');
//
//                    }
//                } else {
//                    console.log("sfesffsfse result.error", result.error)
//                    console.log("eeeeeeeerr")
//                }
//            },
//        });
//    });
//
//    $('#searchBookByAuthor').on('click tap',function(){
//        console.log("here 1 searchBookByAuthor")
//        var author = $('#searchBookByAuthorInput').val();
//        $.ajax({
//            url: '/ajax/booking/searchBookByAuthor',
//            data: {author: author},
//            type: 'POST',
//            timeout: 15000,
//            error: function(result) {
//                console.log(result)
//            },
//            success: function(result) {
//                console.log("sfesffsfse result.error")
//                if (result.error === null) {
//                    window.res = result.result
//                    for (var i = 0; i<result.result.length; i++) {
//                        console.log(result.result[i]);
//                        $("#user_search_table").append('<tr>'+
//                            '<td>'+result.result[i].id+'</td>'+
//                            '<td>'+result.result[i].name+'</td>'+
//                            '<td>'+result.result[i].author+'</td>'+
//                            '<td><a id="adm_book_delete"></a></td>'+
//                            '</tr>');
//
//                    }
//                } else {
//                    console.log("eeeeeeeerr")
//                    console.log("sfesffsfse result.error", result.error)
//                }
//            },
//        });
//    });
//
//    $('#getAllUsersBooks').on('click tap',function(){
//        var userId = getCookie("lib-id");
//        console.log("here 1 getAllUsersBooks window.userId", userId)
//
//        $.ajax({
//            url: '/ajax/booking/getAllTakenBooks',
//            data: {userId: userId},
//            type: 'POST',
//            timeout: 15000,
//            error: function(result) {
//                console.log(result)
//            },
//            success: function(result) {
//                if (result.error === null) {
//                    window.res = result.result
//                    for (var i = 0; i<result.result.length; i++) {
//                        console.log(result.result[i]);
//                        $("#adm_search_table").append('<tr>'+
//                            '<td>'+result.result[i].id+'</td>'+
//                            '<td>'+result.result[i].name+'</td>'+
//                            '<td>'+result.result[i].author+'</td>'+
//                            '<td><a id="adm_book_delete">Повернути книгу</a></td>'+
//                            '</tr>');
//
//                    }
//                } else {
//                    console.log("eeeeeeeerr")
//                }
//            },
//        });
//    });
//    $('#user_book_get').on('click tap',function(){
//        console.log("here 1 user_book_get")
//        var bookId = $('#user_book_get').val();
//        console.log("here 1 searchBookByName", bookId)
//        $.ajax({
//            url: '/ajax/booking/getBook',
//            data: {bookId: bookId, userId: getCookie("lib-id")},
//            type: 'POST',
//            timeout: 15000,
//            error: function(result) {
//                console.log(result)
//            },
//            success: function(result) {
//                console.log("success",result)
//            },
//        });
//    });
//});
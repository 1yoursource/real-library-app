function getById(id) {
    return document.getElementById(id)
}

$(function() {

    $('#searchBtn').on('click tap',function(){
        var filterId = $('#search_filter').val();
        var filterVal = window.userId;
        console.log("filterId: ",filterId)
        console.log("filterVal: ",filterVal)
        searchShell(filterId, filterVal)
    });

});

   function book_return(id){
           $.ajax({
               data: {
                    bookId: id,
                    userId: window.userId,
               },
               url: '/ajax2/usr/book/return',
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

   function searchShell(filterId, filterVal){
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
                           case "4":
                               makeTableResultShell(result.result);
                               break;
                           case "5":
                               makeTableResultShell(result.result);
                               break;
                       }
                   } else {
                       console.log("eeeeeeeerr")
                   }
               },
           });
       }

    function makeTableResultShell(result) {
        if (result == null || result.length == 0) {
            return
        }
        for (var i = 0; i<result.length; i++) {
            $("#user_search_table").append('<tr class="table-row-'+result[i].Id+'">'+
            '<td>'+result[i].Id+'</td>'+
            '<td>'+result[i].Name+'</td>'+
            '<td>'+result[i].Author+'</td>'+
            '<td><button onclick="book_return(this.id)" id="'+result[i].Id+'" class="book_take">Повернути</button></td>'+
            '</tr>');
        }
    }

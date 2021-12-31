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

   function book_del(id){
           $.ajax({
               data: {
                    bookId: id,
               },
               url: '/ajax2/adm/book/delete',
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
               url: '/ajax2/adm/book/search',
               type: 'POST',
               timeout: 15000,
               error: function(result) {
                   console.log(result)
               },
               success: function(result) {
                   if (result.error === null) {
                       if (result.error !== null && result.result.length == 0) {
                           return
                       }
                   $("#adm_search_table td").parent().remove();
                       switch (filterId) {
                           case "1":
                               makeTableResult(result.result);
                               break;
                           case "2":
                               makeTableResult(result.result);
                               break;
                           case "3":
                               makeTableResult(result.result);
                               break;
                           case "4": // должники
                               break;
                       }
                   } else {
                       console.log("eeeeeeeerr")
                   }
               },
           });
       }

    function makeTableResult(result) {
        if (result.length == 0) {
            return
        }
        for (var i = 0; i<result.length; i++) {
            $("#adm_search_table").append('<tr class="table-row-'+result[i].Id+'">'+
            '<td>'+result[i].Id+'</td>'+
            '<td>'+result[i].Name+'</td>'+
            '<td>'+result[i].Author+'</td>'+
            //'<td><button onclick="book_edit(this.id)" id="'+result.result[i].Id+'" class="adm_book_edit">Edit</button>'+
            '<td><button onclick="book_del(this.id)" id="'+result[i].Id+'" class="adm_book_edit">Delete</button></td>'+
            '</tr>');
        }
    }

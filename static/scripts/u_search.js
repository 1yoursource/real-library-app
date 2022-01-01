function getById(id) {
    return document.getElementById(id)
}

$(function() {

    $('#usrSearchBtn').on('click tap',function(){
        console.log("f-",$('#search_filter').val())
        var filterId = $('#search_filter').val();
        var filterVal = $('#input_search').val();
        searchSearch(filterId, filterVal)
    });

});

   function book_take(id){
           $.ajax({
               data: {
                    bookId: id,
                    userId: window.userId,
               },
               url: '/ajax2/usr/book/take',
               type: 'POST',
               timeout: 15000,
               error: function(result) {
                   console.log("err:",result)
               },
               success: function(result) {
                   if (result.error === null) {
                        $("#user_search_table tr").remove('.table-row-'+id);
                   } else {
                       alert(result.error);
                   }
               },
           });
       }

   function searchSearch(filterId, filterVal){
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
                   console.log("rs",result.result)
                       switch (filterId) {
                           case "1":
                               makeTableResultSearch(result.result);
                               break;
                           case "2": // by name
                               makeTableResultSearch(result.result);
                               break;
                           case "3": // by author
                               makeTableResultSearch(result.result);
                               break;
                       }
                   } else {
                       alert(result.error);
                   }
               },
           });
       }

    function makeTableResultSearch(result) {
        if (result == null || result.length == 0) {
            return
        }
        for (var i = 0; i<result.length; i++) {
        console.log("I: ",result[i])
            $("#user_search_table").append('<tr class="table-row-'+result[i].Id+'">'+
            '<td>'+result[i].Id+'</td>'+
            '<td>'+result[i].Name+'</td>'+
            '<td>'+result[i].Author+'</td>'+
            '<td><button onclick="book_take(this.id)" id="'+result[i].Id+'" class="book_take">Взяти</button></td>'+
            '</tr>');
        }
    }

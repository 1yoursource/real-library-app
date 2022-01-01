function getById(id) {
    return document.getElementById(id)
}

$(function() {

    $('#adminAddBookBtn').on('click tap',function(){
            $.ajax({
                data: {
                    name : getById("book_data_name").value,
                    author : getById("book_data_author").value,
                    publishYear : getById("book_data_year").value,
                    publisher : getById("book_data_publisher").value,
                    pagesCount : getById("book_data_pcount").value,
                },
                url: '/ajax2/adm/book/add',
                type: 'POST',
                timeout: 15000,
                success: function(result) {
                console.log("r: ",result)
                    if (result.error === null) {
                        console.log("res:", result)
                        window.location.reload();
                    } else {
                       alert(result.error);
                    }
                },
            });
        });

    $('#admShellSearchBtn').on('click tap',function(){
        var filterId = $('#search_filter').val();
        searchShell(filterId)
    });

});

   function searchShell(filterId){
       $.ajax({
           data: {
               filter: filterId,
           },
           url: '/ajax2/adm/user/dept',
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
               $("#adm_dept_table td").parent().remove();
                   switch (filterId) {
                       case "1":
                           makeTableResultShell(result.result);
                           break;
                       case "2":
                           makeTableResultShell(result.result);
                           break;
                   }
               } else {
                       alert(result.error);
               }
           },
       });
   }

    function makeTableResultShell(result) {
        if (result.length == 0) {
            return
        }
        for (var i = 0; i<result.length; i++) {
            $("#adm_dept_table").append('<tr class="table-row-'+result[i].Id+'">'+
            '<td>'+result[i].TicketNumber+'</td>'+
            '<td>'+result[i].Book+'</td>'+
            '<td>'+result[i].ReturnDate+'</td>'+
            '</tr>');
        }
    }

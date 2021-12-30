function getById(id) {
    return document.getElementById(id)
}

   function book_del(id){
   console.log("del: ", id);
   //        $.ajax({
   //            url: '/ajax2/adm/book/edit',
   //            type: 'POST',
   //            timeout: 15000,
   //            error: function(result) {
   //                console.log(result)
   //            },
   //            success: function(result) {
   //                if (result.error === null) {
   //
   //                } else {
   //                    console.log("ERR")
   //                }
   //            },
   //        });
       }

$(function() {



    $('#searchBtn').on('click tap',function(){
    console.log("f-",$('#search_filter').val())
    var filterValue = $('#search_filter').val();
            $.ajax({
                data: {
                    filter: filterValue,
                },
                url: '/ajax2/adm/book/search',
                type: 'POST',
                timeout: 15000,
                error: function(result) {
                    console.log(result)
                },
                success: function(result) {
                    if (result.error === null) {
                    $("#adm_search_table td").parent().remove();
                        switch (filterValue) {
                            case "1":
                                for (var i = 0; i<result.result.length; i++) {
                                    console.log(result.result[i]);
                                    $("#adm_search_table").append('<tr data-value="'+result.result[i].Id+'">'+
                                    '<td>'+result.result[i].Id+'</td>'+
                                    '<td>'+result.result[i].Name+'</td>'+
                                    '<td>'+result.result[i].Author+'</td>'+
//                                    '<td><button onclick="book_edit(this.id)" id="'+result.result[i].Id+'" class="adm_book_edit">Edit</button>'+
                                    '<td><button onclick="book_del(this.id)" id="'+result.result[i].Id+'" class="adm_book_edit">Delete</button></td>'+
                                    '</tr>');
                                }
                                break;
                            case "2":
                                break;
                            case "3":
                                break;
                        }
                    } else {
                        console.log("eeeeeeeerr")
                    }
                },
            });
        });







});


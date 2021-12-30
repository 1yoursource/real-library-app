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
                    } else {
                        console.log("eeeeeeeerr")
                    }
                },
            });
        });

});

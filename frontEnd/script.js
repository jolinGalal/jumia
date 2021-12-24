// api url
const api_url = 
      "http://localhost:8090/customers/v1"  
  
      $(document).ready(function() {
        $.ajax({
          url: 'http://localhost:8090/customers/v1',
          dataType: 'json',
          success: function(data) {
              console.log(data)
              for (var i=0; i<data.customers.length; i++) {
                  var row = $('<tr><td>' + data.customers[i].ID+ '</td><td>' + data.customers[i].Name +
                   '</td><td>' + data.customers[i].Phone + '</td><td>' + data.customers[i].Country +
                   '</td><td>' + data.customers[i].State +'</td><td>' + data.customers[i].CountryCode+
                    '</td></tr>');
                  $('#myTable').append(row);
              }
              $("#total").html("Total records:"+data.pagination.total_count);
              $("#pages").html("/"+data.pagination.total_pages);

              
           },
          error: function(jqXHR, textStatus, errorThrown){
              alert('Error: ' + textStatus + ' - ' + errorThrown);
          }
        });
      });

      $( "#refresh" ).click(function() {
        var pageNumber = $('#pageNumber').val();
        $('#pageNumber').val(pageNumber);
        $.ajax({ url: 'http://localhost:8090/customers/v1?page_number='+pageNumber,
        dataType: 'json',
        success: function(data) {
            console.log(data)
            $("#myTable tr").remove();
            for (var i=0; i<data.customers.length; i++) {
                var row = $('<tr><td>' + data.customers[i].ID+ '</td><td>' + data.customers[i].Name +
                 '</td><td>' + data.customers[i].Phone + '</td><td>' + data.customers[i].Country +
                 '</td><td>' + data.customers[i].State +'</td><td>' + data.customers[i].CountryCode+
                  '</td></tr>');
                $('#myTable').append(row);
            }
            $("#total").html("Total records:"+data.pagination.total_count);
            $("#pages").html("/"+data.pagination.total_pages+"    ");

            
         },
        error: function(jqXHR, textStatus, errorThrown){
            alert('Error: ' + textStatus + ' - ' + errorThrown);
        }
      });
      });
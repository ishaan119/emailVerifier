/*
(function($) {
    "use strict";

    jQuery(document).ready(function(){
        $('#vform').submit(function(){
            $(this).preventDefault();
            var action = $(this).attr('action');

            $("#message").slideUp(750,function() {
                $('#message').hide();

                $('#submit')
                    .before('<img src="images/ajax-loader.gif" class="contact-loader" />')
                    .attr('disabled','disabled');

                $.ajax(action, {
                        email: $('#email').val()
                    },
                    function(data){
                        document.getElementById('message').innerHTML = data;
                        $('#message').slideDown('slow');
                        $('#cform img.contact-loader').fadeOut('slow',function(){$(this).remove()});
                        $('#vsubmit').removeAttr('disabled');
                        if(data.match('success') != null) $('#vform').slideUp('slow');
                    }
                );

            });

            return false;

        });

    });

}(jQuery));
*/

(function($) {
    "use strict";

    jQuery(document).ready(function(){
        $('#vform').submit(function(){
            console.log("sasdasd")
            $(this).preventDefault();
            var email = $('#email').val(),
            formData = 'email=' + email
            console.log(formData)
            $.ajax({
                type: 'post',
                url: '/email-verifier/api/1.0/emails/verify' + $.param({ email:email}),
                data: formData,
                success: function(results) {
                    $('ul#response').html(results);
                }
            });


        });

    });

}(jQuery));
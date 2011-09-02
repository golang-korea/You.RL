
$(document).ready(function() {
    $("#btn_shorten").click(function() {
        $.ajax({
            url: "/shorten",
            type: "POST",
            data: {url: $("#url").val()},
            dataType: "json",
            success: function(ret) {
                $("#shorten_url").val(ret.url);
                $("#shorten_url").select();
            },
            error: function(xhr, status, err) {
                alert('error: ' + status);
            }
        });
    });
});

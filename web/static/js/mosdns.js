$(document).ready(function() {
    // Get initial MosDNS configuration status
    $.getJSON("/api/config", function(data) {
        $("#config-status").text("MosDNS enabled: " + data.enabled);
    });

    // Update Geo data on button click
    $("#update-geodata").click(function() {
        $("#geodata-update-status").text("Updating...");
        $.post("/api/geodata/update", function(data) {
            $("#geodata-update-status").text(data);
        });
    });

    // Get MosDNS status
    $.getJSON("/api/status", function(data) {
        $("#mosdns-status").text(data);
    });

    // Get MosDNS logs
    $.get("/api/logs", function(data) {
        $("#mosdns-logs").text(data);
    });


    // ... add other JavaScript functions to interact with the API
});

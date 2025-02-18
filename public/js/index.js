function updateArtists(value) {

    data = {
        "userSearch" : document.getElementById("search").value,
        "filters" : {
            "memberCount" : +document.getElementById("memberCount").value,
            "creationDateMin" : +document.getElementById("creationDateMin").value,
            "creationDateMax" : +document.getElementById("creationDateMax").value,
            "firstAlbumMin" : +document.getElementById("firstAlbumMin").value,
            "firstAlbumMax" : +document.getElementById("firstAlbumMax").value,
            "location" : +document.getElementById("location").value,
        }
    }

    $.ajax({
        type: "POST",
        url: "/",
        data: JSON.stringify(data),
        dataType: "text",
        contentType : "application/json;charset=utf-8",
        success: function(data)
        {
            document.getElementById("artists").innerHTML = data
        },
        error: function(xhr, status, error) {
            document.getElementById("artists").innerHTML = "Could not update artists."
        }
    });
}

function updateMemberCountUI(value) {
    if (value == "0") {
        document.getElementById("memberCountLabel").innerHTML = "Member count : " + "No filter"
    } else {
        document.getElementById("memberCountLabel").innerHTML = "Member count : " + value
    }
}
$(document).ready(() => {

    api.getRooms().then(resRooms => {
        if (resRooms.success)
            resRooms.rooms.forEach(resRoom => {
                $(`<a>${resRoom.name}</a>`)
                    .addClass("roomLink")
                    .attr("href", `/room/${resRoom.name}`)
                    .on({
                        click: ev => {
                            $(this).toggleClass("active")
                        }
                    }).appendTo("#rooms");
            });
        else
            $(`<p>${resRooms.answer}</p>`)
                .addClass("error")
                .appendTo("#rooms");
    });

});

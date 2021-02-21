$(document).ready(() => {

    (async () => {
        let rooms = await api.getRooms();
        switch (typeof rooms) {
            case "string":
                $( `<p>${rooms.reason}</p>` )
                    .addClass('error')
                    .appendTo('#rooms');
            case "object":
                if (rooms == null)
                    return $( `<p>Żaden pokój nie został utworzony</p>` )
                        .addClass('error')
                        .appendTo('#rooms');
                rooms.forEach(r =>
                    $( `<a>${r.name}</a>` )
                        .addClass('roomLink')
                        .attr('href', `/room/${r.name}`)
                        .on({
                            click: ev =>
                                $(this).toggleClass('active')})
                        .appendTo('#rooms')
                );
        }
    })();

});

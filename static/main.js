$(document).ready(() => {

    (async () => {
        let rooms = await api.getRooms();
        switch (typeof rooms) {
            case "string":
                $( `<p>${rooms.reason}</p>` )
                    .addClass('error')
                    .appendTo('#rooms');
                break
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
        $('#buildingLayout div[class^="r"]').each((i, r) => (r.id === '') ? $(`<label>${rooms[i].name}</label>`).addClass('roomLabel').appendTo(r) : '');

    })();

    (async () =>
        (await api.getStands()).forEach(stand =>
            $( `<div>${stand.name}</div>` )
                .addClass('stand')
                .on({click: () => window.open(stand.description, '_blank')})
                .appendTo($('#stands'))))();

    let prevScale = 0.9;
    let scaleMap = () => {
            while (buildingLayout.getBoundingClientRect().height + 120 >= $(window).height()) {
                if (prevScale < 0.5) return;
                prevScale -= .05
                $('#buildingLayout').css('transform', `translate(-50%, -50%) scale(${prevScale.toString()})`);
            }
            while (buildingLayout.getBoundingClientRect().height + 200 <= $(window).height()) {
                if (prevScale > 1.5) return;
                prevScale += .05
                $('#buildingLayout').css('transform', `translate(-50%, -50%) scale(${prevScale.toString()})`);
            }
    }
    scaleMap(); window.onresize = scaleMap;


    $('#toilet').on({click: ev => {
        ev.preventDefault();
        (new Audio('/static/easter.mp3')).play();
    }});

    $( '#navAccountRegBtn' ).on({click: ev => {
                ev.preventDefault();
                $( '#accountRegister' ).show();
                // jQuery animate breaks Firefox?
                accountRegister.animate([
                        {transform: 'translateY(-10px)'},
                        {transform: 'translateY(0px)'}],
                    {duration: 200, iterations: 1, fill: 'forwards', easing: 'ease-out'});
    }});

    $( '#accountRegisterSubmit' ).on({click: async ev => {
        ev.preventDefault();
        $('#accountRegisterSubmit')
            .empty()
            .prop('disabled', 'true');
        $( '<div></div>' )
            .addClass('loading')
            .css('width', '16px')
            .css('height', '16px')
            .appendTo($('#accountRegisterSubmit'));
        await api.newUser($('#accountRegister form input:nth-child(1)').val(), $('#accountRegister form input:nth-child(3)').val(),
            $('#accountRegister form input:nth-child(2)').val()).then(res => {
           if (res.success) {
               $('#accountRegisterSubmit')
                   .empty()
                   .prop('disabled', 'false')
                   .text('Utwórz konto');
               $('#accountRegister').hide();
               $('#accountDiscordConfirm')
                   .attr('userID', res.uuid)
                   .show();
               accountDiscordConfirm.animate([
                       {transform: 'translateY(-10px)'},
                       {transform: 'translateY(0px)'}],
                   {duration: 200, iterations: 1, fill: 'forwards', easing: 'ease-out'});
           }
        });
    }});

    $( '#accountDiscordConfirmSubmit' ).on({click: async ev => {
        ev.preventDefault();
            $('#accountDiscordConfirmSubmit')
                .empty()
                .prop('disabled', 'true');
            $( '<div></div>' )
                .addClass('loading')
                .css('width', '16px')
                .css('height', '16px')
                .appendTo($('#accountDiscordConfirmSubmit'));
            await api.confirmDiscord($('#accountDiscordConfirm').attr('userID'),
                $('#accountDiscordConfirm form input:nth-child(1)').val()).then(res => {
                    $('#accountDiscordConfirm').hide();
                    $(`<div>${res.reason}</div>`)
                        .addClass('error')
                        .appendTo(document.body);
                });
    }});

});

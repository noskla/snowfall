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
                rooms.forEach((r, i) =>
                    $( `<a>${r.name.replace('l¢', 'łó')}</a>` )
                        .addClass('roomLink')
                        .attr('href', `/room/${r.name}`)
                        .on({
                            click: ev => {
                                ev.preventDefault();
                                $(`<div></div>`)
                                    .css('background', 'rgba(0,0,0,.75)')
                                    .css('position', 'fixed')
                                    .css('width', '100%')
                                    .css('height', '100%')
                                    .css('transition', 'all ease .2s')
                                    .css('opacity', 0)
                                    .attr('id', 'darken')
                                    .on({click: () => {
                                            darken.animate([{opacity: 1}, {opacity: 0}], {duration: 150, fill: 'forwards'})
                                            setTimeout(() => $('#darken').remove(), 150);
                                            $(`#player`).remove();
                                        }})
                                    .appendTo(document.body);
                                darken.animate([{opacity: 0}, {opacity: 1}], {duration: 200, fill: 'forwards'});
                                $(`<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/${rooms[i].streamURL.split('watch?v=')[1]}" frameborder="0" allow="autoplay; clipboard-write; encrypted-media;" allowfullscreen></iframe>`)
                                    .attr('id', 'player')
                                    .css('position', 'absolute')
                                    .css('top', '50%')
                                    .css('left', '50%')
                                    .css('transform', 'translate(-50%, -50%')
                                    .css('box-shadow', '0 0 15px #000000')
                                    .appendTo(document.body);
                            }
                        })
                        .appendTo('#rooms')
                );
        }
        let s = 0;
        $($('#buildingLayout div[class^="r"]').get().reverse()).each((i, r) => {
            if (r.id === '') {
                if (i === 5 || i === 7) {
                    $(i === 5 ? `<label>Org Room</label>` : `<label>Schowek na miotły</label>`).addClass('roomLabel').appendTo(r);
                    s--;
                } else {
                    $(`<label room="${i+s}">${rooms[i + s].name.replace('l¢', 'łó')}</label>`).addClass('roomLabel').css('cursor', 'pointer')
                        .css('text-decoration', rooms[i+s].streamURL === '' ? 'none' : 'underline')
                        .on({
                        click: ev => {
                            ev.preventDefault();
                            $(`<div></div>`)
                                .css('background', 'rgba(0,0,0,.75)')
                                .css('position', 'fixed')
                                .css('width', '100%')
                                .css('height', '100%')
                                .css('transition', 'all ease .2s')
                                .css('opacity', 0)
                                .attr('id', 'darken')
                                .on({click: () => {
                                        darken.animate([{opacity: 1}, {opacity: 0}], {duration: 150, fill: 'forwards'})
                                        setTimeout(() => $('#darken').remove(), 150);
                                        $(`#player`).remove();
                                    }})
                                .appendTo(document.body);
                            darken.animate([{opacity: 0}, {opacity: 1}], {duration: 200, fill: 'forwards'});
                            $(`<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/${rooms[ev.target.getAttribute('room')].streamURL.split('watch?v=')[1]}" frameborder="0" allow="autoplay; clipboard-write; encrypted-media;" allowfullscreen></iframe>`)
                                .attr('id', 'player')
                                .css('position', 'absolute')
                                .css('top', '50%')
                                .css('left', '50%')
                                .css('transform', 'translate(-50%, -50%')
                                .css('box-shadow', '0 0 15px #000000')
                                .appendTo(document.body);
                        }
                    }).appendTo(r);
                }
            } else {
                $(`<label>Toalety</label>`).addClass('roomLabel').appendTo(r);
                s--;
            }
        });

    })();

    (async () =>
        (await api.getStands()).forEach(stand =>
            $( `<div>${stand.name}</div>` )
                .addClass('stand')
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

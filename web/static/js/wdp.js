let lost = false;

function main() {
    const socket = new WebSocket(`ws://${location.host}/.wdp/ws`);

    socket.onopen = () => {
        console.log("[wdp] init");
        if (lost) {
            socket.close();
            location.reload();
        }
    }

    socket.onmessage = () => {
        console.log("[wdp] reload");
        location.reload();
    }

    socket.onclose = e => {
        if (e.wasClean)
            console.log("[wdp] close");
        else
            console.log("[wdp] died");

        lost = true;
        main();
    };

    socket.onerror = () => {
        console.log("[wdp] error");
    };
}

main();
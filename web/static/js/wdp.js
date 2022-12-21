console.log("[wdp] init");

const socket = new WebSocket(`ws://${location.host}/.wdp/ws`);

socket.onopen = () => {
    console.log("[wdp] init");
}

socket.onmessage = () => {
    console.log("[wdp] reload");
    location.reload();
}

socket.onclose = (e) => {
    if (e.wasClean) {
        console.log("[wdp] close");
    } else {
        console.log("[wdp] died");
    }
};

socket.onerror = e => {
    console.log("[wdp] error");
};
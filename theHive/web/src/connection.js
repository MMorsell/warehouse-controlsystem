import { advancePath } from "./index.js"

const socket = new WebSocket("ws://localhost:8000/websocketConnection");
const status = document.getElementById("status");

socket.onopen = function () {
    status.innerHTML += "Status: Connected\n";
};

socket.onmessage = function (e) {
    const newMessage = JSON.parse(e.data);
    console.log(newMessage)
    advancePath({id: newMessage.RobotId, x: XPosition, y: YPosition})
};
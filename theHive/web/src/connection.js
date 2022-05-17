import { advancePath, clearPath } from "./index.js"

const socket = new WebSocket("ws://localhost:8000/websocketConnection");
const status = document.getElementById("status");

socket.onopen = function () {
    status.innerHTML += "Status: Connected\n";
};

socket.onmessage = function (e) {
    const newMessage = JSON.parse(e.data);
    console.log(newMessage)
    
    if (newMessage.XPosition === -1 && newMessage.YPosition === -1) {
        clearPath(newMessage.RobotId);
        return
    }
    advancePath({id: newMessage.RobotId, x: newMessage.XPosition, y: newMessage.YPosition})
};
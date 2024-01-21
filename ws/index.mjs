import { WebSocketServer } from "ws";

const wss = new WebSocketServer({ port: 8080 });

wss.on("open", function connection(ws) {
  console.log("open");
  window.location.reload();
});

wss.on("close", function close() {
  console.log("disconnected");
});

wss.on("error", function error() {
  console.log("error");
});

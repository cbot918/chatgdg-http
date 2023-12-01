const express = require("express");
const app = express();
const http = require("http");
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server);

app.get("/", (req, res) => {
  res.sendFile(__dirname + "/ui/index.html");
});

io.on("connection", (socket) => {
  console.log("a user connected");

  socket.broadcast.emit("hi");

  socket.on("chat message", (msg) => {
    console.log(msg);
    io.emit("chat message", msg);
  });
});

const port = 3001;
server.listen(port, () => {
  console.log(port);
});

window.onload = function () {};

let url = "ws://localhost:8889/ws";
let ws = {};
let userCount = document.querySelector("#user_count");
function sendID(e) {
  if (e.key === "Enter") {
    s.setName(e.target.value);
    // console.log(s.getName());
    e.target.value = "";
    enableDisable("#user_name", true);
    enableDisable("#message", false);
    connect(s.getName());
  }
}

function sendMessage(event) {
  if (event.key === "Enter" && ws.readyState === WebSocket.OPEN) {
    let messageInput = document.getElementById("message");

    encodedMessage = encode("temp", messageInput.value);

    console.log(encodedMessage);
    ws.send(encodedMessage);
    messageInput.value = ""; // Clear the input field after sending
  }
}

function enableDisable(name, state) {
  document.querySelector(name).disabled = state;
}

function connect(name) {
  if (name) {
    url += "?name=" + name;
  }

  ws = new WebSocket(url);
  ws.onopen = (e) => {
    console.log("ws opened");
  };

  ws.onmessage = (e) => {
    console.log(s.getName() + ": " + e.data);
    let strArr = decode(e.data);
    if (IsUserCount(strArr[0])) {
      userCount.innerHTML = strArr[1];
    }
  };

  ws.onclose = (e) => {
    console.log("ws closeed");
    console.log(e);
  };

  ws.onerror = (err) => {
    console.log("ws error");
    console.log(err);
  };
}

function store() {
  let _name = "";

  function getName() {
    return _name;
  }
  function setName(name) {
    _name = name;
  }

  return {
    setName,
    getName,
  };
}
window.s = store();

function decode(m) {
  return m.split(";");
}

function encode(channel, data) {
  return channel + ";" + data;
}

function IsUserCount(type) {
  return type === "userCount";
}

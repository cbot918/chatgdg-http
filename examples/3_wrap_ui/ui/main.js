window.onload = function () {};

let url = "ws://localhost:8889/ws";
let ws = {};

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
    let message = messageInput.value;
    ws.send(message);
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
    // ws.send("hihi");
  };

  ws.onmessage = (e) => {
    console.log(s.getName() + ": " + e.data);
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
  let nameState = true;
  let messageState = false;

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

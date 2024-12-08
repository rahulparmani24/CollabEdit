// Get the WebSocket URL from environment variables or use a default
const backendUrl = process.env.BACKEND_API_URL || "ws://localhost:8080/ws";
const socket = new WebSocket(backendUrl);

// Get references to DOM elements
const editor = document.getElementById("editor");
const userList = document.getElementById("user-list");
const chatBox = document.getElementById("chat-box");
const sendButton = document.getElementById("send-button");

// Handle incoming WebSocket messages
socket.onmessage = (event) => {
  const message = JSON.parse(event.data);
  if (message.type === "update") {
    editor.innerText = message.content;
  } else if (message.type === "user_list") {
    updateUserList(message.users);
  }
};

// Send content updates to the server
editor.addEventListener("input", () => {
  const message = {
    type: "update",
    content: editor.innerText,
  };
  socket.send(JSON.stringify(message));
});

// Send chat messages
sendButton.addEventListener("click", () => {
  const chatMessage = chatBox.value;
  if (chatMessage.trim() !== "") {
    const message = {
      type: "chat",
      content: chatMessage,
    };
    socket.send(JSON.stringify(message));
    chatBox.value = "";
  }
});

// Update the active users list
function updateUserList(users) {
  userList.innerHTML = "";
  users.forEach((user) => {
    const li = document.createElement("li");
    li.textContent = user;
    userList.appendChild(li);
  });
}

// Handle WebSocket connection errors
socket.onerror = (error) => {
  console.error("WebSocket error:", error);
};

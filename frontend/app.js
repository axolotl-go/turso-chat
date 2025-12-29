const API_URL = "http://127.0.0.1:8080/api";
let currentRoomId = null;
let currentRoomName = null;

async function loadRooms() {
  try {
    const response = await fetch(`${API_URL}/rooms`);
    const rooms = await response.json();

    const roomsList = document.getElementById("roomsList");
    roomsList.innerHTML = "";

    rooms.forEach((room) => {
      const roomItem = document.createElement("div");
      roomItem.className = "room-item";
      if (room.ID === currentRoomId) {
        roomItem.classList.add("active");
      }
      roomItem.onclick = () => selectRoom(room.ID, room.name);
      roomItem.innerHTML = `<div class="room-name">${room.name}</div>`;
      roomsList.appendChild(roomItem);
    });
  } catch (error) {
    showError("Error al cargar las salas: " + error.message);
  }
}

async function createRoom() {
  const input = document.getElementById("roomNameInput");
  const name = input.value.trim();

  if (!name) return;

  try {
    const response = await fetch(`${API_URL}/rooms`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name }),
    });

    if (response.ok) {
      input.value = "";
      await loadRooms();
    }
  } catch (error) {
    showError("Error al crear la sala: " + error.message);
  }
}

async function selectRoom(roomId, roomName) {
  currentRoomId = roomId;
  currentRoomName = roomName;
  document.getElementById("currentRoomName").textContent = roomName;

  const roomItems = document.querySelectorAll(".room-item");
  roomItems.forEach((item) => item.classList.remove("active"));
  event.currentTarget.classList.add("active");

  await loadMessages(roomId);
}

async function loadMessages(roomId) {
  try {
    const response = await fetch(`${API_URL}/rooms/${roomId}/messages`);
    const messages = await response.json();

    const container = document.getElementById("messagesContainer");
    container.innerHTML = "";

    if (!messages || messages.length === 0) {
      container.innerHTML =
        '<div class="empty-state">No hay mensajes aún. ¡Sé el primero en escribir!</div>';
      return;
    }

    messages.forEach((msg) => {
      const messageDiv = document.createElement("div");
      messageDiv.className = "message received";
      messageDiv.innerHTML = `
                        <div class="message-sender">${msg.sender}</div>
                        <div class="message-content">${msg.content}</div>
                    `;
      container.appendChild(messageDiv);
    });

    container.scrollTop = container.scrollHeight;
  } catch (error) {
    showError("Error al cargar los mensajes: " + error.message);
  }
}

async function sendMessage(event) {
  event.preventDefault();

  if (!currentRoomId) {
    showError("Por favor selecciona una sala primero");
    return;
  }

  const messageInput = document.getElementById("messageInput");
  const content = messageInput.value.trim();

  if (!content) return;

  try {
    const response = await fetch(`${API_URL}/messages`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        room_id: currentRoomId,
        content: content,
      }),
    });

    if (response.ok) {
      messageInput.value = "";
      await loadMessages(currentRoomId);
    }
  } catch (error) {
    showError("Error al enviar el mensaje: " + error.message);
  }
}

function showError(message) {
  const errorContainer = document.getElementById("errorContainer");
  errorContainer.innerHTML = `<div class="error">${message}</div>`;
  setTimeout(() => {
    errorContainer.innerHTML = "";
  }, 5000);
}

loadRooms();
setInterval(() => {
  if (currentRoomId) {
    loadMessages(currentRoomId);
  }
}, 3000);

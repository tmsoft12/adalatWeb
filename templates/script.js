const messagesContainer = document.getElementById('messages');
const messageInput = document.getElementById('messageInput');
const sendButton = document.getElementById('sendButton');

let userId = '';
let socket;

// WebSocket baglanyşygyny dörediň
const connectWebSocket = () => {
    socket = new WebSocket(`ws://192.168.100.180:3000/api/chat/ws?user_id=${userId}`);

    socket.onopen = () => {
        console.log("WebSocket baglanyşygy açyldy.");
    };

    socket.onmessage = (event) => {
        const msg = JSON.parse(event.data);
        displayMessage(msg);
    };

    socket.onclose = () => {
        console.log("WebSocket baglanyşygy ýapyldy.");
    };
};

// Habary görkezmek
const displayMessage = (msg) => {
    const messageDiv = document.createElement('div');
    messageDiv.className = 'message';
    messageDiv.textContent = `${msg.content}`;
    messagesContainer.appendChild(messageDiv);
    messagesContainer.scrollTop = messagesContainer.scrollHeight;
    console.log()
};

// Habary ibermek
const sendMessage = () => {
    const content = messageInput.value.trim(); // Inputdan hatanyň boş bolmazlygyny üpjün ediň
    if (content && userId) {
        const msg = {
            user_id: userId,
            content: content,
            created_at: new Date().toISOString() // Habar wagty
        };
        socket.send(JSON.stringify(msg));
        messageInput.value = ''; // Input meýdançasy boşaldylýar
    }
};

// Ulanyjy ID-ni almak
const getUserId = async () => {
    try {
        const response = await fetch('http://192.168.100.180:3000/api/chat/me');
        if (!response.ok) throw new Error('Network response was not ok');
        const data = await response.json();
        userId = "d1227c44-ff9f-429d-a28e-f97a748dabd5";
        console.log(userId) // Ulanyjynyň ID-sini al
        connectWebSocket(); // WebSocket baglanyşygyny dörediň
    } catch (error) {
        console.error('Ulanyjy ID-sini almakda ýalňyşlyk:', error);
    }
};

// Obyektiň event listenerlerini düzüň
sendButton.addEventListener('click', sendMessage);
messageInput.addEventListener('keypress', (e) => {
    if (e.key === 'Enter') {
        sendMessage();
    }
});

// Ulanyjy ID-ni alanyňdan soň WebSocket baglanyşygyny dörediň
getUserId();

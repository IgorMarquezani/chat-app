import { useEffect, useState } from 'react';
import { Box } from '@radix-ui/themes';

let ws

const sendMessage = () => {
  const textInput = document.getElementById("text-input")
  if (textInput.value.trim().length == 0) {
    return
  }
  if (!ws) {
    return
  }

  try {
    ws.send(JSON.stringify({
      type: 1,
      data: textInput.value
    }))
  } catch {
    console.log("error while sending websocket message")
  }
}

const Chat = ({ user, chatId }) => {
  const [messages, setMessages] = useState([]);

  const [input, setInput] = useState('');

  useEffect(() => {
    if (ws) {
      ws.close()
    }

    const chatBox = document.getElementById("chat-box")
    chatBox.innerHTML = ""

    const proto = document.location.protocol === "https:" ? "wss" : "ws"
    ws = new WebSocket(`${proto}://${document.location.host}/api/ws/private/${chatId}/connect`)

    ws.onopen = () => {
      console.log("Connected to WebSocket")
    };

    ws.onmessage = (e) => {
      try {
        const data = JSON.parse(e.data);
        console.log("Message received:", data)
        setMessages((prev) => [...prev, data])
      } catch {
        console.log("Message received:", e.data)
      }
    };

    ws.onclose = () => {
      console.log("WebSocket closed")
    };

    ws.onerror = (err) => {
      console.error("WebSocket error:", err)
    };
  }, [chatId])

  return (
    <div className="flex flex-col h-full">

      <div id='chat-box' className="flex-1 overflow-y-auto p-4 space-y-3 bg-gray-50">
        {messages.map((msg) => (
          <div
            className={`max-w-xs px-3 py-2 rounded-lg ${msg.sender === user?.id
              ? 'bg-blue-500 text-white self-end ml-auto'
              : 'bg-gray-300 text-black'
              }`}
          >
            {msg.data}
          </div>
        ))}
      </div>

      <form onSubmit={(e) => { e.preventDefault(); sendMessage() }}>
        <Box maxWidth="1000px" className='p-3 bg-white flex gap-2 border-t'>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Type a message..."
            className="flex-1 rounded-full px-4 py-2 bg-gray-200"
            size="3"
            radius='full'
            id='text-input'
          >
          </input>
          <button
            className="bg-blue-500 text-white px-4 rounded-full hover:bg-blue-600"
          >
            Send
          </button>
        </Box>
      </form>
    </div>
  );
}

export default Chat

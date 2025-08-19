import { useState } from "react";
import Chat from "../components/Chat";
import ChatSideBar from '../components/ChatSideBar';
import Navbar from '../components/Navbar';

let mainChat

const initChat = () => {
  const id = 1
  const proto = document.location.protocol === "https:" ? "wss" : "ws"
  mainChat = new WebSocket(`${proto}://${document.location.host}/api/ws/private/${id}/connect`)

  mainChat.onopen = () => {
    console.log("Connected to WebSocket")
  };

  mainChat.onmessage = (e) => {
    try {
      const data = JSON.parse(e.data);
      console.log("Message received:", data)
    } catch {
      console.log("Message received:", e.data)
    }
  };

  mainChat.onclose = () => {
    console.log("WebSocket closed")
  };

  mainChat.onerror = (err) => {
    console.error("WebSocket error:", err)
  };
}

initChat()

const submitListener = () => {
  const input = document.getElementById("text-input")

  mainChat.send(JSON.stringify({ data: input.value }))
}

export default function Home() {
  const [chats, setChats] = useState([])

  return (
    <div className="h-screen flex bg-gray-100">
      <ChatSideBar chats={chats} setChats={setChats} />

      <div className="flex flex-col w-full h-full">
        <Navbar chats={chats} setChats={setChats} />
        <div className="flex-1 min-h-0">
          <Chat onClickSubmit={submitListener} />
        </div>
      </div>
    </div>
  );
}


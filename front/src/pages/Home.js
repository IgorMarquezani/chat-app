import { useEffect, useState } from "react";
import Chat from "../components/Chat";
import ChatSideBar from '../components/ChatSideBar';
import Navbar from '../components/Navbar';

export default function Home() {
  const [chats, setChats] = useState([])
  const [user, setUser] = useState({})
  const [chatId, setChatId] = useState('')

  useEffect(() => {
    fetch("/api/users/me", {
      method: "GET"
    }).then((resp) => {
      return resp.json()
    }).then((json) => {
      console.log(json)
      setUser(json?.data)
      setChatId(json.data?.last_chat_id)
    }).catch((err) => {
      console.log(err)
    })
  }, {})

  return (
    <div className="h-screen flex bg-gray-100">
      <ChatSideBar user={user} setChatId={setChatId} chats={chats} setChats={setChats} />

      <div className="flex flex-col w-full h-full">
        <Navbar user={user} chats={chats} setChats={setChats} />
        <div className="flex-1 min-h-0">
          <Chat user={user} chatId={chatId} />
        </div>
      </div>
    </div>
  );
}


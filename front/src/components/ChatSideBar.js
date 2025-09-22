import { UserIcon } from "lucide-react"
import { useEffect } from "react"

const ChatSideBar = ({ user, setChatId, chats, setChats }) => {
  useEffect(() => {
    fetch(`/api/private/chat/list`, {
      method: "GET",
    })
      .then((resp) => resp.json())
      .then((json) => {
        if (json?.data) {
          console.log(json.data)
          setChats(json.data);
        } else {
          setChats([]);
        }
      })
      .catch(() => setChats([]))

  }, [])

  return (
    <div className="w-64 bg-white border-r flex flex-col overflow-y-auto">
      <h1 className="flex border-b p-3 font-bold gap-1"><UserIcon />{user.name}</h1>
      <h1 className="p-3 font-bold border-b">Contatos:</h1>
      {chats.map((chat) => (
        <div className="flex p-3 border-b hover:cursor-pointer" onClick={() => setChatId(chat.chat_id)}>
          <h1 key={chat.chat_id}>{chat.friend_name}</h1>
        </div>
      ))}
    </div>
  )
}

export default ChatSideBar

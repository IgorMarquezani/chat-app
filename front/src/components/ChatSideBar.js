import { useEffect } from "react"

const startWebsocket = () => {
}

const ChatSideBar = ({ chats, setChats }) => {
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
      <h1 className="p-3 font-bold border-b">Contatos</h1>
      {chats.map((chat, i) => (
        <div className="flex p-3 border hover:cursor-pointer" onClick={() => startWebsocket()}>
          <h1 key={chat.chat_id}>{chat.friend_name}</h1>
        </div>
      ))}
    </div>
  )
}

export default ChatSideBar

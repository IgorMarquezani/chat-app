import { useState } from 'react';
import { Box, Avatar, Flex } from '@radix-ui/themes';
import Navbar from '../components/Navbar';

const Chat = () => {
  const [messages, setMessages] = useState([
    { id: 1, text: 'Hey there!', sender: 'other' },
    { id: 2, text: 'Hi! How are you?', sender: 'me' },
  ]);

  const [input, setInput] = useState('');

  const sendMessage = () => {
    if (!input.trim()) return;
    setMessages([...messages, { id: Date.now(), text: input, sender: 'me' }]);
    setInput('');
  };

  return (
    <div className="flex h-screen bg-gray-100">
      {/* LEFT SIDEBAR */}
      <div className="w-64 bg-white border-r flex flex-col overflow-y-auto">
        <h1 className="p-3 font-bold border-b">Contatos</h1>
        <div className="flex-1 p-3 space-y-2">
          <Flex>
            <Avatar fallback="AA" />
            <h1>Contato 1</h1>
          </Flex>
        </div>
      </div>

      {/* CHAT AREA */}
      <div className="flex flex-col flex-1">
        {/* Header */}
        <div className="bg-gray-900 text-white p-3 font-bold"><Navbar /></div>

        {/* Messages */}
        <div className="flex-1 overflow-y-auto p-4 space-y-3 bg-gray-50">
          {messages.map((msg) => (
            <div
              key={msg.id}
              className={`max-w-xs px-3 py-2 rounded-lg ${msg.sender === 'me'
                ? 'bg-blue-500 text-white self-end ml-auto'
                : 'bg-gray-300 text-black'
                }`}
            >
              {msg.text}
            </div>
          ))}
        </div>

        {/* Input */}
        <Box maxWidth="1000px" className='p-3 bg-white flex gap-2 border-t'>
          <input
            type="text"
            value={input}
            onChange={(e) => setInput(e.target.value)}
            placeholder="Type a message..."
            className="flex-1 rounded-full px-4 py-2 bg-gray-200"
            size="3"
            radius='full'
          >
          </input>
          <button
            onClick={sendMessage}
            className="bg-blue-500 text-white px-4 rounded-full hover:bg-blue-600"
          >
            Send
          </button>
        </Box>
      </div>
    </div>
  );
}

export default Chat

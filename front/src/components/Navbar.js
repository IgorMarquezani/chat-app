import { useState, useEffect } from "react";
import * as NavigationMenu from "@radix-ui/react-navigation-menu";
import * as Popover from "@radix-ui/react-popover";
import { Link } from "react-router-dom";
import { GearIcon } from "@radix-ui/react-icons";
import { Plus } from "lucide-react";


const createNewChat = (setChats) => {
  const input = document.getElementById("user-search")
  const userID = input.getAttribute("userid")

  fetch("/api/private/chat/create", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ target_id: parseInt(userID) })
  }).then((resp) => {
    if (resp.status === 208) {
      return {}
    }

    return resp.json()
  }).then((json) => {
    if (json?.data && json.data.chat_id?.length > 0) {
      console.log(json.data)
      setChats((prevChats) => [json.data, ...prevChats,])
    }
  }).catch((e) => {
    console.log(e)
  })
}

export default function Navbar({ setChats }) {
  const [userName, setUserName] = useState("");
  const [userID, setUserID] = useState(0);
  const [open, setOpen] = useState(false);
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (userName.trim().length < 1 && userName.length === 0) {
      setUsers([]);
      setOpen(false);
      return;
    }

    setLoading(true);

    fetch(`/api/users/search/${userName}`, {
      method: "GET",
    })
      .then((resp) => resp.json())
      .then((json) => {
        if (json?.data) {
          setUsers(json.data);
        } else {
          setUsers([]);
        }
      })
      .catch(() => setUsers([]))
      .finally(() => setLoading(false));
  }, [userName]);

  return (
    <div className="bg-gray-900 text-white p-3 font-bold">
      <NavigationMenu.Root className="bg-gray-900 text-white p-3 flex gap-4">

        <div className="relative w-100 ms-auto me-auto">
          <div className="flex gap-2">
            <input
              id="user-search"
              placeholder="Search for friends"
              value={userName}
              userid={userID}
              onChange={(e) => {
                setUserName(e.target.value);
                setOpen(e.target.value.length > 0);
              }}
              className="w-full rounded-full px-3 py-1 bg-white text-black"
            >
            </input>
            <div className="mt-auto mb-auto bg-white rounded hover:bg-green-400 focus:bg-green-400" onClick={() => createNewChat(setChats)}>
              <Plus className="text-black hover:text-white hover:cursor-pointer transition" />
            </div>
          </div>

          {open && (
            <div className="absolute left-0 right-0 mt-1 bg-white text-black rounded-lg shadow-lg z-10">
              {loading ? (
                <div className="p-2 text-gray-500">Loading...</div>
              ) : users.length > 0 ? (
                <ul className="divide-y divide-gray-200">
                  {users.map((user, i) => (
                    i < 4 ?
                      <li
                        key={i}
                        className="p-2 hover:bg-gray-100 cursor-pointer"
                        onClick={() => {
                          setUserName(user.name);
                          setUserID(user.id);
                          setOpen(false);
                        }}
                      >
                        {user.name}
                        <p className="text-sm text-gray-500">id: {user.id}</p>
                      </li> : {}
                  ))}
                </ul>
              ) : (
                <div className="p-2 text-gray-500">No users found</div>
              )}
            </div>
          )}
        </div>

        <NavigationMenu.List clasName="flex gap-4">
          <NavigationMenu.Item>
            <Link to="/about" className="hover:underline">
              About
            </Link>
          </NavigationMenu.Item>
        </NavigationMenu.List>

        {/* Config Popover */}
        <Popover.Root>
          <Popover.Trigger asChild>
            <button className="rounded-full hover:bg-gray-700 transition">
              <GearIcon className="w-5 h-5" />
            </button>
          </Popover.Trigger>

          <Popover.Content
            className="bg-white text-black rounded-lg shadow-lg p-3 w-48"
            sideOffset={8}
          >
            <h3 className="font-bold mb-2">Settings:</h3>
            <ul className="space-y-2">
              <li>
                <button className="hover:underline">Profile</button>
              </li>
              <li>
                <button className="hover:underline">Logout</button>
              </li>
            </ul>
            <Popover.Arrow className="fill-white" />
          </Popover.Content>
        </Popover.Root>
      </NavigationMenu.Root>
    </div>
  );
}

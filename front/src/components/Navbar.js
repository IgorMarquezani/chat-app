import * as NavigationMenu from '@radix-ui/react-navigation-menu';
import * as Popover from '@radix-ui/react-popover';
import { TextField } from '@radix-ui/themes';
import { Link } from 'react-router-dom';
import { GearIcon } from '@radix-ui/react-icons';


export default function Navbar() {
  return (
    <NavigationMenu.Root className="bg-gray-900 text-white p-3 flex gap-4">
      <NavigationMenu.List className="flex gap-4">
        <NavigationMenu.Item>
          <Link to="/about" className="hover:underline">About</Link>
        </NavigationMenu.Item>
      </NavigationMenu.List>
      <input
        placeholder="Search for friends"
        className="w-72 ms-auto me-auto rounded-full px-3 py-1 bg-white text-black"
      >
      </input>
      <Popover.Root>
        <Popover.Trigger asChild>
          <button className="p-2 rounded-full hover:bg-gray-700 transition">
            <GearIcon className="w-5 h-5" />
          </button>
        </Popover.Trigger>

        <Popover.Content
          className="bg-white text-black rounded-lg shadow-lg p-3 w-48"
          sideOffset={8}
        >
          <h3 className="font-bold mb-2">Settings</h3>
          <ul className="space-y-2">
            <li><button className="hover:underline">Profile</button></li>
            <li><button className="hover:underline">Account</button></li>
            <li><button className="hover:underline">Logout</button></li>
          </ul>
          <Popover.Arrow className="fill-white" />
        </Popover.Content>
      </Popover.Root>
    </NavigationMenu.Root>
  );
}


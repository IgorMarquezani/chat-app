import Chat from "../components/Chat";

export default function Home() {
  const cookies = document.cookie.trim(";")

  if (cookies) {
    return (<Chat />)
  }

  document.location = "/login"
}


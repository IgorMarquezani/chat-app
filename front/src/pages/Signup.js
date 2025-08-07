import Label from "../components/Label"
import Input from "../components/Input"
import "./Signup.css"
import { useState } from "react"
import { Link } from "react-router"

const submitHandler = (e) => {
  e.preventDefault()

  console.log("preventedDefault")
}

const Signup = () => {
  const [name, setName] = useState()
  const [email, setEmail] = useState()
  const [password, setPassword] = useState()
  const [cpassword, setCPassword] = useState()

  return (
    <div className="w-screen flex justify-start align-middle main-div">
      <div className="bg-white flex flex-col w-full h-screen justify-center align-middle max-w-md p-6">
        <form className="flex flex-col gap-4 p-5" onChange={(e) => submitHandler(e)}>
          <h1 className="font-extrabold text-3xl text-center">Sign up</h1>
          <Label htmlFor="name">Name:</Label>
          <Input
            type="text"
            value={name}
            id="name"
            required={true}
            min={2} max={55}
            placeholder="Type your name"
            onChange={setName}
          />
          <Label htmlFor="email">Email:</Label>
          <Input
            type="email"
            value={email}
            id="email"
            required={true}
            placeholder="Type your email"
            onChange={setEmail}
          />
          <Label htmlFor="password">Password:</Label>
          <Input
            type="password"
            id="password"
            value={password}
            required={true}
            min={8} max={40}
            onChange={setPassword}
            placeholder="Type your password"
          />
          <Label htmlFor="confirm-password">Password:</Label>
          <Input
            type="password"
            id="confirm-password"
            value={cpassword}
            required={true}
            min={8} max={40}
            onChange={setCPassword}
            placeholder="Type your password again"
          />
          <p className="font-thin text-center">Already Signed up? <Link to="/login" className="text-blue-500 underline">Log in</Link></p>
          <div className="text-center">
            <button
              type="submit"
              className="bg-green-400 text-white mt-2 px-6 py-2 rounded hover:bg-green-500 transition"
            >
              Submit
            </button>
          </div>
        </form>

      </div>

      <div className="w-full h-screen flex flex-col justify-center text-center">
        <h1 className="text-4xl font-extrabold text-white main-message">Talk to your friends anywhere, any time</h1>
      </div>
    </div>
  )
}

export default Signup

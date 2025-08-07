import { useState } from "react";
import Label from "../components/Label";
import Input from "../components/Input";
import "./Login.css"
import { Link } from "react-router";

const submitHandler = (e) => {
  e.preventDefault()

  console.log("preventedDefault")
}

const Login = () => {
  const [name, setName] = useState()
  const [password, setPassword] = useState()

  return (
    <div className="flex items-center justify-center h-screen w-screen bg-gray-900 main-div">
      <div className="rounded-md bg-white w-full max-h-screen max-w-md p-6">
        <form className="flex flex-col gap-4 p-5" onSubmit={(e) => submitHandler(e)}>
          <h1 className="text-center text-3xl font-extrabold">Login</h1>

          <div className="flex flex-col">
            <Label htmlFor="name">Name:</Label>
            <Input
              type="text"
              required={true}
              value={name}
              onChange={setName}
              min={2}
              max={55}
              id="name"
            />
          </div>

          <div className="flex flex-col">
            <Label htmlFor="password" placeholder="Type your password">Password:</Label>
            <Input
              type="password"
              required={true}
              value={password}
              onChange={setPassword}
              id="password"
            />
          </div>

          <p className="text-center font-thin">Don't have an account? <Link to="/signup" className="text-blue-600 underline">Sign up</Link></p>

          <div className="text-center mt-2">
            <button
              type="submit"
              className="bg-green-400 text-white px-6 py-2 rounded hover:bg-green-500 transition"
            >
              Submit
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;


import { useState } from "react";
import Label from "../components/Label";
import Input from "../components/Input";
import BtnSuccess from "../components/BtnSuccess";
import "./Login.css"
import { Link } from "react-router";
import { LoaderIcon } from "lucide-react";

const submitHandler = (e, setPending) => {
  e.preventDefault()

  const emailError = document.getElementById("email-error")
  emailError.innerHTML = ""

  const email = document.getElementById("email").value
  const password = document.getElementById("password").value

  const body = {
    email,
    password,
  }

  setPending(true)

  let logged = false

  fetch("/api/users/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(body),
    credentials: "include"
  })
    .then(async (resp) => {
      if (resp.status === 200 || resp.status === 208) {
        logged = true;
      } else if (resp.status === 400) {
        emailError.innerHTML = "invalid credentials";
      } else if (resp.status === 500) {
        emailError.innerHTML = "internal server error";
      } else {
        emailError.innerHTML = "unexpected error";
      }

      return {
        status: resp.status,
        txt: await resp.text()
      };
    })
    .then((data) => {
      console.log(`api response status: ${data.status}`);
      console.log("api message: " + data.txt);
    })
    .catch((err) => {
      console.error("Fetch error:", err);
      emailError.innerHTML = "network error";
    })
    .finally(() => {
      setPending(false);
      if (logged) {
        document.location = "/";
      }
    });
}

const Login = () => {
  const [pending, setPending] = useState()
  const [email, setEmail] = useState()
  const [password, setPassword] = useState()

  return (
    <div className="flex items-center justify-center h-screen w-screen bg-gray-900 main-div">
      <div className="rounded-md bg-white w-full max-h-screen max-w-md p-6">
        <form className="flex flex-col gap-1 p-5" onSubmit={(e) => submitHandler(e, setPending)}>
          <h1 className="text-center text-3xl font-extrabold">Log in</h1>
          <Label htmlFor="name">E-mail:</Label>
          <Input
            type="email"
            required={true}
            value={email}
            onChange={setEmail}
            min={2}
            max={55}
            id="email"
          />
          <p id="email-error" className="text-red-500"></p>
          <Label htmlFor="password" placeholder="Type your password">Password:</Label>
          <Input
            type="password"
            required={true}
            value={password}
            onChange={setPassword}
            id="password"
          />

          <p className="text-center font-thin">Don't have an account? <Link to="/signup" className="text-blue-600 underline">Sign up</Link></p>

          <BtnSuccess disabled={pending} disabledInnerHTML={<LoaderIcon className="animate-spin" />} />
        </form>
      </div>
    </div>
  );
};

export default Login;


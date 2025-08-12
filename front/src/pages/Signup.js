import Label from "../components/Label"
import Input from "../components/Input"
import BtnSuccess from "../components/BtnSuccess"
import "./Signup.css"
import { Link } from "react-router"
import { CheckCircle, LoaderIcon, XCircle } from "lucide-react";
import { useState } from "react"

const detectNameError = (name) => {
  if (name.trim(" ").length < 3) {
    return "empty name"
  }
  if (name.length < 3) {
    return "name should be at least 3 chars long"
  }
  if (name.length > 55) {
    return "name should have be only 55 chars long"
  }
}

const validateName = (name) => {
  if (name.trim(" ").length < 3) {
    return false
  }
  if (name.length < 3) {
    return false
  }
  if (name.length > 55) {
    return false
  }

  return true
}

const validatePassword = (passwd, confirm) => {
  return passwd === confirm
}

const submitHandler = async (e, setPending) => {
  e.preventDefault()

  const emailError = document.getElementById("email-error")
  emailError.innerHTML = ""

  const nameError = document.getElementById("name-error")
  nameError.innerHTML = ""

  const name = document.getElementById("name").value
  const email = document.getElementById("email").value
  const passwd = document.getElementById("password").value
  const confirm = document.getElementById("confirm-password").value

  if (!validateName(name)) {
    nameError.innerHTML = detectNameError(name)
    return
  }
  if (!validatePassword(passwd, confirm)) {
    return
  }

  const body = {
    name: name,
    email: email,
    password: passwd
  }


  fetch("/api/users/signup", {
    method: "POST",
    body: JSON.stringify(body)
  }).then(async (resp) => {
    setPending(true)

    if (resp.status === 200) {
      document.location = "/login"
      return
    }
    if (resp.status === 208) {
      emailError.innerHTML = "e-mail already in use"
      return
    }

    return {
      status: resp.status,
      txt: await resp.text()
    }
  }).then((data) => {
    if (data !== undefined) {
      console.log("unexpected error")
      console.log("error: " + data.txt)
      console.log("status: " + data.status.toString())
    }
  }).finally(() => {
    setPending(false)
  })
}

const Signup = () => {
  const [pending, setPending] = useState()
  const [name, setName] = useState()
  const [email, setEmail] = useState()
  const [password, setPassword] = useState()
  const [cpassword, setCPassword] = useState()

  const requirements = [
    { label: "At least 8 characters", valid: password !== undefined && password.length > 7 },
    { label: "Contains a number", valid: /\d/.test(password) },
    { label: "Contains an uppercase letter", valid: /[A-Z]/.test(password) },
    { label: "Contains a special character", valid: /[^A-Za-z0-9]/.test(password) },
  ];

  const passwordsMatch = password !== cpassword

  return (
    <div className="w-screen flex justify-start align-middle main-div">
      <div className="bg-white flex flex-col w-full h-screen justify-center align-middle max-w-md p-6">
        <form className="flex flex-col gap-1 p-5" onSubmit={(e) => submitHandler(e, setPending)}>
          <h1 className="font-extrabold text-3xl text-center">Sign up</h1>
          <Label htmlFor="name">Name:</Label>
          <Input
            type="text"
            value={name}
            id="name"
            min={2} max={55}
            placeholder="Type your name"
            onChange={setName}
          />
          <p id="name-error" className="text-red-500 m-0"></p>
          <Label htmlFor="email">E-mail:</Label>
          <Input
            type="email"
            value={email}
            id="email"
            required={true}
            placeholder="Type your email"
            onChange={setEmail}
          />
          <p id="email-error" className="text-red-500 m-0"></p>
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
          <ul className="mt-2">
            {requirements.map((req, i) => (
              <li key={i} className="flex items-center">
                {
                  req.valid ? (
                    <CheckCircle size={20} className="text-green-400" />
                  ) : (
                    <XCircle size={20} className="text-red-400" />
                  )
                }
                <span className="ml-2">{req.label}</span>
              </li>
            ))}
          </ul>
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
          <p className={passwordsMatch ? "text-red-500 m-0" : "m-0"}>{passwordsMatch ? "Passwords don't match" : ""}</p>
          <BtnSuccess disabled={pending} disabledInnerHTML={<LoaderIcon className="animate-spin" />} />
          <p className="font-thin text-center">Already Signed up? <Link to="/login" className="text-blue-500 underline">Log in</Link></p>
        </form>
      </div>

      <div className="w-full h-screen flex flex-col justify-center text-center">
        <h1 className="text-4xl font-extrabold text-white main-message">Talk to your friends anywhere, any time</h1>
      </div>
    </div>
  )
}

export default Signup

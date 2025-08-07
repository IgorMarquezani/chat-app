const Input = ({ type, required, value, placeholder, onChange, min, max, id }) => {
  return (
    <input
      type={type}
      required={required ? true: false}
      value={value}
      placeholder={placeholder}
      onChange={(e) => onChange(e.target.value)}
      min={min}
      max={max}
      id={id}
      className="w-full px-4 py-2 rounded-md bg-gray-200 border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500"
    />

  )
}

export default Input

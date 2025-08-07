const Label = ({ htmlFor, children }) => {
  return (
    <label className="font-bold mb-2" htmlFor={htmlFor}>{children}</label>
  )
}

export default Label

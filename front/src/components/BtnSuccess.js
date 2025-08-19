const BtnSuccess = ({ disabled, disabledInnerHTML, innerHTML, onClick }) => {
  if (disabledInnerHTML === undefined) {
    disabledInnerHTML = "Submit"
  }
  if (innerHTML === undefined) {
    innerHTML = "Submit"
  }

  return (
    <div className="text-center">
      <button
        type="submit"
        className="bg-green-400 text-white mt-2 px-6 py-2 rounded hover:bg-green-500 transition"
        disabled={disabled}
        onClick={onClick}
      >
        {disabled ? disabledInnerHTML : innerHTML}
      </button>
    </div>

  )
}

export default BtnSuccess

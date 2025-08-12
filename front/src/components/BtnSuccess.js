const BtnSuccess = ({ disabled, disabledInnerHTML }) => {
  if (disabledInnerHTML === undefined) {
    disabledInnerHTML = "Submit"
  }

  return (
    <div className="text-center">
      <button
        type="submit"
        className="bg-green-400 text-white mt-2 px-6 py-2 rounded hover:bg-green-500 transition"
        disabled={disabled}
      >
        {disabled ? disabledInnerHTML : "Submit"}
      </button>
    </div>

  )
}

export default BtnSuccess

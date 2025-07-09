

const Input = (props) => {
   const { type, placeholder, value, onChange, label, multiple = false } = props;

   return (
      <div>
         <label htmlFor={label} className="block text-secondary text-lg mb-2">{label}</label>
         <input
            id={label}
            type={type}
            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-primary"
            placeholder={placeholder}
            value={value}
            onChange={onChange}
            multiple={multiple}
         />
      </div>
   )
}

export default Input;
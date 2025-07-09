

const TextArea = (props) => {
   const { placeholder, value, onChange, label } = props;
   return (
      <div>
         <label htmlFor={label} className="block text-secondary text-lg mb-2">{label}</label>
         <textarea 
            id={label}
            className="w-full mt-4 p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-primary"
            type="text"
            placeholder={placeholder}
            value={value}
            onChange={onChange}
            rows="6"
         ></textarea>
      </div>
   )
}

export default TextArea;
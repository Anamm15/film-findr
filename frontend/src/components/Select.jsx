

const Select = (props) => {
   const { label, value, onChange, options } = props;
   return (
      <div className="">
         <label htmlFor={label} className="block text-gray-600 text-sm mb-2">{label}</label>
         <select
            id={label}
            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-primary"
            value={value}
            // defaultValue={genres && genres[0]?.id}
            onChange={onChange}
         >
         {
            options && options.map((option) => (
               <option key={option.id} value={option.id}>{option.nama}</option>
            ))
         }
         </select>
      </div>
   )
}

export default Select;
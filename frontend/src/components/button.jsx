

const Button = (props) => {
    const {children, className, type} = props;
    return (
        <button 
            type={type}
            className={`font-semibold text-white bg-primary duration-150 hover:scale-[98%] hover:bg-secondary px-4 py-1.5 mb-2 ${className}`}                
        >
            {children}
        </button>
    )
}

export default Button
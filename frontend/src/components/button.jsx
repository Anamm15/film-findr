

const Button = (props) => {
    const {children, className, type, onClick} = props;
    return (
        <button 
            type={type}
            onClick={onClick}
            className={`font-semibold text-white bg-primary duration-150 hover:scale-[98%] hover:bg-secondary px-4 py-1.5 ${className}`}                
        >
            {children}
        </button>
    )
}

export default Button
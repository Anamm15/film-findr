

const Button = (props) => {
    const {children, className, type, onClick} = props;
    return (
        <button 
        type={type}
            onClick={onClick}
            className={`
                font-semibold text-white
                px-4 py-1.5 rounded-full
                bg-gradient-primary hover:bg-gradient-secondary
                transition-all duration-200 ease-in-out transform
                hover:scale-[98%]
                ${className}
            `}
            >
            {children}
        </button>
    )
}

export default Button
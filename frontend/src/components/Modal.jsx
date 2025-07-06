import { X } from "react-feather";

const Modal = (props) => {
    const { onClick, children } = props
    return (
        <div className="fixed z-50 top-0 left-0 flex justify-center items-center w-screen h-screen">
            <div className="w-1/4 max-h-[80vh] bg-white shadow-lg relative p-4 rounded-md">
                <X className="absolute top-1 right-1 cursor-pointer" onClick={onClick}/>
                {children}
            </div>
        </div>
    )
}

export default Modal;
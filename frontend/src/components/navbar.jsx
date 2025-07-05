import { logoutUser } from "../service/user";
import { Link } from "react-router-dom";

const Navbar = () => {

    const hanldeLogout = async () => {
        try {
            const response = await logoutUser();
            console.log(response);
            
            if (response.status === 200) {
                localStorage.removeItem("token");
                window.location.reload();
            }
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <nav className="fixed z-50 w-screen bg-primary flex justify-between h-20 text-white items-center shadow-md top-0 left-0 px-10">
            <div className="text-2xl font-bold cursor-pointer">
                <a href="/">Film-Findr</a>
            </div>
            <div className="flex gap-10 font-semibold text-xl">
                <a href="/">Top Film</a>
                <a href="/menu">Profile</a>
                <a href="">Watch List</a>
            </div>
            <div>
                {
                    localStorage.getItem("token") ? (
                        <button 
                            onClick={hanldeLogout}
                            className="bg-secondary w-28 py-2 rounded-3xl font-semibold hover:bg-tertiary duration-150">Logout</button>
                    ) : (
                        <div className="flex gap-5">
                            <button className="bg-secondary w-28 py-2 rounded-3xl font-semibold hover:bg-tertiary duration-150">
                                <Link to="/login">Login</Link>
                            </button>
                            <button className="bg-secondary w-28 py-2 rounded-3xl font-semibold hover:bg-tertiary duration-150">
                                <Link to="/register">Register</Link>
                            </button>
                        </div>
                    )
                }
            </div>
        </nav>
    )
}


export default Navbar;
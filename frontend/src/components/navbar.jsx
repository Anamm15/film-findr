import { logoutUser } from "../service/user";
import { Link } from "react-router-dom";
import { useContext } from "react";
import { AuthContext } from "../contexts/authContext";

const Navbar = () => {
    const { user, loading } = useContext(AuthContext);

    if (loading) return null;

    const hanldeLogout = async () => {
        try {
            const response = await logoutUser();
            console.log(response);
            
            if (response.status === 200) {
                window.location.reload();
            }
        } catch (error) {
            console.log(error);
        }
    }

    return (
        <nav className="fixed z-50 w-screen bg-primary flex justify-between h-20 text-white items-center shadow-md top-0 left-0 px-10">
            <div className="text-2xl font-bold cursor-pointer">
                <Link to="/">Film-Findr</Link>
            </div>
            <div className="flex gap-10 font-semibold text-xl">
                <Link to="/">Top Film</Link>
                <Link to={`/profile/${user?.id}`}>Profile</Link>
                <Link to="/watchlist">Watch List</Link>
            </div>
            <div>
                {
                    user ? (
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
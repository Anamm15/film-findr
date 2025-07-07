import { logoutUser } from "../service/user";
import { Link } from "react-router-dom";
import { useContext } from "react";
import { AuthContext } from "../contexts/authContext";
import Button from "./Button";

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
        <nav className="fixed z-50 w-screen bg-background border-b border-[#e0e0e0] h-20 shadow-md top-0 left-0 flex justify-center items-center">
            <div className="flex justify-between items-center w-full xl:w-[1280px] p-4">
                <div className="text-3xl font-bold cursor-pointer text-primary">
                    <Link to="/">Film-Findr</Link>
                </div>
                <div className="flex gap-10 font-bold text-xl text-primary">
                    <Link to="/">Top Film</Link>
                    <Link to={`/profile/${user?.id}`}>Profile</Link>
                    <Link to="/watchlist">Watch List</Link>
                </div>
                <div>
                    {
                        user ? (
                            <Button 
                                onClick={hanldeLogout}
                                className="w-28 py-2 rounded-3xl font-semibold">Logout</Button>
                        ) : (
                            <div className="flex gap-5">
                                <Button className="w-28 py-2 rounded-3xl font-semibold">
                                    <Link to="/login">Login</Link>
                                </Button>
                                <Button className="w-28 py-2 rounded-3xl font-semibold">
                                    <Link to="/register">Register</Link>
                                </Button>
                            </div>
                        )
                    }
                </div>
            </div>
        </nav>
    )
}


export default Navbar;
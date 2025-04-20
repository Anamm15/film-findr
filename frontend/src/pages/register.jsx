import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { registerUser } from "../api/user";

const RegisterPage = () => {
    const [nama, setNama] = useState("");
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    const handleRegister = async (e) => {
        e.preventDefault();
        try {
            const data = {
                nama : nama,
                username : username,
                password : password
            }

            const response = await registerUser(data);
            if (response.status === 201) {
                setMessage(response.data.message);
                navigate("/login");
            }
        } catch (error) {
            setMessage(error.response.data.message);
        }
    };

    return (
        <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="bg-white p-8 rounded-lg shadow-lg w-96">
                <h2 className="text-2xl font-semibold text-center text-gray-700 mb-6">Login</h2>
                <form onSubmit={handleRegister}>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Nama</label>
                        <input
                            type="text"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Enter your name"
                            value={nama}
                            onChange={(e) => setNama(e.target.value)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Username</label>
                        <input
                            type="text"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Enter your username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                        />
                    </div>
                    <div className="mb-2">
                        <label className="block text-gray-600 text-sm mb-2">Password</label>
                        <input
                            type="password"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Enter your password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </div>
                    <p>Sudah punya akun? <Link to="/login" className="text-blue-500">Login</Link> </p>
                    <button
                        type="submit"
                        className="mt-2 w-full bg-blue-500 hover:bg-blue-700 text-white py-2 rounded transition duration-200"
                    >
                        Register
                    </button>
                    {message && <p className="text-red-500 mt-2">{message}</p>}
                </form>
            </div>
        </div>
    );
};

export default RegisterPage;

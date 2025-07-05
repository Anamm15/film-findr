import { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import { createFilm } from "../../service/film";
import { getAllGenre } from "../../service/genre";

const DashboardPage = () => {
    const [judul, setJudul] = useState("");
    const [status, setStatus] = useState("");
    const [sinopsis, setSinopsis] = useState("");
    const [sutradara, setSutradara] = useState("");
    const [tanggalRilis, setTanggalRilis] = useState("");
    const [durasi, setDurasi] = useState(null);
    const [totalEpisode, setTotalEpisode] = useState(null);
    const [imageFiles, setImageFiles] = useState([]);
    const [genres, setGenres] = useState([]);
    const [selectedGenres, setSelectedGenres] = useState([]);
    const [message, setMessage] = useState("");
    const [isGenresLoaded, setIsGenresLoaded] = useState(false);


    useEffect(() => {
        const fetchGenres = async () => {
            try {
                const response = await getAllGenre();
                setGenres(response.data.data);
            } catch (error) {
                console.error("Error fetching genres:", error.data.message);
            }
        };

        fetchGenres();
        if (!isGenresLoaded) {
            setIsGenresLoaded(true);
        }
    }, [isGenresLoaded]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const formData = new FormData();
            formData.append("judul", judul);
            formData.append("status", status);
            formData.append("sinopsis", sinopsis);
            formData.append("sutradara", sutradara);
            formData.append("tanggal_rilis", tanggalRilis);
            formData.append("durasi", durasi);
            formData.append("total_episode", totalEpisode);
            selectedGenres.forEach((g) => formData.append("genres", g));

            for (let i = 0; i < imageFiles.length; i++) {
                formData.append("images", imageFiles[i]);
            }

            const response = await createFilm(formData);
            if (response.status === 200) {
                setMessage(response.data.message);
            }
        } catch (error) {
            setMessage(error.data.message);
        }
    }

    return (
        <>
            <div className="flex items-center justify-center min-h-screen bg-gray-100">
            <div className="bg-white p-8 rounded-lg shadow-lg w-96">
                <h2 className="text-2xl font-semibold text-center text-gray-700 mb-6">Login</h2>
                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Judul</label>
                        <input
                            type="text"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Masukkan Judul"
                            value={judul}
                            onChange={(e) => setJudul(e.target.value)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Status</label>
                        <input
                            type="text"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Pilih Status"
                            value={status}
                            onChange={(e) => setStatus(e.target.value)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Sinopsis</label>
                        <input
                            type="text"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Deskripsikan sinosis"
                            value={sinopsis}
                            onChange={(e) => setSinopsis(e.target.value)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Sutradara</label>
                        <input
                            type="text"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Masukkan Sutradara"
                            value={sutradara}
                            onChange={(e) => setSutradara(e.target.value)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Tanggal Rilis</label>
                        <input
                            type="date"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Masukkan tanggal rilis"
                            value={tanggalRilis}
                            onChange={(e) => setTanggalRilis(e.target.value)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Total Episode</label>
                        <input
                            type="number"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Masukkan Total Episode"
                            value={totalEpisode}
                            onChange={(e) => setTotalEpisode(e.target.value)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Durasi</label>
                        <input
                            type="number"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Masukkan Total Episode"
                            value={durasi}
                            onChange={(e) => setDurasi(e.target.value)}
                        />
                    </div>
                    <div className="mb-2">
                        <label className="block text-gray-600 text-sm mb-2">Masukkan Gambar Pendukung</label>
                        <input
                            type="file"
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            placeholder="Enter your password"
                            multiple
                            onChange={(e) => setImageFiles(e.target.files)}
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-600 text-sm mb-2">Pilih genre</label>
                        <select
                            className="w-full p-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-400"
                            value={selectedGenres}
                            defaultValue={genres && genres[0].id}
                            onChange={(e) => setSelectedGenres(e.target.value)}
                        >
                            {
                                genres && genres.map((genre) => (
                                    <option key={genre.id} value={genre.id}>{genre.nama}</option>
                                ))
                            }
                        </select>
                    </div>
                    <p>Sudah punya akun? <Link to="/login" className="text-blue-500">Login</Link> </p>
                    <button
                        type="submit"
                        className="mt-2 w-full bg-blue-500 hover:bg-blue-700 text-white py-2 rounded transition duration-200"
                    >
                        Submit
                    </button>
                    {message && <p className="text-red-500 mt-2">{message}</p>}
                </form>
            </div>
        </div>
        </>
    )
}

export default DashboardPage;
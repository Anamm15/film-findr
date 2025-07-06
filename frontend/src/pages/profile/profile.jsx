import { useState, useEffect, useContext } from "react";
import { useParams } from "react-router-dom";
import { getReviewByUserId } from "../../service/review";
import { getUserById } from "../../service/user";
import Review from "../detailFilm/components/Review";
import { AuthContext } from "../../contexts/authContext";
import { getUserFilmByUserId } from "../../service/userFilm";
import Informasi from "./components/Informasi";
import WatchlistLayout from "../../layouts/WatchlistLayout";

const ProfilePage = () => {
    const params = useParams();
    const routeId = params.id;
    const { user: currentUser } = useContext(AuthContext);

    const [user, setUser] = useState(null);
    const [review, setReview] = useState([]);
    const [watchlists, setWatchlists] = useState([]);
    const [page, setPage] = useState(1);

    const finalId = routeId || currentUser?.id;

    useEffect(() => {
        const fetchUser = async () => {
            if (!finalId) return;
            try {
                const response = await getUserById(finalId);
                setUser(response.data.data);
            } catch (error) {
                console.error("Error fetching user:", error?.response?.data?.message || error.message);
            }
        };

        fetchUser();
    }, [finalId, routeId, currentUser]);

    useEffect(() => {
        const fetchReview = async () => {
            if (!finalId) return;
            try {
                const response = await getReviewByUserId(finalId, page);
                setReview(response.data.data);
            } catch (error) {
                console.error("Error fetching review:", error?.response?.data?.message || error.message);
                console.log(error);
            }
        };

        fetchReview();
    }, [finalId, page]);

    useEffect(() => {
        const fetchWatchlists = async () => {
            if (!finalId) return;
            try {
                const response = await getUserFilmByUserId(finalId);
                if (response.status === 200) {
                    setWatchlists(response.data.data);
                }
            } catch (error) {
                console.error("Error fetching watchlists:", error?.response?.data?.message || error.message);
            }
        };

        fetchWatchlists();
    }, [finalId]);

    return (
        <>
            <div className="mt-28 flex justify-center px-4">
                <div className="w-full max-w-4xl bg-gradient-to-br from-indigo-50 to-white rounded-3xl shadow-xl p-8 space-y-6">
                    <div className="flex items-center gap-6">
                        <div className="w-24 h-24 rounded-full bg-gradient-to-tr from-indigo-400 to-purple-400 flex items-center justify-center text-white text-3xl font-bold shadow-md">
                            {user?.nama?.charAt(0).toUpperCase() || "?"}
                        </div>
                        <div>
                            <h1 className="text-2xl font-bold text-gray-800">{user?.nama}</h1>
                            <p className="text-gray-500">@{user?.username}</p>
                        </div>
                    </div>

                    <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6">
                        <div className="bg-white p-4 rounded-xl shadow-sm border border-gray-100">
                            <h2 className="text-lg font-semibold mb-2 text-indigo-700">Informasi Pribadi</h2>
                            <ul className="text-gray-700 space-y-1">
                                <li><span className="font-medium">Nama:</span> {user?.nama}</li>
                                <li><span className="font-medium">Username:</span> {user?.username}</li>
                                {/* Tambahkan data lain di sini jika ada */}
                            </ul>
                        </div>

                        <div className="bg-white p-4 rounded-xl shadow-sm border border-gray-100">
                            <h2 className="text-lg font-semibold mb-2 text-indigo-700">Aktivitas</h2>
                            <p className="text-gray-600">
                                Total Ulasan: {review?.reviews?.length || 0}
                            </p>
                            <p className="text-gray-600">
                                Total Watchlist: {watchlists?.length || 0}
                            </p>
                        </div>
                    </div>

                    <div className="bg-white p-4 rounded-xl shadow-sm border border-gray-100">
                        <h2 className="text-lg font-semibold mb-2 text-indigo-700">Tentang Saya</h2>
                        <p className="text-gray-600">{user?.bio}</p>
                    </div>
                </div>
            </div>

            <div className="mt-12 px-4 max-w-4xl mx-auto space-y-4">
                <h2 className="text-3xl font-semibold ps-4 pt-4">Watchlist</h2>
                {
                    watchlists && watchlists.map((watch) => (
                        <WatchlistLayout key={watch.id} watchlist={watch}>
                            <Informasi watch={watch} />
                        </WatchlistLayout>
                    ))
                }
            </div>

            <div className="mt-12 px-4 max-w-4xl mx-auto">
                <Review review={review} setPage={setPage} page={page} />
            </div>
        </>
    );
};

export default ProfilePage;

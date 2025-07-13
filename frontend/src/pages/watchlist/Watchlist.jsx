import React, { useContext, useEffect, useState } from "react";
import { AuthContext } from "../../contexts/authContext";
import { getUserFilmByUserId } from "../../service/userFilm";
import Detail from "./components/Detail";
import Informasi from "./components/Informasi";
import WatchlistLayout from "../../layouts/WatchlistLayout";

const WatchListPage = () => {
    const { user, loading } = useContext(AuthContext);
    const [watchlist, setWatchlist] = useState();
    const [page, setPage] = useState(1);

    useEffect(() => {
        if (loading) return;
        const fetchWatchlist = async () => {
            try {
                const response = await getUserFilmByUserId(user.id, page);
                if (response.status === 200) {
                    setWatchlist(response.data.data);
                    console.log(response.data.data);
                }
            } catch (error) {
                console.error("Error fetching watchlist:", error);
            }
        };

        fetchWatchlist();
    }, [user, loading, page]);

    return (
        <div className="mx-auto max-w-4xl mt-28 px-4 space-y-6">
            <h1 className="text-3xl font-bold text-center mb-8">🎬 Your Watchlist</h1>
            {watchlist && watchlist.user_films?.map((user_film) => (
                <WatchlistLayout watchlist={user_film}>
                    <Informasi watch={user_film} />
                    <Detail watchlist={user_film} />
                </WatchlistLayout>
            ))}

            {/* Pagination */}
            <div className="flex justify-center mt-2 pb-2 space-x-2">
                {watchlist && Array.from({ length: watchlist.count_page }, (_, i) => (
                    <button
                        key={i}
                        className={`px-3 py-1 rounded bg-gray-100 hover:bg-gray-200 text-sm ${page === i + 1 ? "bg-primary text-white" : ""
                            }`}
                        onClick={() => setPage(i + 1)}
                    >
                        {i + 1}
                    </button>
                ))}
            </div>
        </div>
    );
};

export default WatchListPage;

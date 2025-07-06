import React, { useContext, useEffect, useState } from "react";
import { AuthContext } from "../../contexts/authContext";
import { getUserFilmByUserId } from "../../service/userFilm";
import Detail from "./components/Detail";
import Informasi from "./components/Informasi";
import WatchlistLayout from "../../layouts/WatchlistLayout";

const WatchListPage = () => {
    const { user, loading } = useContext(AuthContext);
    const [watchlist, setWatchlist] = useState([]);

    useEffect(() => {
        if (loading) return;
        const fetchWatchlist = async () => {
            try {
                const response = await getUserFilmByUserId(user.id);
                if (response.status === 200) {
                    setWatchlist(response.data.data);
                }
            } catch (error) {
                console.error("Error fetching watchlist:", error);
            }
        };

        fetchWatchlist();
    }, [user, loading]);

  return (
    <div className="mx-auto max-w-4xl mt-28 px-4 space-y-6">
        <h1 className="text-3xl font-bold text-center mb-8">🎬 Your Watchlist</h1>
        { watchlist && watchlist.map((watch) => (
            <WatchlistLayout key={watch.id} watchlist={watch}>
                <Informasi watch={watch} />
                <Detail watchlist={watch} />
            </WatchlistLayout>
        ))}
    </div>
  );
};

export default WatchListPage;

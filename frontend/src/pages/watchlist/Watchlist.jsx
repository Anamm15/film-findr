import { useContext } from "react";
import { AuthContext } from "../../contexts/authContext";
import Detail from "./components/Detail";
import Informasi from "./components/Informasi";
import WatchlistLayout from "../../layouts/WatchlistLayout";
import Pagination from "../../components/Pagination"
import { useFetchWatchlist } from "../../hooks/watchlist/useFetchWatchlist";
import Loading from "../../components/Loading";

const WatchListPage = () => {
    const { user } = useContext(AuthContext);
    const { watchlists, loading: watchlistLoading, page, setPage } = useFetchWatchlist(user?.id);
    return (
        <div className="mx-auto max-w-4xl mt-28 px-4 space-y-6">
            <h1 className="text-3xl font-bold text-center mb-8">🎬 Your Watchlist</h1>
            {
                watchlistLoading ? (
                    <Loading>Loading your watchlist...</Loading>
                ) : (
                    watchlists?.user_films && watchlists.user_films.map((user_film) => (
                        <WatchlistLayout key={user_film.id} watchlist={user_film}>
                            <Informasi watch={user_film} />
                            <Detail watchlist={user_film} />
                        </WatchlistLayout>
                    ))
                )
            }

            <Pagination contents={watchlists} page={page} setPage={setPage} />
        </div>
    );
};

export default WatchListPage;

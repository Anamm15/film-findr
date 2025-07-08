import WatchlistLayout from "../../../layouts/WatchlistLayout";
import Informasi from "./Informasi";

const Watchlist = (props) => {
    const { watchlists, watchlistPage, setWatchlistPage } = props;
    return (
        <div className="mt-12 px-4 max-w-4xl mx-auto space-y-4">
            <h2 className="text-3xl font-semibold ps-4 pt-4">Watchlist</h2>
            {
                watchlists && watchlists.user_films.map((user_film) => (
                    <WatchlistLayout watchlist={user_film}>
                        <Informasi watch={user_film} />
                    </WatchlistLayout>
                ))
            }

            {/* Pagination */}
            <div className="flex justify-center mt-2 pb-2 space-x-2">
                {watchlists && Array.from({ length: watchlists.count_page }, (_, i) => (
                    <button
                    key={i}
                    className={`px-3 py-1 rounded bg-gray-100 hover:bg-gray-200 text-sm ${
                        watchlistPage === i + 1 ? "bg-primary text-white" : ""
                    }`}
                    onClick={() => setWatchlistPage(i + 1)}
                    >
                    {i + 1}
                    </button>
                ))}
            </div>
        </div>
    )
}

export default Watchlist;
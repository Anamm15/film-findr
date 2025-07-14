import WatchlistLayout from "../../../layouts/WatchlistLayout";
import Informasi from "./Informasi";
import Pagination from "../../../components/Pagination";

const Watchlist = (props) => {
    const { watchlists, watchlistPage, setWatchlistPage } = props;
    return (
        <div className="mt-12 px-4 max-w-4xl mx-auto space-y-4">
            <h2 className="text-3xl font-semibold ps-4 pt-4">Watchlist</h2>
            {
                watchlists?.user_films && watchlists.user_films.map((user_film) => (
                    <WatchlistLayout watchlist={user_film}>
                        <Informasi watch={user_film} />
                    </WatchlistLayout>
                ))
            }

            <Pagination contents={watchlists} page={watchlistPage} setPage={setWatchlistPage} />
        </div>
    )
}

export default Watchlist;
import ListFilm from "./components/ListFilm";
import { useFetchTopFilms } from "../../hooks/film/useFetchTopFilms";
import { useFetchTrendingFilms } from "../../hooks/film/useFetchTrendingFilms";

const TopFilmPage = () => {
   const { topFilms, loading: topFilmLoading } = useFetchTopFilms();
   const { trendingFilms, loading: trendingFilmLoading } = useFetchTrendingFilms();

   if (topFilmLoading || trendingFilmLoading) {
      return <div className="flex justify-center items-center h-screen">Loading...</div>;
   }

   return (
      <>
         <div className="p-4 xl:max-w-[1280px] mx-auto mt-12">
            <h1 className="text-3xl md:text-4xl font-bold mb-4 text-text mt-10">Top Film</h1>
            {
               topFilms && <ListFilm films={topFilms} />
            }

            <h1 className="text-3xl md:text-4xl font-bold mb-4 text-text mt-10">Trending Film</h1>
            {
               trendingFilms && <ListFilm films={trendingFilms} />
            }
         </div>
      </>
   )
}

export default TopFilmPage;
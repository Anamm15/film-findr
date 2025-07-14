import { useEffect, useState } from "react";
import { getTopFilm, getTrendingFilm } from "../../service/film";
import ListFilm from "./components/ListFilm";

const TopFilmPage = () => {
   const [topFilms, setTopFilms] = useState(null);
   const [trendingFilms, setTrendingFilms] = useState(null);

   useEffect(() => {
      const fetchTopFilms = async () => {
         try {
            const response = await getTopFilm();
            setTopFilms(response.data.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         }
      }

      fetchTopFilms();
   }, []);

   useEffect(() => {
      const fetchTrendingFilms = async () => {
         try {
            const response = await getTrendingFilm();
            setTrendingFilms(response.data.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         }
      }

      fetchTrendingFilms();
   }, []);

   return (
      <>
         <div className="p-4 xl:max-w-[1280px] mx-auto mt-28">
            <h1 className="text-4xl font-bold mb-4 text-text mt-10">Top Film</h1>
            {
               topFilms && <ListFilm films={topFilms} />
            }

            <h1 className="text-4xl font-bold mb-4 text-text mt-10">Trending Film</h1>
            {
               trendingFilms && <ListFilm films={trendingFilms} />
            }
         </div>
      </>
   )
}

export default TopFilmPage;
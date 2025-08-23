import { useEffect, useState } from "react";
import { getTrendingFilm } from "../../service/film";

export function useFetchTrendingFilms() {
   const [trendingFilms, setTrendingFilms] = useState(null);
   const [loading, setLoading] = useState(true);

   useEffect(() => {
      const fetchTrendingFilms = async () => {
         setLoading(true);
         try {
            const response = await getTrendingFilm();
            setTrendingFilms(response.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         } finally {
            setLoading(false);
         }
      }

      fetchTrendingFilms();
   }, []);

   return { trendingFilms, loading };
}
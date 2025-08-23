import { useState, useEffect } from "react";
import { getTopFilm } from "../../service/film";

export function useFetchTopFilms() {
   const [topFilms, setTopFilms] = useState(null);
   const [loading, setLoading] = useState(true);

   useEffect(() => {
      const fetchTopFilms = async () => {
         setLoading(true);
         try {
            const response = await getTopFilm();
            setTopFilms(response.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         } finally {
            setLoading(false);
         }
      }

      fetchTopFilms();
   }, []);

   return { topFilms, loading };
}
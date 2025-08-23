import { useEffect, useState } from "react";
import { getFilmById } from "../../service/film";

export function useFetchFilm(id) {
   const [film, setFilm] = useState(null);
   const [loading, setLoading] = useState(false);

   useEffect(() => {
      const fetchFilm = async () => {
         setLoading(true);
         try {
            const response = await getFilmById(id);
            setFilm(response.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         } finally {
            setLoading(false);
         }
      };

      fetchFilm();
   }, [id]);

   return { film, loading };
}

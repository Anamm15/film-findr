import { useEffect, useState } from "react";
import { getAllFilm } from "../../service/film";
import { INIT_PAGE_NUMBER } from "../../utils/constant";

export function useFetchFilms() {
   const [films, setFilms] = useState(null);
   const [page, setPage] = useState(INIT_PAGE_NUMBER);
   const [loading, setLoading] = useState(false);

   useEffect(() => {
      const fetchAllFilms = async () => {
         setLoading(true);
         try {
            const response = await getAllFilm(page);
            setFilms(response.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         } finally {
            setLoading(false);
         }
      };

      fetchAllFilms();
   }, [page]);

   return { films, setFilms, page, setPage, loading };
}

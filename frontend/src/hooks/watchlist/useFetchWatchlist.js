import { useState, useEffect } from "react";
import { getUserFilmByUserId } from "../../service/userFilm";
import { INIT_PAGE_NUMBER } from "../../utils/constant";

export function useFetchWatchlist(userId) {
   const [watchlists, setWatchlists] = useState(null);
   const [loading, setLoading] = useState(false);
   const [page, setPage] = useState(INIT_PAGE_NUMBER);

   useEffect(() => {
      const fetchWatchlists = async () => {
         if (!userId) return;
         setLoading(true);
         try {
            const response = await getUserFilmByUserId(userId, page);
            setWatchlists(response.data);
         } catch (error) {
            console.error("Error fetching watchlists:", error);
         } finally {
            setLoading(false);
         }
      };

      fetchWatchlists();
   }, [userId, page]);

   return { watchlists, page, setPage, loading };
}
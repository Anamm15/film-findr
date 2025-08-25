import { useEffect, useState } from "react";
import { getReviewByFilmId } from "../../service/review";
import { INIT_PAGE_NUMBER } from "../../utils/constant";

export function useFetchReviewByFilmId(id) {
   const [reviews, setReviews] = useState(null);
   const [loading, setLoading] = useState(false);
   const [page, setPage] = useState(INIT_PAGE_NUMBER);

   useEffect(() => {
      const fetchReviews = async () => {
         setLoading(true);
         try {
            const response = await getReviewByFilmId(id, page);
            setReviews(response.data);
         } catch (error) {
            console.error("Error fetching review:", error);
         } finally {
            setLoading(false);
         }
      };

      fetchReviews();
   }, [id, page]);

   return { reviews, setReviews, page, setPage, loading };
}
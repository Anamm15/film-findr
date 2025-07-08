import { useState, useEffect, useContext } from "react";
import { useParams } from "react-router-dom";
import { getReviewByUserId } from "../../service/review";
import { getUserById } from "../../service/user";
import { AuthContext } from "../../contexts/authContext";
import { getUserFilmByUserId } from "../../service/userFilm";
import ReviewLayout from "../../layouts/ReviewLayout";
import Watchlist from "./components/Watchlist";
import Akun from "./components/Akun";

const ProfilePage = () => {
   const params = useParams();
   const routeId = params.id;
   const { user: currentUser } = useContext(AuthContext);

   const [user, setUser] = useState(null);
   const [review, setReview] = useState([]);
   const [watchlists, setWatchlists] = useState(null);
   const [watchlistPage, setWatchlistPage] = useState(1);
   const [reviewPage, setReviewPage] = useState(1);

   const finalId = routeId || currentUser?.id;

   useEffect(() => {
   const fetchUser = async () => {
      if (!finalId) return;
      try {
         const response = await getUserById(finalId);
         setUser(response.data.data);
      } catch (error) {
         console.error("Error fetching user:", error);
      }
   };

   fetchUser();
   }, [finalId, routeId, currentUser]);

   useEffect(() => {
   const fetchReview = async () => {
      if (!finalId) return;
      try {
         const response = await getReviewByUserId(finalId, reviewPage);
         setReview(response.data.data);
      } catch (error) {
         console.error("Error fetching review:", error);
      }
   };

   fetchReview();
   }, [finalId, reviewPage]);

   useEffect(() => {
      const fetchWatchlists = async () => {
         if (!finalId) return;
         try {
            const response = await getUserFilmByUserId(finalId, watchlistPage);
            if (response.status === 200) {
               setWatchlists(response.data.data);
            }
         } catch (error) {
            console.error("Error fetching watchlists:", error);
         }
      };

      fetchWatchlists();
   }, [finalId, watchlistPage]);

   return (
      <>
         <Akun user={user} review={review} watchlists={watchlists} />
         <Watchlist watchlists={watchlists} watchlistPage={watchlistPage} setWatchlistPage={setWatchlistPage} />
         <div className="mt-12 px-4 max-w-4xl mx-auto">
            <ReviewLayout review={review} setPage={setReviewPage} page={reviewPage} />
         </div>
      </>
   );
};

export default ProfilePage;

import { useState, useEffect, useContext } from "react";
import { useParams } from "react-router-dom";
import { getReviewByUserId } from "../../service/review";
import { getUserByUsername } from "../../service/user";
import { AuthContext } from "../../contexts/authContext";
import { getUserFilmByUserId } from "../../service/userFilm";
import ReviewLayout from "../../layouts/ReviewLayout";
import Watchlist from "./components/Watchlist";
import Akun from "./components/Akun";

const ProfilePage = () => {
   const params = useParams();
   const usernameParams = params.username;
   const { user: currentUser } = useContext(AuthContext);

   const [user, setUser] = useState(null);
   const [review, setReview] = useState([]);
   const [watchlists, setWatchlists] = useState(null);
   const [watchlistPage, setWatchlistPage] = useState(1);
   const [reviewPage, setReviewPage] = useState(1);

   const finalUsername = usernameParams || currentUser?.username;

   useEffect(() => {
      const fetchUser = async () => {
         if (!finalUsername) return;
         try {
            const response = await getUserByUsername(finalUsername);
            setUser(response.data.data);
         } catch (error) {
            console.error("Error fetching user:", error);
         }
      };

      fetchUser();
   }, [finalUsername, usernameParams, currentUser]);

   useEffect(() => {
      const fetchReview = async () => {
         if (!finalUsername && !user) return;
         try {
            const response = await getReviewByUserId(user.id, reviewPage);
            setReview(response.data.data);
         } catch (error) {
            console.error("Error fetching review:", error);
         }
      };

      fetchReview();
   }, [finalUsername, reviewPage, user]);

   useEffect(() => {
      const fetchWatchlists = async () => {
         if (!finalUsername && !user) return;
         try {
            const response = await getUserFilmByUserId(user.id, watchlistPage);
            if (response.status === 200) {
               setWatchlists(response.data.data);
               console.log(response.data.data);

            }
         } catch (error) {
            console.error("Error fetching watchlists:", error);
         }
      };

      fetchWatchlists();
   }, [finalUsername, watchlistPage, user]);

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

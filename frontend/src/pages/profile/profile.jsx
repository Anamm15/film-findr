import { useContext } from "react";
import { useParams } from "react-router-dom";
import { AuthContext } from "../../contexts/authContext";
import ReviewLayout from "../../layouts/ReviewLayout";
import Watchlist from "./components/Watchlist";
import Akun from "./components/Akun";
import { useFetchUserByUsername } from "../../hooks/user/useFetchUserByUsername";
import { useFetchReviewByUserId } from "../../hooks/review/useFetchReviewByUserId";
import { useFetchWatchlist } from "../../hooks/watchlist/useFetchWatchlist";
import PageLoading from "../../components/PageLoading";

const ProfilePage = () => {
   const params = useParams();
   const usernameParams = params.username;
   const { user: currentUser } = useContext(AuthContext);
   const finalUsername = usernameParams || currentUser?.username;

   const { user, loading: userLoading } = useFetchUserByUsername(finalUsername);
   const { reviews, loading: reviewLoading, page: reviewPage, setPage: setReviewPage } = useFetchReviewByUserId(user?.id);
   const { watchlists, loading: watchlistLoading, page: watchlistPage, setPage: setWatchlistPage } = useFetchWatchlist(user?.id);

   if (userLoading) {
      return <PageLoading message="Loading user" />;
   }

   return (
      <>
         <Akun user={user} review={reviews} watchlists={watchlists} />
         <Watchlist watchlists={watchlists} watchlistPage={watchlistPage} setWatchlistPage={setWatchlistPage} loading={watchlistLoading} />
         <div className="mt-12 px-4 max-w-4xl mx-auto">
            <ReviewLayout review={reviews} setPage={setReviewPage} page={reviewPage} loading={reviewLoading} />
         </div>
      </>
   );
};

export default ProfilePage;
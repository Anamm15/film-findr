import { useParams } from "react-router-dom";
import WatchListForm from "./components/AddWatchlist";
import RekomendasiFilm from "./components/Rekomendasi";
import AddReview from "./components/AddReview";
import Sinopsis from "./components/Sinopsis";
import ReviewLayout from "../../layouts/ReviewLayout";
import InformasiFilm from "./components/InformasiFilm";
import Gambar from "./components/Gambar";
import { useFetchFilm } from "../../hooks/film/useFetchFilm";
import { useFetchFilms } from "../../hooks/film/UseFetchFilms";
import { useFetchReviewByFilmId } from "../../hooks/review/useFetchReviewByFilmId";
import PageLoading from "../../components/PageLoading";

const DetailFilmPage = () => {
   const { id } = useParams();
   const { film, page, setPage, loading: loadingFilms } = useFetchFilm(id);
   const { reviews, setReviews, loading: loadingReviews } = useFetchReviewByFilmId(id, page);
   const { films } = useFetchFilms();

   if (loadingFilms) {
      return <PageLoading message="Loading film..." />;
   }

   return (
      <div className="max-w-5xl mx-auto px-4 pb-10 bg-background">
         {
            film && (
               <>
                  <div className="flex flex-col md:flex-row gap-6 my-8 shadow-md rounded-xl p-4 relative mt-28">
                     <Gambar film={film} />
                     <div>
                        <InformasiFilm film={film} />
                        <WatchListForm id={id} />
                     </div>
                  </div>

                  <Sinopsis sinopsis={film.sinopsis} />
                  <ReviewLayout review={reviews} setReviews={setReviews} setPage={setPage} page={page} loading={loadingReviews} />
                  <AddReview id={id} setReviews={setReviews} />
                  {
                     films && <RekomendasiFilm films={films.films} />
                  }
               </>
            )
         }
      </div>
   );
};

export default DetailFilmPage;
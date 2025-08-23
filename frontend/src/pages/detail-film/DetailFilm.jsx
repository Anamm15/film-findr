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

const DetailFilmPage = () => {
   const { id } = useParams();
   const { film, page, setPage, loading: loadingFilms } = useFetchFilm(id);
   const { reviews, loading: loadingReviews } = useFetchReviewByFilmId(id, page);
   const { films } = useFetchFilms();

   if (loadingFilms || loadingReviews) {
      return <div>Loading...</div>;
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
                  <ReviewLayout review={reviews} setPage={setPage} page={page} />
                  <AddReview id={id} />
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
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { getFilmById, getAllFilm } from "../../service/film";
import { getReviewByFilmId } from "../../service/review";
import WatchListForm from "./components/AddWatchlist";
import RekomendasiFilm from "./components/Rekomendasi";
import AddReview from "./components/AddReview";
import Sinopsis from "./components/Sinopsis";
import Review from "./components/Review";
import InformasiFilm from "./components/InformasiFilm";
import Gambar from "./components/Gambar";

const DetailFilmPage = () => {
    const id = useParams().id;
    const [film, setFilm] = useState(null);
    const [page, setPage] = useState(1);
    const [review, setReview] = useState(null);
    
    const [films, setFilms] = useState([]);
    const [isFilmsFetched, setIsFilmsFetched] = useState(false);

    useEffect(() => {
        const fetchFilm = async () => {
            try {
                const response = await getFilmById(id);
                setFilm(response.data.data);
            } catch (error) {
                console.error("Error fetching film:", error.data.message);
            }
        };

        fetchFilm();
    }, [id]);

    useEffect(() => {
        const fetchReview = async () => {
            try {
                const response = await getReviewByFilmId(id, page);
                setReview(response.data.data);
            } catch (error) {
                console.error("Error fetching review:", error.data.message);
            }
        };

        fetchReview();
    }, [id, page]);

    useEffect(() => {
        const fetchAllFilms = async () => {
            try {
                const response = await getAllFilm();
                setFilms(response.data.data);
            } catch (error) {
                console.error("Error fetching films:", error.data.message);
            }
        }

        fetchAllFilms();
        if (!isFilmsFetched) {
            setIsFilmsFetched(false);
        }
    }, [isFilmsFetched]);

  return (
    <div className="max-w-5xl mx-auto px-4 pb-10">
      {
        film && (
          <>
            <div className="flex gap-6 my-8 shadow-md rounded-xl p-4 relative mt-28">
              <Gambar film={film} />
              <div>
                  <h1 className="text-3xl font-bold mb-4">{film.judul}</h1>
                  <InformasiFilm film={film} />
                  <WatchListForm id={id} />
              </div>
            </div>

            <Sinopsis sinopsis={film.sinopsis} />
            <Review review={review} setPage={setPage} page={page} />
            <AddReview id={id} />
            <RekomendasiFilm films={films} />
          </>
        )
      }
    </div>
  );
};

export default DetailFilmPage;
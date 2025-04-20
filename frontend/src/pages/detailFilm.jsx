import React from "react";
import { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import { getFilmById } from "../api/film";
import { createReview, getReviewByFilmId } from "../api/review";

const DetailFilmPage = () => {
    const id = useParams().id;
    const [film, setFilm] = useState(null);
    const [review, setReview] = useState(null);
    const [page, setPage] = useState(1);
    const [newReview, setNewReview] = useState("");
    const [message, setMessage] = useState("");
    const [rating, setRating] = useState(0);
    const [watchListStatus, setWatchListStatus] = useState(false);

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

    const handleReviewSubmit = async (e) => {
        e.preventDefault();

        try {
            const data = {
              film_id: Number(id), 
              komentar: newReview,
              rating: Number(rating)
            };
            
            const response = await createReview(data, localStorage.getItem("token"));
            setMessage(response.data.message);
            setNewReview("");
          } catch (error) {
          console.log(error);
            setMessage(error.response.data.message);
        }
    };

  return (
    <div className="max-w-5xl mx-auto px-4 py-6">
      {
        film && (
          <>
            {/* Judul */}
            <h1 className="text-3xl font-bold mb-4">{film.judul}</h1>

            {/* Gambar Carousel */}
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mb-6">
              {film.film_gambar.map((img) => (
                <img
                  key={img.id}
                  src={img.url}
                  alt={`Gambar ${img.id + 1}`}
                  className="rounded-lg shadow-md object-cover h-48 w-full"
                />
              ))}
            </div>

            {/* Tambah ke Watchlist */}
            <div className="bg-white">
              <button 
                  className="font-semibold text-white bg-blue-500 px-4 py-2 rounded-full mb-2"                
                >Tambah ke Watchlist</button>
            </div>

            {/* Informasi Film */}
            <div className="bg-white rounded-xl shadow p-6 mb-6 space-y-3">
              <div className="flex flex-wrap gap-2 mb-2">
                {film.genres.slice(0, 3).map((genre) => (
                  <span
                    key={genre.id}
                    className="bg-blue-500 text-white px-3 py-1 rounded-full text-sm"
                  >
                    {genre.nama}
                  </span>
                ))}
              </div>
              <p><strong>Tanggal Rilis:</strong> {film.tanggal_rilis}</p>
              <p><strong>Durasi:</strong> {film.durasi} menit</p>
              <p><strong>Status:</strong> <span className="capitalize">{film.status}</span></p>
              <p><strong>Rating:</strong> ⭐ {film.rating}/10</p>
              <p><strong>Sutradara:</strong> {film.sutradara}</p>
              <p><strong>Total Episode:</strong> {film.total_episode}</p>
            </div>

            {/* Sinopsis */}
            <div className="bg-gray-100 rounded-xl p-5 mb-6">
              <h2 className="text-xl font-semibold mb-2">Sinopsis</h2>
              <p className="text-gray-700">{film.sinopsis}</p>
            </div>

            {/* Review */}
            <div className="bg-white rounded-xl shadow p-5">
              <h2 className="text-xl font-semibold mb-4">Review</h2>
              <div className="space-y-4">
                {review && review.reviews && review.reviews.map((r, idx) => (
                  <div key={idx} className="border-b pb-3">
                    <p className="font-semibold">{r.user.username}</p>
                    <p className="text-gray-600">{r.komentar}</p>

                    <div className="flex items-center space-x-4 mt-2">
                      {/* Like Icon */}
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className={`w-5 h-5 cursor-pointer ${
                          r.user_reaksi?.reaksi === "like" ? "text-blue-500" : "text-gray-400"
                        }`}
                        viewBox="0 0 20 20"
                        fill="currentColor"
                      >
                        <path d="M2 10c0-.6.4-1 1-1h3V5.5A2.5 2.5 0 018.5 3h1.086a1 1 0 01.707.293l3.414 3.414a1 1 0 01.293.707v6.172a1 1 0 01-.293.707l-3.414 3.414A1 1 0 019.586 18H8.5A2.5 2.5 0 016 15.5V13H3a1 1 0 01-1-1v-2z" />
                      </svg>

                      {/* Dislike Icon */}
                      <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className={`w-5 h-5 cursor-pointer ${
                          r.user_reaksi?.reaksi === "dislike" ? "text-red-500" : "text-gray-400"
                        }`}
                        viewBox="0 0 20 20"
                        fill="currentColor"
                      >
                        <path d="M18 10c0 .6-.4 1-1 1h-3v3.5a2.5 2.5 0 01-2.5 2.5h-1.086a1 1 0 01-.707-.293l-3.414-3.414a1 1 0 01-.293-.707V6.414a1 1 0 01.293-.707L9.707 2.293A1 1 0 0110.414 2H11.5A2.5 2.5 0 0114 4.5V7h3a1 1 0 011 1v2z" />
                      </svg>
                    </div>
                  </div>
                ))}
              </div>

              {/* Pagination */}
              <div className="flex justify-center mt-6 space-x-2">
                {review && Array.from({ length: review.count_page }, (_, i) => (
                  <button
                    key={i}
                    className="px-3 py-1 rounded bg-gray-100 hover:bg-gray-200 text-sm"
                    onClick={() => setPage(i + 1)}
                  >
                    {i + 1}
                  </button>
                ))}
              </div>
            </div>

            {/* Add Review */}
            <div className="bg-white rounded-xl shadow p-5">
              <h2 className="text-xl font-semibold mb-4">Tambah Review</h2>
              <form onSubmit={handleReviewSubmit} className="space-y-4">
                <input
                  type="text"
                  placeholder="Your Comment"
                  value={newReview}
                  onChange={(e) => setNewReview(e.target.value)}
                  className="w-full p-2 border rounded"
                />
                <label className="mt-2">Rating</label>
                <input type="number" placeholder="Rating" value={rating} onChange={(e) => setRating(e.target.value)} className="w-full p-2 border rounded" min="1" max="10" />
                <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded">Submit</button>
                {message && <p>{message}</p>}
              </form>
            </div>
          </>
        )
      }
    </div>
  );
};

export default DetailFilmPage;
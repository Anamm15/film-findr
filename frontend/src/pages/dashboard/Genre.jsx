import { useState, useEffect } from "react";
import { deleteGenre, getAllGenre } from "../../service/genre";
// 1. Impor komponen bar chart yang baru
import GenreBarChart from "./components/GenreBarChart";
import AddGenre from "./components/AddGenre";

const GenreDashboardPage = () => {
   // ... (semua state dan fungsi useEffect tetap sama)
   const [genres, setGenres] = useState([]);
   const [message, setMessage] = useState("");
   const [colorMessage, setColorMessage] = useState("");
   const [isGenresLoaded, setIsGenresLoaded] = useState(false);

   useEffect(() => {
      const fetchGenres = async () => {
         try {
            const response = await getAllGenre();
            const genresWithData = response.data.data.map(genre => ({
               ...genre,
               jumlahFilm: Math.floor(Math.random() * 200) + 50,
            }));

            // Urutkan genre dari yang paling sedikit filmnya ke paling banyak
            // agar di bar chart horizontal, yang terpanjang ada di atas
            genresWithData.sort((a, b) => a.jumlahFilm - b.jumlahFilm);

            setGenres(genresWithData);
         } catch (error) {
            console.error("Error fetching genres:", error);
         }
      };

      fetchGenres();
      if (!isGenresLoaded) {
         setIsGenresLoaded(true);
      }
   }, [isGenresLoaded]);

   // ... (fungsi handleDeleteGenre tetap sama)
   const handleDeleteGenre = async (genreId) => {
      try {
         const response = await deleteGenre(genreId);
         if (response.status === 200) {
            setMessage(response.data.message);
            setColorMessage("text-green-600");
            setGenres(genres.filter((genre) => genre.id !== genreId));
         }
      } catch (error) {
         setMessage(error.response.data.error);
         setColorMessage("text-red-600");
      }
   };


   return (
      <div className="grid p-4 md:p-6">
         {/* Kolom kiri untuk manajemen genre tidak berubah */}
         <div className="lg:col-span-2">
            <h1 className="text-3xl font-bold mb-8">Manajemen Genre</h1>
            <div>
               <label className="block text-text text-2xl font-semibold mb-2">Daftar Genre Aktif</label>
               <div className="flex flex-wrap gap-2">
                  {genres.map((genre) => (
                     <div
                        key={genre.id}
                        className="flex items-center px-3 py-1 text-sm rounded-full border text-slate-700 bg-slate-100 hover:bg-tertiary hover:text-white hover:border-tertiary transition-all duration-200 group"
                     >
                        <span>{genre.nama}</span>
                        <button
                           onClick={() => handleDeleteGenre(genre.id)}
                           type="button"
                           className="ml-2 text-slate-400 group-hover:text-white font-bold"
                        >
                           ×
                        </button>
                     </div>
                  ))}
               </div>
            </div>
            {message && (
               <div className={`mt-4 ${colorMessage}`}>
                  <p>{message}</p>
               </div>
            )}
         </div>

         {/* Form untuk menambah genre */}
         <AddGenre />

         {/* Kolom kanan untuk grafik */}
         <div className="lg:col-span-3 bg-white p-6 rounded-2xl shadow-lg mt-10">
            <h2 className="text-2xl font-bold text-slate-800 mb-4">Popularitas Genre</h2>
            <div className="relative h-[400px]">
               {genres.length > 0 ? (
                  // 2. Ganti komponen di sini
                  <GenreBarChart genres={genres} />
               ) : (
                  <p className="text-center text-slate-500">Memuat data grafik...</p>
               )}
            </div>
         </div>
      </div>
   );
};

export default GenreDashboardPage;
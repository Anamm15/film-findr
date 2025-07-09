import { useState, useEffect } from "react";
import { deleteGenre, getAllGenre } from "../../service/genre";

const GenreDashboardPage = () => {
   const [genres, setGenres] = useState([]);
   const [message, setMessage] = useState("");
   const [colorMessage, setColorMessage] = useState("");
   const [isGenresLoaded, setIsGenresLoaded] = useState(false);

   useEffect(() => {
         const fetchGenres = async () => {
            try {
               const response = await getAllGenre();
               setGenres(response.data.data);
            } catch (error) {
               console.error("Error fetching genres:", error.data.message);
            }
         };
   
         fetchGenres();
         if (!isGenresLoaded) {
            setIsGenresLoaded(true);
         }
      }, [isGenresLoaded]);

   const handleDeleteGenre = async(genreId) => {
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
   }

   return (
      <div>
         <h1 className="text-3xl font-bold mb-8">Genre</h1>
         <div className="mt-4">
            <label className="block text-gray-600 text-sm mb-2">Pilih Genre</label>

            <div className="flex flex-wrap gap-4 mb-4">
               {genres.map((genre) => (
                     <div
                     key={genre.id}
                     className="flex items-center px-3 py-0.5 text-lg rounded-full border text-text hover:bg-primary hover:text-white border-primary"
                  >
                     {genre.nama}
                     <button
                        onClick={() => handleDeleteGenre(genre.id)}
                        type="button"
                        className="ml-2 hover:scale-105"
                     >
                        ×
                     </button>
                  </div>
               ))}
            </div>
         </div>
         {
            message && (
               <div className={`mt-4 ${colorMessage}`}>
                  <p>{message}</p>
               </div>
            )
         }
      </div>
   )
}

export default GenreDashboardPage;
import { useEffect, useState } from "react";
import { getAllFilm } from "../../service/film"
import FilmCard from "../../components/FilmCard";
import { useNavigate } from "react-router-dom";
import Button from "../../components/Button";

const FilmDashboardPage = () => {
   const [films, setFilms] = useState(null);
   const [page, setPage] = useState(1);
   const navigate = useNavigate();

   useEffect(() => {
      const fetchAllFilms = async () => {
         try {
            const response = await getAllFilm();
            setFilms(response.data.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         }
      }

      fetchAllFilms();
   }, []);

   const handleClickCard = (id) => {
      navigate(`/film/${id}`);
   };

   const handleButtonAddFilm = () => {
      navigate("/dashboard/films/new");
   }

   return (
      <div className="relative">
         <Button 
            type="button"
            className="fixed bottom-8 right-8 px-12 py-2 rounded-xl"
            onClick={handleButtonAddFilm}
         >
            Add Film
         </Button>
         
         <h1 className="text-4xl font-bold mb-4">List Film</h1>
         <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {films && films.films.map((film) => (
                <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)}  />
            ))}
        </div>

         {/* Pagination */}
         <div className="flex justify-center mt-4 pb-2 space-x-2">
         {films && Array.from({ length: films.count_page }, (_, i) => (
               <button
               key={i}
               className={`px-3 py-1 rounded bg-gray-100 hover:bg-gradient-secondary text-sm ${
                  page === i + 1 ? "bg-gradient-primary text-white" : ""
               }`}
               onClick={() => setPage(i + 1)}
               >
               {i + 1}
               </button>
         ))}
         </div> 
      </div>
   )
}

export default FilmDashboardPage;
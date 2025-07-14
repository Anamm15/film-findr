import { useEffect, useState } from "react";
import { getAllFilm } from "../../service/film"
import FilmCard from "../../components/FilmCard";
import { useNavigate } from "react-router-dom";
import Button from "../../components/Button";
import Pagination from "../../components/Pagination";

const FilmDashboardPage = () => {
   const [films, setFilms] = useState(null);
   const [page, setPage] = useState(1);
   const navigate = useNavigate();

   useEffect(() => {
      const fetchAllFilms = async () => {
         try {
            const response = await getAllFilm(page);
            setFilms(response.data.data);
         } catch (error) {
            console.error("Error fetching films:", error);
         }
      }

      fetchAllFilms();
   }, [page]);

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
            className="fixed bottom-4 right-8 px-16 py-2 rounded-xl"
            onClick={handleButtonAddFilm}
         >
            Add Film
         </Button>

         <h1 className="text-4xl font-bold mb-4">List Film</h1>
         <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {films && films.films?.map((film) => (
               <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)} />
            ))}
         </div>

         {/* Pagination */}
         <Pagination contents={films} page={page} setPage={setPage} />
      </div>
   )
}

export default FilmDashboardPage;
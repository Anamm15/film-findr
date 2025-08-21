import { useEffect, useState } from "react";
import { getAllFilm } from "../../service/film"
import { useNavigate } from "react-router-dom";
import Button from "../../components/Button";
import Pagination from "../../components/Pagination";
import AdminFilmCard from "./film/AdminFilmCard";
import { deleteFilm } from "../../service/film";

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

   const handleUpdateFilm = (id) => {
      navigate("/dashboard/films/update/" + id);
   }

   const handleDeleteFilm = async (id) => {
      try {
         await deleteFilm(id);
      } catch (error) {
         console.error("Error deleting film:", error);
      }
   }

   return (
      <div className="relative">
         <Button
            type="button"
            className="fixed bottom-4 right-8 px-12 py-2 rounded-xl"
            onClick={handleButtonAddFilm}
         >
            Add Film
         </Button>

         <h1 className="text-4xl font-bold mb-4">List Film</h1>
         <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
            {films && films.films?.map((film) => (
               <AdminFilmCard
                  key={film.id}
                  movie={film}
                  onClick={() => handleClickCard(film.id)}
                  onUpdate={() => handleUpdateFilm(film.id)}
                  onDelete={() => handleDeleteFilm(film.id)}
               />
            ))}
         </div>

         <Pagination contents={films} page={page} setPage={setPage} />
      </div>
   )
}

export default FilmDashboardPage;
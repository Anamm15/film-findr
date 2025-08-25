import { useNavigate } from "react-router-dom";
import FilmCard from "../../../components/FilmCard";
import FilmCardSkeleton from "../../../components/FilmCardSkeleton";

const ListFilm = (props) => {
   const { films, loading } = props;
   const navigate = useNavigate();

   const handleClickCard = (id) => {
      navigate(`/film/${id}`);
   };

   return (
      <>
         <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
            {
               loading ? (
                  <FilmCardSkeleton count={5} />
               ) : (
                  films && films.map((film) => (
                     <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)} />
                  ))
               )
            }
         </div>
      </>
   )
}


export default ListFilm;
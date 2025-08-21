import Button from "../../../components/Button";
import FilmCard from "../../../components/FilmCard";

const AdminFilmCard = ({ movie, onClick, onUpdate, onDelete }) => {
   return (
      <>
         <FilmCard movie={movie} onClick={onClick}>
            <div className="mt-2 flex items-center gap-2 w-full">
               <Button
                  className="w-full py-0.5 rounded"
                  type="button"
                  onClick={(e) => {
                     e.stopPropagation();
                     onUpdate();
                  }}
               >
                  Update
               </Button>
               <Button
                  className="w-full py-0.5 rounded"
                  type="button"
                  onClick={(e) => {
                     e.stopPropagation();
                     onDelete();
                  }}
               >
                  Delete
               </Button>
            </div>
         </FilmCard>
      </>
   )
}

export default AdminFilmCard;
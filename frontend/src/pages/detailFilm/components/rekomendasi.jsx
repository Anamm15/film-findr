import { useNavigate } from "react-router-dom";
import FilmCard from "../../../components/filmCard";

const RekomendasiFilm = (props) => {
    const {films} = props;
    const navigate = useNavigate();

    const handleClickCard = (id) => {
        navigate(`/film/${id}`);
    };

    return (
        <>
            <h1 className="text-4xl font-bold mb-4 mt-8">Rekomendasi Film</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
              {films && films.slice(0, 4).map((film) => (
                <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)}  />
              ))}
            </div>
        </>
    )
}

export default RekomendasiFilm;
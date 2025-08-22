import { useNavigate } from "react-router-dom";
import FilmCard from "../../../components/FilmCard";
import Pagination from "../../../components/Pagination";

const ListFilm = (props) => {
    const { films, page, setPage } = props;
    const navigate = useNavigate();

    const handleClickCard = (id) => {
        navigate(`/film/${id}`);
    };

    return (
        <>
            <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
                {films && films.films?.map((film) => (
                    <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)} />
                ))}
            </div>

            <Pagination contents={films} page={page} setPage={setPage} />
        </>
    )
}


export default ListFilm;
import { useEffect, useState } from "react";
import { getAllFilm } from "../api/film";
import FilmCard from "../components/filmCard";
import { useNavigate } from "react-router-dom";


const LandingPage = () => {
    const [films, setFilms] = useState([]);
    const [isFilmsFetched, setIsFilmsFetched] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        const fetchAllFilms = async () => {
            try {
                const response = await getAllFilm();
                setFilms(response.data.data);
            } catch (error) {
                console.error("Error fetching films:", error.data.message);
            }
        }

        fetchAllFilms();
        if (!isFilmsFetched) {
            setIsFilmsFetched(false);
        }
    }, [isFilmsFetched]);

    const handleClickCard = (id) => {
        navigate(`/film/${id}`);
    }

    return (
        <>
          <div className="p-4 container mx-auto">
            <h1 className="text-2xl font-bold mb-4">Daftar Film</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-6">
              {films && films.map((film) => (
                <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)}  />
              ))}
            </div>
          </div>
        </>
      );      
}

export default LandingPage;
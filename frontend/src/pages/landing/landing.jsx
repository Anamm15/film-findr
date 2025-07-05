import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../../components/navbar";
import FilmCard from "../../components/filmCard";
import { getAllFilm } from "../../service/film"


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
          <Navbar />
          <div className="p-4 xl:max-w-[1280px] mx-auto mt-28">
            <h1 className="text-4xl font-bold mb-4">Daftar Film</h1>
            <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 gap-6">
              {films && films.map((film) => (
                <FilmCard key={film.id} movie={film} onClick={() => handleClickCard(film.id)}  />
              ))}
            </div>
          </div>
        </>
      );      
}

export default LandingPage;
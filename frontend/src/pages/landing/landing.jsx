import { useEffect, useState } from "react";
import { getAllFilm } from "../../service/film";
import ListFilm from "./components/ListFilm";


const LandingPage = () => {
    const [films, setFilms] = useState(null);
    const [isFilmsFetched, setIsFilmsFetched] = useState(false);

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
        if (!isFilmsFetched) {
            setIsFilmsFetched(false);
        }
    }, [isFilmsFetched]);

    return (
        <>
            <div className="p-4 xl:max-w-[1280px] mx-auto mt-28">
                <h1 className="text-4xl font-bold mb-4 text-text">Daftar Film</h1>
                {
                    films && <ListFilm films={films.films} />
                }
            </div>
        </>
      );      
}

export default LandingPage;
import { useState, useEffect } from "react";
import { getFilmById, updateFilm } from "../../../service/film";
import Input from "../../../components/Input";
import Button from "../../../components/Button";
import TextArea from "../../../components/Textarea";
import { useParams } from "react-router-dom";
import { deleteGenre } from "../../../service/genre";

const UpdateFilmPage = () => {
   const { id } = useParams();
   const [judul, setJudul] = useState("");
   const [status, setStatus] = useState("");
   const [sinopsis, setSinopsis] = useState("");
   const [sutradara, setSutradara] = useState("");
   const [tanggalRilis, setTanggalRilis] = useState("");
   const [durasi, setDurasi] = useState(0);
   const [totalEpisode, setTotalEpisode] = useState(0);
   const [genres, setGenres] = useState([]);
   const [message, setMessage] = useState("");
   const [colorMessage, setColorMessage] = useState("");
   const [isFilmLoaded, setIsFilmLoaded] = useState(false);

   useEffect(() => {
      const fetchFilm = async () => {
         try {
            const response = await getFilmById(id);
            const film = response.data.data;

            setJudul(film.judul);
            setStatus(film.status);
            setSinopsis(film.sinopsis);
            setSutradara(film.sutradara);
            setTanggalRilis(new Date(film.tanggal_rilis).toISOString().split("T")[0]);
            setDurasi(film.durasi);
            setTotalEpisode(film.total_episode);
            setGenres(film.genres);
         } catch (error) {
            console.error("Error fetching genres:", error.data.message);
         }
      };

      if (!isFilmLoaded) {
         fetchFilm();
         setIsFilmLoaded(true);
      }
   }, [isFilmLoaded, id]);

   const handleSubmit = async (e) => {
      e.preventDefault();
      try {
         const data = {
            judul: judul,
            status: status,
            sinopsis: sinopsis,
            sutradara: sutradara,
            tanggal_rilis: new Date(tanggalRilis).toISOString(),
            durasi: Number(durasi),
            total_episode: Number(totalEpisode),
         }

         const response = await updateFilm(id, data);
         if (response.status === 200) {
            setMessage(response.data.message);
            setColorMessage("text-green-600");
         }
      } catch (error) {
         console.log(error);
         setMessage(error.response.data.message);
         setColorMessage("text-red-600");
      }
   };

   const handleDeleteGenre = async (genreId) => {
      try {
         const response = await deleteGenre(genreId);
         if (response.status === 200) {
            setMessage(response.data.message);
            setColorMessage("text-green-600");
            setGenres(genres.filter((genre) => genre.id !== genreId));
         }
      } catch (error) {
         setMessage(error.response.data.error);
         setColorMessage("text-red-600");
      }
   };

   return (
      <>
         <h1 className="text-4xl font-bold mb-5">Update Film</h1>
         <form
            onSubmit={handleSubmit}
            className="w-full relative mb-10">
            <div className="w-full grid lg:grid-cols-2 gap-4">
               <Input
                  type="text"
                  placeholder="Masukkan Judul"
                  label="Judul"
                  value={judul}
                  onChange={(e) => setJudul(e.target.value)}
               />
               <Input
                  type="text"
                  label="Status"
                  placeholder="Pilih Status"
                  value={status}
                  onChange={(e) => setStatus(e.target.value)}
               />
               <Input
                  type="text"
                  label="Sutradara"
                  placeholder="Masukkan Sutradara"
                  value={sutradara}
                  onChange={(e) => setSutradara(e.target.value)}
               />
               <Input
                  type="date"
                  label="Tanggal Rilis"
                  placeholder="Masukkan tanggal rilis"
                  value={tanggalRilis}
                  onChange={(e) => setTanggalRilis(e.target.value)}
               />
               <Input
                  type="number"
                  label="Total Episode"
                  placeholder="Masukkan Total Episode"
                  value={totalEpisode}
                  onChange={(e) => setTotalEpisode(e.target.value)}
               />
               <Input
                  type="number"
                  label="Durasi"
                  placeholder="Masukkan Durasi"
                  value={durasi}
                  onChange={(e) => setDurasi(e.target.value)}
               />
            </div>
            <div className="mb-4">
               <label className="block text-text text-2xl font-semibold mb-2">Daftar Genre Aktif</label>
               <div className="flex flex-wrap gap-2">
                  {genres.map((genre) => (
                     <div
                        key={genre.id}
                        className="flex items-center px-3 py-1 text-sm rounded-full border text-slate-700 bg-slate-100 hover:bg-tertiary hover:text-white hover:border-tertiary transition-all duration-200 group"
                     >
                        <span>{genre.nama}</span>
                        <button
                           onClick={() => handleDeleteGenre(genre.id)}
                           type="button"
                           className="ml-2 text-slate-400 group-hover:text-white font-bold"
                        >
                           ×
                        </button>
                     </div>
                  ))}
               </div>
            </div>
            <TextArea
               label="Sinopsis"
               placeholder="Masukkan Sinopsis"
               value={sinopsis}
               onChange={(e) => setSinopsis(e.target.value)}
            />
            <Button
               type="submit"
               className="mt-4 w-40 py-2 rounded-lg absolute right-0"
            >
               Submit
            </Button>
            {message && <p className={`mt-16 absolute right-0 text-sm ${colorMessage}`}>{message}</p>}
         </form>
      </>
   )
}

export default UpdateFilmPage;
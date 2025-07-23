import Input from "../../../components/Input";
import Button from "../../../components/Button";
import { useState } from "react";
import { createGenre } from "../../../service/genre";

const AddGenre = () => {
   const [genreName, setGenreName] = useState('');
   const [message, setMessage] = useState('');
   const [colorMessage, setColorMessage] = useState('');

   const handleAddGenre = async (e) => {
      e.preventDefault();
      const data = {
         nama: genreName
      }
      try {
         const response = await createGenre(data);
         if (response.status === 201) {
            setMessage(response.data.message);
            setColorMessage("text-green-600");
            setGenreName('');
         }
      } catch (error) {
         setMessage(error.response.data.error);
         setColorMessage("text-red-600");
      }
   }

   return (
      <div className="lg:col-span-2 mt-8">
         <h2 className="text-2xl font-semibold text-text">Tambah Genre Baru</h2>
         <form onSubmit={handleAddGenre} className="max-w-96">
            <Input
               label="Name"
               type="text"
               placeholder="Masukkan Nama Genre"
               onChange={(e) => setGenreName(e.target.value)}
               value={genreName}
               className="w-96"
            />
            <Button
               type="submit"
               className="rounded-md w-40 mt-4"
            >
               Add Genre
            </Button>
            {message && <p className={`${colorMessage} mt-2`}>{message}</p>}
         </form>
      </div>
   )
}

export default AddGenre;
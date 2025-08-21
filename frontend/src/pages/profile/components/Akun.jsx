import { useEffect, useState } from "react";
import Button from "../../../components/Button";
// eslint-disable-next-line no-unused-vars
import { motion, AnimatePresence } from "framer-motion"
import { updateUser } from "../../../service/user";

const Akun = (props) => {
   const { user, review, watchlists } = props;
   const [isUpdating, setIsUpdating] = useState(false);
   const [nama, setNama] = useState(user?.nama);
   const [username, setUsername] = useState(user?.username);
   const [bio, setBio] = useState(user?.bio);
   const [message, setMessage] = useState("");
   const [colorMessage, setColorMessage] = useState("");

   useEffect(() => {
      setNama(user?.nama);
      setUsername(user?.username);
      setBio(user?.bio);
   }, [user]);

   const onUpdateProfile = () => {
      setIsUpdating(true);
   };

   const handleUpdateProfile = async () => {
      const data = {
         nama: nama,
         username: username,
         bio: bio,
      };

      if (user.id) {
         try {
            const response = await updateUser(user.id, data);
            setIsUpdating(false);
            setMessage(response.data.message);
            setColorMessage("text-green-600");
            return;
         } catch (error) {
            setMessage(error.response.data.error);
            setColorMessage("text-red-600");
         }
      }

      alert("User ID not found");
   }

   return (
      <div className="mt-28 flex justify-center px-4">
         <div className="w-full max-w-4xl bg-gradient-to-br from-indigo-50 to-white rounded-3xl shadow-xl p-8 space-y-6 transition-all duration-300">

            <div className="flex items-center gap-6">
               <div className="w-24 h-24 rounded-full bg-gradient-to-r from-indigo-500 to-purple-500 flex items-center justify-center text-white text-3xl font-bold shadow-lg">
                  {user?.nama?.charAt(0).toUpperCase() || "?"}
               </div>
               <div>
                  <h1 className="text-2xl font-bold text-gray-800">{user?.nama}</h1>
                  <p className="text-gray-500">@{user?.username}</p>
               </div>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6">
               <div className="bg-white/70 backdrop-blur-sm p-4 rounded-xl shadow-sm border border-gray-200">
                  <h2 className="text-lg font-semibold mb-1 text-indigo-600 flex items-center gap-2">
                     Informasi Pribadi
                  </h2>
                  <ul className="text-gray-600">
                     <li>
                        <span className="font-medium">Nama:</span>
                        <input
                           disabled={!isUpdating}
                           className={`ml-2 border-b outline-none px-1 transition-all duration-200 
                  ${isUpdating ? "border-indigo-400 focus:border-indigo-600" : "border-transparent bg-transparent text-gray-800"}`}
                           type="text"
                           value={nama}
                           onChange={(e) => setNama(e.target.value)}
                        />
                     </li>
                     <li>
                        <span className="font-medium">Username:</span>
                        <input
                           disabled={!isUpdating}
                           className={`ml-2 border-b outline-none px-1 transition-all duration-200 
                  ${isUpdating ? "border-indigo-400 focus:border-indigo-600" : "border-transparent bg-transparent text-gray-800"}`}
                           type="text"
                           value={username}
                           onChange={(e) => setUsername(e.target.value)}
                        />
                     </li>
                  </ul>
               </div>

               <div className="bg-white/70 backdrop-blur-sm p-4 rounded-xl shadow-sm border border-gray-200">
                  <h2 className="text-lg font-semibold mb-1 text-indigo-600">Aktivitas</h2>
                  <p className="text-gray-600">Total Ulasan: {review?.reviews?.length || 0}</p>
                  <p className="text-gray-600">Total Watchlist: {watchlists?.length || 0}</p>
               </div>
            </div>

            <div className="bg-white/70 backdrop-blur-sm p-4 rounded-xl shadow-sm border border-gray-200">
               <h2 className="text-lg font-semibold text-indigo-600">Tentang Saya</h2>
               <textarea
                  disabled={!isUpdating}
                  value={bio}
                  onChange={(e) => setBio(e.target.value)}
                  className={`w-full resize-none border rounded-lg transition-all duration-200 
            ${isUpdating ? "border-indigo-400 focus:border-indigo-600 p-2" : "border-transparent bg-transparent text-gray-600"}`}
               />
            </div>

            <div className="flex flex-col items-end gap-3">
               <AnimatePresence>
                  {isUpdating ? (
                     <motion.div
                        key="actions"
                        initial={{ opacity: 0, y: 15 }}
                        animate={{ opacity: 1, y: 0 }}
                        exit={{ opacity: 0, y: 15 }}
                        className="flex gap-3"
                     >
                        <Button
                           variant="outline"
                           onClick={() => setIsUpdating(false)}
                           className="rounded-xl px-5 py-2 shadow-md"
                        >
                           Cancel
                        </Button>
                        <Button
                           className="bg-indigo-600 hover:bg-indigo-700 text-white rounded-xl px-5 py-2 shadow-md"
                           onClick={handleUpdateProfile}
                        >
                           Save
                        </Button>
                     </motion.div>
                  ) : (
                     <motion.div
                        key="edit"
                        initial={{ opacity: 0, scale: 0.9 }}
                        animate={{ opacity: 1, scale: 1 }}
                        exit={{ opacity: 0, scale: 0.9 }}
                     >
                        <Button
                           className="bg-indigo-600 hover:bg-indigo-700 text-white rounded-xl px-5 py-2 shadow-md"
                           onClick={onUpdateProfile}
                        >
                           Edit Profile
                        </Button>
                     </motion.div>
                  )}
               </AnimatePresence>
               <p className={`${colorMessage} text-sm}`}>{message}</p>
            </div>
         </div>
      </div>
   )
}

export default Akun;
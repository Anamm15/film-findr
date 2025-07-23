// src/pages/AssignmentDashboardPage.jsx
import { useState, useEffect } from "react";

const AssignmentDashboardPage = () => {
   const [loading, setLoading] = useState(true);
   const [requests, setRequests] = useState([]);

   useEffect(() => {
      // Simulasi fetching data
      setTimeout(() => {
         setRequests([
            {
               id: 1,
               image: "https://via.placeholder.com/200x300",
               title: "Inception",
               releaseDate: "2010-07-16",
               status: "Pending",
               duration: "148 mins",
               genre: "Sci-fi",
               requestedBy: "Alice Johnson",
            },
            {
               id: 2,
               image: "https://via.placeholder.com/200x300",
               title: "Interstellar",
               releaseDate: "2014-11-07",
               status: "Approved",
               duration: "169 mins",
               genre: "Adventure",
               requestedBy: "Bob Smith",
            },
         ]);
         setLoading(false);
      }, 2000);
   }, []);

   const handleApprove = (id) => {
      console.log("Approve", id);
   };

   const handleReject = (id) => {
      console.log("Reject", id);
   };

   return (
      <div className="p-6 bg-gray-50 min-h-screen">
         <h1 className="text-3xl font-bold mb-6">Permintaan Pembuatan Film</h1>
         <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {loading
               ? Array.from({ length: 3 }).map((_, index) => (
                  <div
                     key={index}
                     className="bg-white p-4 rounded-xl shadow animate-pulse"
                  >
                     <div className="w-full h-72 bg-gray-300 rounded mb-4"></div>
                     <div className="h-4 bg-gray-300 rounded w-3/4 mb-2"></div>
                     <div className="h-4 bg-gray-300 rounded w-1/2 mb-2"></div>
                     <div className="h-4 bg-gray-300 rounded w-2/3 mb-2"></div>
                     <div className="h-4 bg-gray-300 rounded w-1/4"></div>
                  </div>
               ))
               : requests.map((req) => (
                  <div
                     key={req.id}
                     className="bg-white p-4 rounded-xl shadow flex flex-col"
                  >
                     <img
                        src={req.image}
                        alt={req.title}
                        className="w-full h-72 object-cover rounded mb-4"
                     />
                     <h2 className="text-xl font-semibold mb-2">{req.title}</h2>
                     <p className="text-sm text-gray-600 mb-1">
                        <strong>Tanggal Rilis:</strong> {req.releaseDate}
                     </p>
                     <p className="text-sm text-gray-600 mb-1">
                        <strong>Status:</strong> {req.status}
                     </p>
                     <p className="text-sm text-gray-600 mb-1">
                        <strong>Durasi:</strong> {req.duration}
                     </p>
                     <p className="text-sm text-gray-600 mb-1">
                        <strong>Genre:</strong> {req.genre}
                     </p>
                     <p className="text-sm text-gray-600 mb-4">
                        <strong>Diminta oleh:</strong> {req.requestedBy}
                     </p>
                     <div className="mt-auto flex justify-between gap-2">
                        <button
                           onClick={() => handleReject(req.id)}
                           className="px-4 py-2 text-sm font-medium text-white bg-red-500 hover:bg-red-600 rounded-lg"
                        >
                           Tolak
                        </button>
                        <button
                           onClick={() => handleApprove(req.id)}
                           className="px-4 py-2 text-sm font-medium text-white bg-green-500 hover:bg-green-600 rounded-lg"
                        >
                           Terima
                        </button>
                     </div>
                  </div>
               ))}
         </div>
      </div>
   );
};

export default AssignmentDashboardPage;

import { useState, useEffect } from "react";
import { getReviewByUserId } from "../api/review";
import { useParams } from "react-router-dom";
import { getUserById } from "../api/user";


const ProfilePage = () => {
    const [user, setUser] = useState(null);
    const [review, setReview] = useState([]);
    const [page, setPage] = useState(1);
    const id = useParams().id;

    useEffect(() => {
        const fetchUser = async () => {
            try {
                const response = await getUserById(id);
                setUser(response.data.data);
            } catch (error) {
                console.error("Error fetching user:", error.data.message);
            }
        };
        fetchUser();
    }, [id]);


    useEffect(() => {
        const fetchReview = async () => {
            try {
                const response = await getReviewByUserId(id, page);
                setReview(response.data.data);
            } catch (error) {
                console.error("Error fetching review:", error.data.message);
            }
        };

        fetchReview();
    }, [id, page]);

    return (
        <>
        <div>
            <div className="bg-white rounded-xl shadow p-5">
                <h2 className="text-xl font-semibold mb-4">Profil</h2>
                <div className="space-y-4">
                    <p className="font-semibold">Nama: {user?.nama}</p>
                    <p className="font-semibold">Username: {user?.username}</p>
                </div>
            </div>
        </div>
        
        <div className="bg-white rounded-xl shadow p-5">
            <h2 className="text-xl font-semibold mb-4">Review</h2>
            <div className="space-y-4">
              {review.reviews && review.reviews.map((r, idx) => (
                <div key={idx} className="border-b pb-3">
                  <p className="font-semibold">{r.user.username}</p>
                  <p className="text-gray-600">{r.komentar}</p>

                  <div className="flex items-center space-x-4 mt-2">
                    <svg
                        xmlns="http://www.w3.org/2000/svg"
                        className={`w-5 h-5 cursor-pointer ${
                        r.user_reaksi?.reaksi === "like" ? "text-blue-500" : "text-gray-400"
                        }`}
                        fill="currentColor"
                        viewBox="0 0 20 20"
                    >
                        <path d="M2 10c0-.6.4-1 1-1h3V5.5A2.5 2.5 0 018.5 3h1.1c.2 0 .4.1.6.3l3.4 3.4c.2.2.3.5.3.7v6.2c0 .3-.1.5-.3.7l-3.4 3.4c-.2.2-.4.3-.6.3h-1.1A2.5 2.5 0 016 15.5V13H3a1 1 0 01-1-1v-2z" />
                    </svg>

                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      className={`w-5 h-5 cursor-pointer ${
                        r.user_reaksi?.reaksi === "dislike" ? "text-red-500" : "text-gray-400"
                      }`}
                      viewBox="0 0 20 20"
                      fill="currentColor"
                    >
                      <path d="M18 10c0 .6-.4 1-1 1h-3v3.5a2.5 2.5 0 01-2.5 2.5h-1.086a1 1 0 01-.707-.293l-3.414-3.414a1 1 0 01-.293-.707V6.414a1 1 0 01.293-.707L9.707 2.293A1 1 0 0110.414 2H11.5A2.5 2.5 0 0114 4.5V7h3a1 1 0 011 1v2z" />
                    </svg>
                  </div>
                </div>
              ))}
            </div>

            {/* Pagination */}
            <div className="flex justify-center mt-6 space-x-2">
              {Array.from({ length: review.count_page }, (_, i) => (
                <button
                  key={i}
                  className="px-3 py-1 rounded bg-gray-100 hover:bg-gray-200 text-sm"
                  onClick={() => setPage(i + 1)}
                >
                  {i + 1}
                </button>
              ))}
            </div>
          </div>
        </>
    )
}

export default ProfilePage
import { ThumbsUp, ThumbsDown } from "react-feather";
import Pagination from "../components/Pagination";
import { updateReaksiReview } from "../service/review";

const ReviewLayout = (props) => {
   const { review, setPage, page } = props;

   const handleReact = async (id, reaksi) => {
      try {
         const payload = {
            reaksi: reaksi
         }
         const response = await updateReaksiReview(id, payload);
         console.log(response);
      } catch (error) {
         console.log(error);
      }
   }

   return (
      <div className="bg-white rounded-xl shadow p-2 mb-8">
         <h2 className="text-2xl md:text-3xl font-semibold mb-2 ps-4 pt-4">Review</h2>
         <div className="space-y-4">
            {review && review.reviews && review.reviews.map((r, idx) => (
               <div
                  key={idx}
                  className={`px-4 py-2 rounded-md ${idx % 2 === 0 ? "bg-white" : "bg-gray-100"
                     }`}
               >
                  <p className="font-semibold text-xl lg:text-2xl">{r.user.username}</p>
                  <p className="text-text lg:text-lg">{r.komentar}</p>

                  <div className="flex items-center gap-x-6 mt-2">
                     <ThumbsUp
                        className={`w-6 h-6 cursor-pointer ${r.user_reaksi && r.user_reaksi.reaksi === "like" ? "fill-green-500 text-green-800" : "text-green-500"}`}
                        onClick={() => handleReact(r.id, "like")}
                     />
                     <ThumbsDown
                        className={`w-6 h-6 mt-2 cursor-pointer ${r.user_reaksi && r.user_reaksi.reaksi === "dislike" ? "fill-red-500 text-red-900" : "text-red-500"}`}
                        onClick={() => handleReact(r.id, "dislike")}
                     />
                  </div>
               </div>
            ))}
         </div>

         <Pagination contents={review} page={page} setPage={setPage} />
      </div>
   )
}

export default ReviewLayout;
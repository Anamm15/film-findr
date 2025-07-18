import { ThumbsUp, ThumbsDown } from "react-feather";
import Pagination from "../components/Pagination";

const ReviewLayout = (props) => {
    const { review, setPage, page } = props;

    return (
        <div className="bg-white rounded-xl shadow p-2 mb-8">
            <h2 className="text-3xl font-semibold mb-4 ps-4 pt-4">Review</h2>
            <div className="space-y-4">
                {review && review.reviews && review.reviews.map((r, idx) => (
                    <div
                        key={idx}
                        className={`px-4 py-2 rounded-md ${idx % 2 === 0 ? "bg-white" : "bg-gray-100"
                            }`}
                    >
                        <p className="font-semibold text-2xl">{r.user.username}</p>
                        <p className="text-text text-lg">{r.komentar}</p>

                        <div className="flex items-center gap-x-6 mt-2">
                            {/* Like Icon */}
                            <ThumbsUp
                                className={`w-6 h-6 text-green-500 cursor-pointer ${review.reaksi_review && review.reaksi_review[idx] === "like" ? "text-green-500" : ""}`} />
                            {/* Dislike Icon */}
                            <ThumbsDown
                                className={`w-6 h-6 text-red-500 mt-2 cursor-pointer ${review.reaksi_review && review.reaksi_review[idx] === "dislike" ? "text-red-500" : ""}`} />
                        </div>
                    </div>
                ))}
            </div>

            <Pagination contents={review} page={page} setPage={setPage} />
        </div>
    )
}


export default ReviewLayout;
import { useState } from "react"
import Button from "../../../components/Button";
import { createUserFilm } from "../../../service/userFilm";
import { WATCH_LIST_STATUS } from "../../../utils/constant";

const WatchListForm = (props) => {
    const {id} = props
    const [watchListStatus, setWatchListStatus] = useState("");
    const [message, setMessage] = useState("");
    const [colorMessage, setColorMessage] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const data = {
                film_id: Number(id),
                status: watchListStatus
            }
            const response = await createUserFilm(data)
            if (response.status === 201) {
                setMessage(response.data.message);
                setColorMessage("text-green-600");
            }
        } catch (error) {
            setMessage(error.response.data.error);
            setColorMessage("text-red-600");
        }
    }

    return(
        <div className="absolute bottom-4 right-4">
            <form 
            className="rounded-xl flex flex-col sm:flex-row gap-4 sm:gap-5 items-center"
            onSubmit={handleSubmit}
            >
                <div className="relative max-w-sm text-lg">
                    <select
                    className="appearance-none px-4 py-[5px] pr-10 rounded-full border border-gray-300 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500 transition duration-200"
                    onChange={(e) => setWatchListStatus(e.target.value)}
                    >
                    <option value="">Pilih Status</option>
                    {WATCH_LIST_STATUS.map((status) => (
                        <option key={status} value={status}>
                        {status}
                        </option>
                    ))}
                    </select>

                    {/* Custom arrow */}
                    <div className="pointer-events-none absolute inset-y-0 right-4 flex items-center text-gray-500">
                    <svg
                        className="w-5 h-5"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                    >
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                    </svg>
                    </div>
                </div>
                <Button 
                    type="submit"
                    className="text-lg rounded-full"
                >
                    Tambah ke Watchlist
                </Button>
            </form>
            <p className={`${colorMessage} text-sm mt-1 ps-2`}>{message}</p>
        </div>
    )
}

export default WatchListForm
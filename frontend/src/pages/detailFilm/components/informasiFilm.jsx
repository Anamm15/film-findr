

const InformasiFilm = (props) => {
    const { film } = props;
    return (
        <div className="bg-white mb-6 space-y-3 text-lg text-secondary">
            <div className="flex flex-wrap gap-2 mb-2">
                {film.genres.map((genre) => (
                <span
                    key={genre.id}
                    className="bg-tertiary text-white px-3 py-1 rounded-full"
                >
                    {genre.nama}
                </span>
                ))}
            </div>
            <p><strong>Tanggal Rilis:</strong> {film.tanggal_rilis}</p>
            <p><strong>Durasi:</strong> {film.durasi} menit</p>
            <p><strong>Status:</strong> <span className="capitalize">{film.status}</span></p>
            <p><strong>Rating:</strong> ⭐ {film.rating}/10</p>
            <p><strong>Sutradara:</strong> {film.sutradara}</p>
            <p><strong>Total Episode:</strong> {film.total_episode}</p>
        </div>
    )
}

export default InformasiFilm
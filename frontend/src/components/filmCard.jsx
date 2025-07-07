import React from "react";

const FilmCard = ({ movie, onClick }) => {
  const {
    film_gambar,
    judul,
    genres,
    tanggal_rilis,
    // rating,
    status,
    durasi,
  } = movie;

  return (
    <div 
        className="max-w-sm bg-background rounded-2xl overflow-hidden shadow-lg hover:shadow-xl transition-shadow duration-300 cursor-pointer hover:border border-tertiary"
        onClick={onClick}
    >
      <img
        className="w-full h-[320px] object-cover"
        src={film_gambar[0].url}
        alt={judul}
      />
      <div className="px-4 pt-2 pb-4">
        <h2 className="text-xl font-semibold mb-1 truncate text-text">{judul}</h2>
        <div className="flex flex-wrap gap-2 text-sm text-white mb-2">
          {genres && genres.map((genre) => (
            <span
              key={genre.id}
              className="bg-tertiary px-2 py-0.5 rounded-full"
            >
              {genre.nama}
            </span>
          ))}
        </div>

        <div className="text-sm text-gray-600 space-y-1">
          <p><strong>Tanggal Rilis:</strong> {tanggal_rilis}</p>
          <p><strong>Durasi:</strong> {durasi} menit</p>
          <p><strong>Status:</strong> <span className="capitalize">{status}</span></p>
          <p><strong>Rating:</strong> ⭐ {2}/10</p>
        </div>
      </div>
    </div>
  );
};

export default FilmCard;

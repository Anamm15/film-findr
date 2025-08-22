const Gambar = (props) => {
    const { film } = props;
    return (
        <div className="w-max h-max mx-auto md:mx-0">
            {film.film_gambar.map((img) => (
                <img
                    key={img.id}
                    src={img.url}
                    alt={`Gambar ${img.id + 1}`}
                    className="rounded-lg shadow-md object-cover h-[450px] w-[300px]"
                />
            ))}
        </div>
    )
}

export default Gambar;
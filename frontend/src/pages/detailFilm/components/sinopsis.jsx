

const Sinopsis = (props) => {
    const { sinopsis } = props;
    return (
        <div className="bg-gray-100 rounded-xl p-6 mb-8">
            <h2 className="text-3xl font-semibold mb-2">Sinopsis</h2>
            <p className="text-text text-xl">{sinopsis}</p>
        </div>
    )
}

export default Sinopsis;
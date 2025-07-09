

const Akun = (props) => {
    const { user, review, watchlists } = props;
    return (
        <div className="mt-28 flex justify-center px-4">
            <div className="w-full max-w-4xl bg-gradient-to-br from-indigo-50 to-white rounded-3xl shadow-xl p-8 space-y-6">
                <div className="flex items-center gap-6">
                    <div className="w-24 h-24 rounded-full bg-gradient-primary flex items-center justify-center text-white text-3xl font-bold shadow-md">
                        {user?.nama?.charAt(0).toUpperCase() || "?"}
                    </div>
                    <div>
                        <h1 className="text-2xl font-bold text-text">{user?.nama}</h1>
                        <p className="text-secondary">@{user?.username}</p>
                    </div>
                </div>

                <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mt-6">
                    <div className="bg-background p-4 rounded-xl shadow-sm border border-gray-100">
                        <h2 className="text-lg font-semibold mb-2 text-primary">Informasi Pribadi</h2>
                        <ul className="text-secondary space-y-1">
                            <li><span className="font-medium">Nama:</span> {user?.nama}</li>
                            <li><span className="font-medium">Username:</span> {user?.username}</li>
                            {/* Tambahkan data lain di sini jika ada */}
                        </ul>
                    </div>

                    <div className="bg-background p-4 rounded-xl shadow-sm border border-gray-100">
                        <h2 className="text-lg font-semibold mb-2 text-primary">Aktivitas</h2>
                        <p className="text-secondary">
                            Total Ulasan: {review?.reviews?.length || 0}
                        </p>
                        <p className="text-secondary">
                            Total Watchlist: {watchlists?.length || 0}
                        </p>
                    </div>
                </div>

                <div className="bg-background p-4 rounded-xl shadow-sm border border-gray-100">
                    <h2 className="text-lg font-semibold mb-2 text-primary">Tentang Saya</h2>
                    <p className="text-secondary whitespace-pre-line">{user?.bio}</p>
                </div>
            </div>
        </div>
    )
}

export default Akun;
const PageLoading = ({ message }) => {
   return (
      <div className="fixed inset-0 flex flex-col items-center justify-center bg-background z-50">
         {/* Animasi Lingkaran */}
         <div className="relative w-20 h-20">
            <div className="absolute inset-0 rounded-full border-4 border-gray-300"></div>
            <div className="absolute inset-0 rounded-full border-4 border-tertiary border-t-transparent animate-spin"></div>
         </div>

         {/* Judul Animasi */}
         <h1 className="mt-6 text-2xl font-bold text-text animate-pulse tracking-widest">
            {message}
         </h1>

         {/* Progress bar di bawah */}
         <div className="mt-4 w-48 h-2 bg-gray-300 rounded-full overflow-hidden">
            <div className="h-2 bg-tertiary animate-[progress_2s_ease-in-out_infinite]"></div>
         </div>

         {/* Custom Keyframes Tailwind */}
         <style>
            {`
          @keyframes progress {
            0% { transform: translateX(-100%); }
            50% { transform: translateX(0); }
            100% { transform: translateX(100%); }
          }
        `}
         </style>
      </div>
   );
};

export default PageLoading;
const FilmCardSkeletonItem = () => (
   <div className="max-w-sm bg-background rounded-2xl overflow-hidden shadow-lg animate-pulse">
      <div className="w-full h-[260px] md:h-[320px] bg-gray-300"></div>

      <div className="px-4 pt-2 pb-4">
         <div className="h-5 bg-gray-300 rounded mb-2 w-3/4"></div>
         <div className="hidden md:flex flex-wrap gap-2 mb-2">
            <div className="h-4 w-16 bg-gray-300 rounded-full"></div>
            <div className="h-4 w-20 bg-gray-300 rounded-full"></div>
            <div className="h-4 w-14 bg-gray-300 rounded-full"></div>
         </div>
         <div className="hidden md:block space-y-2">
            <div className="h-4 bg-gray-300 rounded w-1/2"></div>
            <div className="h-4 bg-gray-300 rounded w-1/3"></div>
            <div className="h-4 bg-gray-300 rounded w-2/3"></div>
            <div className="h-4 bg-gray-300 rounded w-1/4"></div>
         </div>
         <div className="flex justify-between items-center text-sm md:hidden mt-2">
            <div className="h-4 w-16 bg-gray-300 rounded"></div>
            <div className="h-4 w-12 bg-gray-300 rounded"></div>
         </div>
      </div>
   </div>
);

const FilmCardSkeleton = ({ count = 1 }) => {
   return (
      <>
         {Array.from({ length: count }).map((_, i) => (
            <FilmCardSkeletonItem key={i} />
         ))}
      </>
   );
};

export default FilmCardSkeleton;

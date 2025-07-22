import React from "react";

const TopContributorsPlaceholder = () => {
   const dummyUsers = Array.from({ length: 5 });

   return (
      <div className="p-6">
         <h2 className="text-2xl font-semibold mb-4 text-gray-800">Top Contributors</h2>
         <div className="bg-white/30 backdrop-blur-lg shadow-xl rounded-2xl p-4 space-y-4">
            {dummyUsers.map((_, index) => (
               <div
                  key={index}
                  className="flex items-center space-x-4 animate-pulse bg-white/40 rounded-xl p-3 hover:shadow-md transition-all duration-300"
               >
                  {/* Avatar */}
                  <div className="w-12 h-12 rounded-full bg-gray-300 shadow-inner" />

                  {/* Info */}
                  <div className="flex-1 space-y-2">
                     <div className="h-4 bg-gray-300 rounded w-2/3" />
                     <div className="h-3 bg-gray-200 rounded w-1/3" />
                  </div>
               </div>
            ))}
         </div>
      </div>
   );
};

export default TopContributorsPlaceholder;

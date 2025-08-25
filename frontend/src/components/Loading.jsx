const Loading = ({ children }) => {
   return (
      <div className="flex flex-col items-center justify-center py-6">
         <div className="w-8 h-8 border-4 border-gray-300 border-t-tertiary rounded-full animate-spin"></div>
         <p className="mt-3 text-sm text-gray-600 animate-pulse">
            {children}
         </p>
      </div>
   );
};
export default Loading;
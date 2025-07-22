

const Pagination = (props) => {
   const { contents, page, setPage } = props;
   return (
      <div className="flex justify-center mt-4 pb-2 space-x-2">
         {contents && Array.from({ length: contents.count_page }, (_, i) => (
            <button
               key={i}
               className={`px-3 py-1 rounded bg-gray-100 hover:bg-gradient-secondary hover:text-white text-sm ${page === i + 1 ? "bg-gradient-primary text-white" : ""
                  }`}
               onClick={() => setPage(i + 1)}
            >
               {i + 1}
            </button>
         ))}
      </div>
   )
}

export default Pagination;
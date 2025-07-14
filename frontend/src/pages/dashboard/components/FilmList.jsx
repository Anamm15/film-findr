export default function FilmList({ title, films, icon }) {
   return (
      <div className="bg-background border border-slate-300 p-6 rounded-2xl shadow-lg">
         <div className="flex items-center space-x-2 mb-4">
            {icon}
            <h2 className="text-lg font-semibold text-text">{title}</h2>
         </div>
         <ul className="space-y-4">
            {films.map((film, index) => (
               <li key={film.id} className="flex items-center justify-between p-3 rounded-lg hover:bg-slate-500/50 transition-colors duration-200">
                  <div className="flex items-center space-x-2">
                     <span className="text-lg font-bold text-text w-4">{index + 1}</span>
                     <span className="font-medium text-text">{film.title}</span>
                  </div>
                  <span className="text-sm text-white bg-tertiary font-mono px-2 py-1 rounded">
                     {film.metric}
                  </span>
               </li>
            ))}
         </ul>
      </div>
   );
}
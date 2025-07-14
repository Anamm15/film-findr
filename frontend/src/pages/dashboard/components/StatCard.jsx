

const StatCard = (props) => {
   const { icon, title, value } = props;
   return (
      <div className="p-5 border border-slate-300 rounded-2xl flex items-center space-x-4 shadow-lg duration-300">
         <div className="p-3 rounded-lg">
            {icon}
         </div>
         <div>
            <p className="text-text font-medium">{title}</p>
            <p className="text-3xl font-bold text-secondary">{value}</p>
         </div>
      </div>
   );
}

export default StatCard
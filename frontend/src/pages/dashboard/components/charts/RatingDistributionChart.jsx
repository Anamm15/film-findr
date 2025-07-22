import { Pie } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";

ChartJS.register(ArcElement, Tooltip, Legend);

export const RatingDistributionChart = () => {
   const data = {
      labels: ["⭐️ 1", "⭐️ 2", "⭐️ 3", "⭐️ 4", "⭐️ 5"],
      datasets: [
         {
            label: "Distribusi Rating",
            data: [5, 12, 20, 40, 23],
            backgroundColor: ["#EF4444", "#F59E0B", "#EAB308", "#10B981", "#3B82F6"],
         },
      ],
   };

   return (
      <div className="bg-white p-4 shadow-md rounded-2xl">
         <h2 className="text-xl font-semibold mb-4">Distribusi Rating</h2>
         <Pie data={data} />
      </div>
   );
};

export default RatingDistributionChart;

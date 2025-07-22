import { Line } from "react-chartjs-2";
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Tooltip, Legend } from "chart.js";

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Tooltip, Legend);

export const ReviewsOverTimeChart = () => {
   const data = {
      labels: ["Jan", "Feb", "Mar", "Apr", "May"],
      datasets: [
         {
            label: "Review per Bulan",
            data: [50, 65, 80, 60, 90],
            borderColor: "#6366F1",
            backgroundColor: "#6366F1",
         },
      ],
   };

   const options = {
      responsive: true,
      scales: {
         y: { beginAtZero: true },
      },
   };

   return (
      <div className="bg-white p-4 shadow-md rounded-2xl">
         <h2 className="text-xl font-semibold mb-4">Review per Bulan</h2>
         <Line data={data} options={options} />
      </div>
   );
};

export default ReviewsOverTimeChart;
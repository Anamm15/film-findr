import { Bar } from "react-chartjs-2";
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Tooltip, Legend } from "chart.js";

ChartJS.register(CategoryScale, LinearScale, BarElement, Tooltip, Legend);

export const TopUsersChart = () => {
   const data = {
      labels: ["Alice", "Bob", "Charlie", "Diana", "Eve"],
      datasets: [
         {
            label: "Jumlah Review",
            data: [45, 38, 35, 30, 28],
            backgroundColor: "#4F46E5",
         },
      ],
   };

   const options = {
      responsive: true,
      plugins: {
         legend: { display: false },
      },
      scales: {
         y: { beginAtZero: true },
      },
   };

   return (
      <div className="bg-white p-4 shadow-md rounded-2xl">
         <h2 className="text-xl font-semibold mb-4">Top 5 User</h2>
         <Bar data={data} options={options} />
      </div>
   );
};

export default TopUsersChart;
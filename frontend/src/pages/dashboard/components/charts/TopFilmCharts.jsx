import { Bar } from "react-chartjs-2";
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Tooltip, Legend } from "chart.js";

ChartJS.register(CategoryScale, LinearScale, BarElement, Tooltip, Legend);

export const TopFilmsChart = () => {
   const data = {
      labels: ["Avengers", "Inception", "Titanic", "Interstellar", "Joker"],
      datasets: [
         {
            label: "Jumlah Review",
            data: [120, 100, 95, 90, 88],
            backgroundColor: "#10B981",
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
         <h2 className="text-xl font-semibold mb-4">Top 5 Film</h2>
         <Bar data={data} options={options} />
      </div>
   );
};

export default TopFilmsChart;
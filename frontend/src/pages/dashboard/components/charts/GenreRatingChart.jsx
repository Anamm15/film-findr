import { Bar } from "react-chartjs-2";
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Tooltip, Legend } from "chart.js";

ChartJS.register(CategoryScale, LinearScale, BarElement, Tooltip, Legend);

export const GenreRatingChart = () => {
   const data = {
      labels: ["Action", "Drama", "Comedy", "Sci-fi", "Romance"],
      datasets: [
         {
            label: "Rata-rata Rating",
            data: [4.2, 3.8, 4.0, 4.5, 3.9],
            backgroundColor: "#F97316",
         },
      ],
   };

   const options = {
      responsive: true,
      scales: {
         y: { beginAtZero: true, max: 5 },
      },
   };

   return (
      <div className="bg-white p-4 shadow-md rounded-2xl">
         <h2 className="text-xl font-semibold mb-4">Rata-rata Rating per Genre</h2>
         <Bar data={data} options={options} />
      </div>
   );
};

export default GenreRatingChart;
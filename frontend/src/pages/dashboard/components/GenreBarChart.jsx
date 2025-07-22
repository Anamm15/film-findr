import { Bar } from 'react-chartjs-2';
import {
   Chart as ChartJS,
   CategoryScale,
   LinearScale,
   BarElement,
   Title,
   Tooltip,
   Legend,
} from 'chart.js';

// Daftarkan elemen-elemen yang dibutuhkan oleh Bar Chart
ChartJS.register(
   CategoryScale,
   LinearScale,
   BarElement,
   Title,
   Tooltip,
   Legend
);

const GenreBarChart = ({ genres }) => {
   // Opsi untuk kustomisasi tampilan grafik
   const options = {
      indexAxis: 'y', // **Ini yang membuat grafik menjadi horizontal**
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
         legend: {
            display: false, // Kita tidak butuh legenda untuk satu set data
         },
         tooltip: {
            backgroundColor: '#1e293b', // Warna tooltip (slate-800)
            titleColor: '#cbd5e1',
            bodyColor: '#94a3b8',
         },
      },
      scales: {
         x: { // Sumbu horizontal (nilai)
            ticks: { color: '#64748b' }, // Warna angka (slate-500)
            grid: { color: '#e2e8f0' }, // Warna garis grid (slate-200)
            title: {
               display: true,
               text: 'Jumlah Film',
               color: '#334155'
            }
         },
         y: { // Sumbu vertikal (kategori/genre)
            ticks: { color: '#334155', font: { weight: '600' } }, // Warna & ketebalan teks (slate-700)
            grid: { display: false }, // Hilangkan garis grid vertikal
         },
      },
   };

   // Memproses data genre untuk grafik
   const chartData = {
      labels: genres.map(genre => genre.nama),
      datasets: [
         {
            label: 'Jumlah Film',
            data: genres.map(genre => genre.jumlahFilm),
            backgroundColor: [ // Warna-warni untuk setiap bar
               'rgba(59, 130, 246, 0.7)',  // blue-500
               'rgba(16, 185, 129, 0.7)', // emerald-500
               'rgba(239, 68, 68, 0.7)',  // red-500
               'rgba(249, 115, 22, 0.7)', // orange-500
               'rgba(139, 92, 246, 0.7)', // violet-500
               'rgba(236, 72, 153, 0.7)', // pink-500
               'rgba(20, 184, 166, 0.7)', // teal-500
               'rgba(245, 158, 11, 0.7)'  // amber-500
            ],
            borderColor: [ // Warna border untuk setiap bar
               'rgb(59, 130, 246)',
               'rgb(16, 185, 129)',
               'rgb(239, 68, 68)',
               'rgb(249, 115, 22)',
               'rgb(139, 92, 246)',
               'rgb(236, 72, 153)',
               'rgb(20, 184, 166)',
               'rgb(245, 158, 11)'
            ],
            borderWidth: 1,
            borderRadius: 5, // Membuat sudut bar menjadi tumpul
         },
      ],
   };

   return <Bar options={options} data={chartData} />;
};

export default GenreBarChart;
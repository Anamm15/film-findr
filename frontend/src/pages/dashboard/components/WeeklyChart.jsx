

// src/components/WeeklyChart.jsx

import { Line } from 'react-chartjs-2';
import {
   Chart as ChartJS,
   CategoryScale,
   LinearScale,
   PointElement,
   LineElement,
   Title,
   Tooltip,
   Legend,
   Filler,
} from 'chart.js';

ChartJS.register(
   CategoryScale,
   LinearScale,
   PointElement,
   LineElement,
   Title,
   Tooltip,
   Legend,
   Filler
);

const WeeklyChart = ({ title, data: chartData, icon, borderColor, backgroundColor }) => {
   const options = {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
         legend: { display: false },
         tooltip: {
            backgroundColor: '#1e293b', // slate-800
            titleColor: '#cbd5e1', // slate-300
            bodyColor: '#94a3b8', // slate-400
            padding: 10,
            cornerRadius: 8,
            displayColors: false,
         },
      },
      scales: {
         x: {
            ticks: { color: '#94a3b8' }, // slate-400
            grid: { color: '#334155' }, // slate-700
         },
         y: {
            ticks: { color: '#94a3b8' }, // slate-400
            grid: { color: '#334155' }, // slate-700
            beginAtZero: true,
         },
      },
      elements: {
         line: {
            tension: 0.4, // Membuat garis lebih melengkung halus
         },
      },
   };

   const data = {
      labels: chartData.labels,
      datasets: [
         {
            label: title,
            data: chartData.values,
            fill: true,
            borderColor: borderColor,
            backgroundColor: backgroundColor,
            pointBackgroundColor: borderColor,
            pointHoverRadius: 7,
            pointHoverBackgroundColor: '#ffffff',
            pointHoverBorderColor: borderColor,
         },
      ],
   };

   return (
      <div className="bg-background border border-slate-700 p-6 rounded-2xl shadow-lg col-span-1 md:col-span-2">
         <div className="flex items-center space-x-2 mb-4">
            {icon}
            <h2 className="text-lg font-semibold text-white">{title}</h2>
         </div>
         <div className="h-64">
            <Line options={options} data={data} />
         </div>
      </div>
   );
}

export default WeeklyChart
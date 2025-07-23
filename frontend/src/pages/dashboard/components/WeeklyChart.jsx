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
            backgroundColor: '#1e293b',
            titleColor: '#cbd5e1',
            bodyColor: '#94a3b8',
            padding: 10,
            cornerRadius: 8,
            displayColors: false,
         },
      },
      scales: {
         x: {
            ticks: { color: '#94a3b8' },
            grid: { color: '#334155' },
         },
         y: {
            ticks: { color: '#94a3b8' },
            grid: { color: '#334155' },
            beginAtZero: true,
         },
      },
      elements: {
         line: {
            tension: 0.4,
         },
      },
   };

   const data = {
      labels: chartData?.labels,
      datasets: [
         {
            label: title,
            data: chartData?.values,
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
      <div className="bg-background border border-slate-300 p-6 rounded-2xl shadow-lg col-span-1 md:col-span-2">
         <div className="flex items-center space-x-2 mb-4">
            {icon}
            <h2 className="text-lg font-semibold text-white">{title}</h2>
         </div>
         <div className="h-64">
            {
               data && <Line options={options} data={data} />
            }
         </div>
      </div>
   );
}

export default WeeklyChart
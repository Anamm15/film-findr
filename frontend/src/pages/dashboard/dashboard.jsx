import { Users, Star, TrendingUp, Award, Film, LineChart } from 'lucide-react';
import StatCard from './components/StatCard';
import WeeklyChart from './components/WeeklyChart';
import FilmList from './components/FilmList';

const weeklyReviewData = {
   labels: ['Minggu 1', 'Minggu 2', 'Minggu 3', 'Minggu 4', 'Minggu 5', 'Minggu 6'],
   values: [65, 59, 80, 81, 56, 95],
};

const weeklyUserData = {
   labels: ['Minggu 1', 'Minggu 2', 'Minggu 3', 'Minggu 4', 'Minggu 5', 'Minggu 6'],
   values: [30, 42, 55, 60, 78, 88],
};

const trendingFilms = [
   {
      id: 1,
      title: 'Dune: Part Two',
      metric: '150,234 views',
      metricValue: 150234,
      metricType: 'views',
   },
   {
      id: 2,
      title: 'Godzilla x Kong: The New Empire',
      metric: '135,876 views',
      metricValue: 135876,
      metricType: 'views',
   },
   {
      id: 3,
      title: 'The Fall Guy',
      metric: '110,453 views',
      metricValue: 110453,
      metricType: 'views',
   },
   {
      id: 4,
      title: 'Kingdom of the Planet of the Apes',
      metric: '98,211 views',
      metricValue: 98211,
      metricType: 'views',
   },
   {
      id: 5,
      title: 'Civil War',
      metric: '85,789 views',
      metricValue: 85789,
      metricType: 'views',
   },
];

const topFilms = [
   {
      id: 1,
      title: 'The Shawshank Redemption',
      metric: '9.3 Rating',
      metricValue: 9.3,
      metricType: 'rating',
   },
   {
      id: 2,
      title: 'The Godfather',
      metric: '9.2 Rating',
      metricValue: 9.2,
      metricType: 'rating',
   },
   {
      id: 3,
      title: 'The Dark Knight',
      metric: '9.0 Rating',
      metricValue: 9.0,
      metricType: 'rating',
   },
   {
      id: 4,
      title: 'Pulp Fiction',
      metric: '8.9 Rating',
      metricValue: 8.9,
      metricType: 'rating',
   },
   {
      id: 5,
      title: 'Forrest Gump',
      metric: '8.8 Rating',
      metricValue: 8.8,
      metricType: 'rating',
   },
];



const DashboardPage = () => {
   return (
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

         {/* Kartu Statistik */}
         <StatCard title="Total Pengguna" value="14,287" icon={<Users className="w-10 h-10 text-cyan-400" />} />
         <StatCard title="Total Review" value="89,541" icon={<Star className="w-10 h-10 text-amber-400" />} />

         {/* Grafik */}
         <WeeklyChart
            title="Review per Minggu"
            data={weeklyReviewData}
            icon={<LineChart className="text-violet-400" />}
            borderColor="#a78bfa" // violet-400
            backgroundColor="#a78bfa33"
         />
         <WeeklyChart
            title="Pendaftaran per Minggu"
            data={weeklyUserData}
            icon={<Users className="text-emerald-400" />}
            borderColor="#34d399" // emerald-400
            backgroundColor="#34d39933"
         />

         {/* Daftar Film */}
         <FilmList title="Film Trending Minggu Ini" films={trendingFilms} icon={<TrendingUp className="text-rose-500" />} />
         <FilmList title="Top Film Sepanjang Masa" films={topFilms} icon={<Award className="text-amber-400" />} />

      </div>
   )
}

export default DashboardPage;
import { Users, Star, TrendingUp, Award, Film, LineChart } from 'lucide-react';
import StatCard from './components/StatCard';
import WeeklyChart from './components/WeeklyChart';
import FilmList from './components/FilmList';
import { useEffect, useState } from 'react';
import { getDashboard } from '../../service/dashboard';
import { convertTimestampToDate } from '../../utils/helpers';

const DashboardPage = () => {
   const [stats, setStats] = useState(null);
   const [weeklyUsers, setWeeklyUsers] = useState(null);
   const [weeklyReviews, setWeeklyReviews] = useState(null);

   useEffect(() => {
      const fetchStats = async () => {
         try {
            const response = await getDashboard();
            if (response.status === 200) {
               setStats(response.data.data);

               const rawWeeklyUsers = response.data.data.weekly_users;
               const rawWeeklyReviews = response.data.data.weekly_reviews;

               setWeeklyUsers({
                  labels: rawWeeklyUsers.map((item) => {
                     const convertedDate = convertTimestampToDate(item.label);
                     return convertedDate;
                  }),
                  values: rawWeeklyUsers.map((item) => item.value)
               });
               setWeeklyReviews({
                  labels: rawWeeklyReviews.map((item) => {
                     const convertedDate = convertTimestampToDate(item.label);
                     return convertedDate;
                  }),
                  values: rawWeeklyReviews.map((item) => item.value)
               });
            }
         } catch (error) {
            console.error('Error fetching stats:', error);
         }
      };

      fetchStats();
   }, []);

   return (
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">

         {/* Kartu Statistik */}
         <StatCard title="Total Users" value={stats?.count_users} icon={<Users className="w-10 h-10 text-cyan-400" />} />
         <StatCard title="Total Reviews" value={stats?.count_review} icon={<Star className="w-10 h-10 text-amber-400" />} />

         {/* Grafik */}
         <WeeklyChart
            title="Pendaftaran per Minggu"
            data={weeklyUsers}
            icon={<Users className="text-emerald-400" />}
            borderColor="#34d399"
            backgroundColor="#34d39933"
         />
         <WeeklyChart
            title="Review per Minggu"
            data={weeklyReviews}
            icon={<LineChart className="text-violet-400" />}
            borderColor="#a78bfa"
            backgroundColor="#a78bfa33"
         />

         {/* Daftar Film */}
         <FilmList title="Film Trending Minggu Ini" films={stats?.trending_films} icon={<TrendingUp className="text-rose-500" />} />
         <FilmList title="Top Film Sepanjang Masa" films={stats?.top_films} icon={<Award className="text-amber-400" />} />

      </div>
   )
}

export default DashboardPage;
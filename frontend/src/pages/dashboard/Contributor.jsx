import TopUsersChart from "./components/charts/TopUsersChart";
import TopFilmsChart from "./components/charts/TopFilmCharts";
import RatingDistributionChart from "./components/charts/RatingDistributionChart";
import ReviewsOverTimeChart from "./components/charts/ReviewOverTimeChart";
import GenreRatingChart from "./components/charts/GenreRatingChart";
import TopContributorsPlaceholder from "./components/TopContributor";

const ContributorDashboardPage = () => {
   return (
      <div className="p-6 bg-gray-50 min-h-screen">
         <h1 className="text-3xl font-bold mb-6">Review Dashboard</h1>
         <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
            <TopContributorsPlaceholder />
            <TopContributorsPlaceholder />
         </div>
      </div>
   )
}

export default ContributorDashboardPage;
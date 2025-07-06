import { Routes, Route } from "react-router-dom";
import LandingPage from "../pages/landing/landing";
import LoginPage from "../pages/login/login";
import RegisterPage from "../pages/register/register";
import DetailFilmPage from "../pages/detailFilm/detailFilm";
import DashboardPage from "../pages/dashboard/dashboard";
import ProfilePage from "../pages/profile/profile";
import ProtectedRoute from "../utils/protectedRoute";
import WatchListPage from "../pages/watchlist/watchlist";

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={<LandingPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/film/:id" element={<DetailFilmPage />} />
      <Route path="/profile/:id" element={<ProfilePage />} />
      <Route 
        path="/watchlist" 
        element={
            <ProtectedRoute>
                <WatchListPage />
            </ProtectedRoute>} 
      />
      <Route
        path="/dashboard"
        element={
            <ProtectedRoute>
                <DashboardPage />
            </ProtectedRoute>
        }
      />
    </Routes>
  );
};


export default AppRoutes;
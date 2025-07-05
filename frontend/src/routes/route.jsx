import { createBrowserRouter } from "react-router-dom";
import LandingPage from "../pages/landing/landing";
import LoginPage from "../pages/login/login";
import RegisterPage from "../pages/register/register";
import DetailFilmPage from "../pages/detailFilm/detailFilm";
import DashboardPage from "../pages/dashboard/dashboard";
import ProfilePage from "../pages/profile/profile";

const router = createBrowserRouter([{
    path: "/",
    element: <LandingPage />
},
{
    path: "/login",
    element: <LoginPage />
},
{
    path: "/register",
    element: <RegisterPage />
},
{
    path: "/film/:id",
    element: <DetailFilmPage />
},
{
    path: "/dashboard",
    element: <DashboardPage />
},
{
    path: "/profile/:id",
    element: <ProfilePage />
}
])

export default router
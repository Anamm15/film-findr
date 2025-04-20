import { createBrowserRouter } from "react-router-dom";
import LoginPage from "../pages/login";
import RegisterPage from "../pages/register";
import LandingPage from "../pages/landing";
import DetailFilmPage from "../pages/detailFilm";
import DashboardPage from "../pages/dashboard";
import ProfilePage from "../pages/profile";

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
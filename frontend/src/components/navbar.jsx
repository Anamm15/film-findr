import { useState, useEffect } from "react";
import { NavLink, useLocation } from "react-router-dom";
import { logoutUser } from "../service/user";
import { useContext } from "react";
import { AuthContext } from "../contexts/authContext";
import Button from "./Button";
import "../styles/Navbar.css";

const Navbar = () => {
   const { user } = useContext(AuthContext);
   const [isMenuOpen, setIsMenuOpen] = useState(false);
   const location = useLocation();

   useEffect(() => {
      setIsMenuOpen(false);
   }, [location]);

   useEffect(() => {
      if (isMenuOpen) {
         document.body.style.overflow = 'hidden';
      } else {
         document.body.style.overflow = 'auto';
      }
      // Cleanup function
      return () => {
         document.body.style.overflow = 'auto';
      };
   }, [isMenuOpen]);

   const hanldeLogout = async () => {
      try {
         const response = await logoutUser();
         console.log(response);

         if (response.status === 200) {
            window.location.reload();
         }
      } catch (error) {
         console.log(error);
      }
   }

   const navLinks = user ? (
      <>
         <NavLink to="/" className="nav-link" end>Home</NavLink>
         <NavLink to="/top-films" className="nav-link" end>Top Film</NavLink>
         <NavLink to={`/profile/${user?.username}`} className="nav-link">Profile</NavLink>
         <NavLink to="/watchlist" className="nav-link">Watch List</NavLink>
      </>
   ) : null;

   return (
      <>
         {/* Navbar Utama */}
         <nav className="fixed z-50 w-full bg-background backdrop-blur-sm border-b border-gray-200/50 h-20 top-0 left-0 flex justify-center items-center">
            <div className="flex justify-between items-center w-full max-w-[1280px] px-4">
               {/* Logo */}
               <div className="text-3xl font-bold cursor-pointer text-primary z-50">
                  <NavLink to="/">Film-Findr</NavLink>
               </div>

               {/* Menu Navigasi Desktop (sembunyi di mobile) */}
               {user && (
                  <div className="hidden md:flex items-center gap-10 font-bold text-xl text-primary">
                     {navLinks}
                  </div>
               )}

               {/* Tombol Auth Desktop (sembunyi di mobile) */}
               <div className="hidden md:flex">
                  {user ? (
                     <Button onClick={hanldeLogout} className="w-28 py-2 rounded-lg font-semibold">Logout</Button>
                  ) : (
                     <div className="flex gap-5">
                        <Button className="w-28 py-2 rounded-lg font-semibold"><NavLink to="/login">Login</NavLink></Button>
                        <Button className="w-28 py-2 rounded-lg font-semibold"><NavLink to="/register">Register</NavLink></Button>
                     </div>
                  )}
               </div>

               {/* Tombol Hamburger (hanya tampil di mobile) */}
               <div className="md:hidden z-50">
                  <button onClick={() => setIsMenuOpen(!isMenuOpen)} className="hamburger-button">
                     <span className={`hamburger-line ${isMenuOpen ? 'open' : ''}`}></span>
                     <span className={`hamburger-line ${isMenuOpen ? 'open' : ''}`}></span>
                     <span className={`hamburger-line ${isMenuOpen ? 'open' : ''}`}></span>
                  </button>
               </div>
            </div>
         </nav>

         {/* Overlay Latar Belakang */}
         <div
            onClick={() => setIsMenuOpen(false)}
            className={`fixed inset-0 bg-black/50 z-30 transition-opacity duration-300
                    ${isMenuOpen ? 'opacity-100' : 'opacity-0 pointer-events-none'}`}
         />

         {/* Mobile Menu Panel (Slide-in) */}
         <div
            className={`fixed bottom-0 right-0 h-full w-4/5 max-w-sm bg-background z-40
                    transform transition-transform pt-20 duration-500 ease-in-out md:hidden
                    ${isMenuOpen ? 'translate-x-0' : 'translate-x-full'}`}
         >
            <div className="flex flex-col h-full p-5">
               <nav className="flex flex-col flex-grow text-2xl font-semibold">
                  <ul className={`mobile-menu-list ${isMenuOpen ? 'visible' : ''}`}>
                     {user ? (
                        <>
                           <li><NavLink to="/" end>Home</NavLink></li>
                           <li className="mt-4"><NavLink to="/top-films" end>Top Film</NavLink></li>
                           <li className="mt-4"><NavLink to={`/profile/${user?.username}`}>Profile</NavLink></li>
                           <li className="mt-4"><NavLink to="/watchlist">Watch List</NavLink></li>
                        </>
                     ) : (
                        <>
                           <Button className="w-full py-0.5 text-lg rounded-xl"><NavLink to="/login">Login</NavLink></Button>
                           <Button className="w-full py-0.5 text-lg rounded-xl mt-2"><NavLink to="/register">Register</NavLink></Button>
                        </>
                     )}
                  </ul>
               </nav>

               {user && (
                  <div className={`mobile-menu-list ${isMenuOpen ? 'visible' : ''}`}>
                     <Button onClick={hanldeLogout} className="w-full py-0.5 rounded-xl text-lg">Logout</Button>
                  </div>
               )}
            </div>
         </div>
      </>
   );
};

export default Navbar;
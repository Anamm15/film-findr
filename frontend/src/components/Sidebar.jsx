

// components/Sidebar.jsx
import { Link, useLocation } from "react-router-dom";

const menu = [
  { name: "Dashboard", path: ["/dashboard"] },
  { name: "Films", path: ["/dashboard/films", "/dashboard/films/new"] },
  { name: "Genres", path: ["/dashboard/genres"] },
  { name: "Reviews", path: ["/dashboard/reviews"] },
  { name: "Contributors", path: ["/dashboard/contributors"] },
  { name: "Assingments", path: ["/dashboard/assignments"] },
];

export default function Sidebar() {
  const location = useLocation();
  return (
    <aside className="w-64 pt-28 bg-background h-screen border-r border-gray-200 p-4">
      <h1 className="text-2xl font-bold mb-6">Dashboard</h1>
      <nav className="space-y-2">
        {menu.map((item) => (
          <Link
            key={item.name}
            to={item.path[0]}
            className={`block px-4 py-2 rounded-lg ${
              item.path.includes(location.pathname)
                ? "bg-blue-500 text-white"
                : "text-gray-700 hover:bg-gray-100"
            }`}
          >
            {item.name}
          </Link>
        ))}
      </nav>
    </aside>
  );
}

import { Outlet } from "react-router-dom";
import Sidebar from "../components/Sidebar";

export default function DashboardLayout() {
  return (
    <div className="flex min-h-screen xl:max-w-[1280px] mx-auto">
      <Sidebar />
      <div className="flex-1 flex flex-col bg-background">
        <main className="p-6 flex-1 overflow-y-auto pt-28 relative">
          <Outlet />
        </main>
      </div>
    </div>
  );
}

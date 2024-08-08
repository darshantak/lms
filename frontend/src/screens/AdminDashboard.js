import React, { useState } from "react";
import { ReactDOM } from "react-dom/client";
import Sidebar from "../components/Sidebar";
import {createColumnHelper,flexRender,getCoreRowModel,useReactTable} from "@tanstack/react-table";


const AdminDashboard = () => {
  const [sidebarOpen, setSidebarOpen] = useState(true);

  return (
    <div className="flex">
      <Sidebar open={sidebarOpen} setOpen={setSidebarOpen}/>
      <div className={`flex-1 ${sidebarOpen ? "ml-72" : "ml-10"} p-8 bg-gray-100 min-h-screen`}>
        <h1>Hello World</h1>
      </div>
    </div>
  );
};

export default AdminDashboard;

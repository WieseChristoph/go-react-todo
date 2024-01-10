import { type FC } from "react";

import Navbar from "~/components/Navigation/Navbar";
import { Outlet } from "@tanstack/react-router";

const Root: FC = () => {
  return (
    <>
      <Navbar />
      <main className="p-3">
        <Outlet />
      </main>
    </>
  );
};

export default Root;

import React from "react";
import Authorization from "./pages/Authorization";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Home from "./pages/Home";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Authorization />,
  },
  {
    path: "homepage",
    element: <Home></Home>,
  },
]);

function App() {
  return <RouterProvider router={router}></RouterProvider>;
}

export default App;

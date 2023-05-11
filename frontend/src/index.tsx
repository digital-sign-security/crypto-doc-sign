import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import LoadingIndicator from "./components/LoadingIndicator";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <App /> <LoadingIndicator></LoadingIndicator>
  </React.StrictMode>
);

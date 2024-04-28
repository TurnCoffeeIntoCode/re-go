import React from "react";
import ReactDOM from "react-dom/client";
import App from "./page";

const root = ReactDOM.hydrateRoot(document.getElementById("app"), <App {...((window as any).PROPS || {})} />);

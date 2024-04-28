import React from "react";
import { renderToString } from "react-dom/server";
import App from "./page";

function renderApp(props) {
  return renderToString(<App {...props} />);
}

globalThis.renderApp = renderApp;

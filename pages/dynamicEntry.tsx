import React from "react";
import { renderToString } from "react-dom/server";
import App from "./page"; // Import all your components here
// import OtherComponent from './OtherComponent';

// Create a mapping of component names to components
const components = {
  App, // Add all your components to this object
  // OtherComponent,
};

function renderApp(componentName, props) {
  const Component = components[componentName];
  if (!Component) {
    throw new Error(`No component found for name ${componentName}`);
  }
  return renderToString(<Component {...props} />);
}

globalThis.renderApp = renderApp;

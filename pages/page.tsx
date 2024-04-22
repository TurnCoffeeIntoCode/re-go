import React from "react";

function App(props) {
  console.log("APP rendered", props);
  return (
    <div>
      <h1>{props.Name}</h1>
    </div>
  );
}

export default App;

import React from "react";

function App(props) {
  console.log("APP rendered", props);
  return (
    <main>
      <h1>Rendered on Server by ReGo</h1>
      {props.numbers.map((row, index) => (
        <p key={index}>{row}</p>
      ))}
    </main>
  );
}

export default App;

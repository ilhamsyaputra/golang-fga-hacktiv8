import axios from "axios";
import { useState, useEffect } from "react";

const App = () => {
  const [data, setData] = useState();

  useEffect(() => {
    axios.get("http://localhost:8080/api/weather").then((response) => {
       setData(response.data);
    })
  }, [])

  useEffect(() => {
    const interval = setInterval(() => (
      axios.get("http://localhost:8080/api/weather").then((response) => {
        setData(response.data);
      })
    ), 3000);
    return () => {
      clearInterval(interval);
    };
  }, []);

  if (!data) return null;

  return (
    <>
      <h1>
        {data.status.wind}
      </h1>
    </>
  )
}

export default App;
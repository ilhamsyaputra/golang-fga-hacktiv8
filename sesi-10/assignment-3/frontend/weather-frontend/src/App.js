import axios from "axios";
import { useState, useEffect } from "react";

const getWaterStatus = (data) => {
  switch (true) {
    case data.status.water <= 5:
      return "aman"
    case 6 <= data.status.water <= 8:
      return "siaga"
    case data.status.water > 8:
      return "bahaya"
  }
}

const getWindStatus = (data) => {
  switch (true) {
    case data.status.wind <= 6:
      return "aman"
    case 7 <= data.status.wind <= 15:
      return "siaga"
    case data.status.wind > 15:
      return "bahaya"
  }
}

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
    ), 15000);
    return () => {
      clearInterval(interval);
    };
  }, []);

  if (!data) return null;

  return (
    <>
      <h1>
        Assignment-3 - M Ilham Syaputra
      </h1>
      <h3>
        Wind: {data.status.wind} m/s ({getWindStatus(data)})
      </h3>
      <h3>
        Water: {data.status.water} m ({getWaterStatus(data)})
      </h3>
    </>
  )
}

export default App;
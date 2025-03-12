import React, { useEffect, useState } from "react";
import { SensorData } from "../../domain/entities/sensorData";
import WebSocketAdapter from "../adapters/websocketAdapter";
import { GetSensorData } from "../../application/usecases/getSensorData";
import { SensorController } from "../controllers/sensorController";
import WS_URL from "../config/webSocket";

const styles = {
    container: {
        textAlign: "center",
        padding: "20px",
        fontFamily: "Arial, sans-serif"
    },
    title: {
        color: "#333",
        fontSize: "24px",
        marginBottom: "20px"
    },
    table: {
        width: "80%",
        margin: "auto",
        borderCollapse: "collapse",
        boxShadow: "0px 4px 8px rgba(0, 0, 0, 0.1)",
        borderRadius: "8px",
        overflow: "hidden",
        backgroundColor: "#f9f9f9"
    },
    th: {
        backgroundColor: "#4CAF50",
        color: "white",
        padding: "10px",
        textAlign: "center"
    },
    td: {
        padding: "12px",
        borderBottom: "1px solid #ddd",
        textAlign: "center"
    },
    tr: {
        transition: "background-color 0.3s",
    },
    noData: {
        fontSize: "18px",
        color: "#777"
    }
};

styles.tr[":hover"] = {
    backgroundColor: "#e0f7fa"
};

const SensorTable = () => {
  const [sensorData, setSensorData] = useState([]);

  useEffect(() => {
    // Conectamos WebSocket
    const webSocketAdapter = new WebSocketAdapter();
    webSocketAdapter.connect(WS_URL);

    // Inicializamos el caso de uso y controlador
    const getSensorData = new GetSensorData({
      getAll: () => Promise.resolve(sensorData),
      store: (data) => {
        setSensorData((prevData) => [...prevData, new SensorData(data.temperature, data.humidity)]);
      },
    });

    new SensorController(webSocketAdapter, getSensorData);

    // Cleanup on unmount
    return () => {
      webSocketAdapter.socket.close();
    };
  }, [sensorData]);

  return (
    <div style={styles.container}>
      <h2 style={styles.title}>Sensor Data</h2>
      <table style={styles.table}>
        <thead>
          <tr>
            <th style={styles.th}>Temperature</th>
            <th style={styles.th}>Humidity</th>
          </tr>
        </thead>
        <tbody>
          {sensorData.length === 0 ? (
            <tr>
              <td colSpan="2" style={styles.noData}>No data available</td>
            </tr>
          ) : (
            sensorData.map((data, index) => (
              <tr key={index} style={styles.tr}>
                <td style={styles.td}>{data.temperature}</td>
                <td style={styles.td}>{data.humidity}</td>
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
};

export default SensorTable;

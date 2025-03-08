// src/components/SensorDataDisplay.js

import React, { useEffect, useState } from "react";
import { GetSensorDataUseCase } from "../application/useCases/GetSensorDatauseCase"; 
import { SensorService } from "../domain/services/SensorService";
import { Api } from "../infraestructure/api/api"; 

const SensorDataDisplay = () => {
    const [sensorData, setSensorData] = useState([]);

    useEffect(() => {
        const api = new Api();
        const sensorService = new SensorService(api);
        const useCase = new GetSensorDataUseCase(sensorService);

        const fetchData = async () => {
            const data = await useCase.execute();
            setSensorData(data);
        };

        const interval = setInterval(fetchData, 5000);
        fetchData(); // Llamada inicial

        return () => clearInterval(interval);
    }, []);

    return (
        <div style={styles.container}>
            <h1 style={styles.title}>Sensor Data</h1>
            {sensorData.length > 0 ? (
                <table style={styles.table}>
                    <thead>
                        <tr>
                            <th style={styles.th}>ID</th>
                            <th style={styles.th}>Temperature (°C)</th>
                            <th style={styles.th}>Humidity (%)</th>
                        </tr>
                    </thead>
                    <tbody>
                        {sensorData.map((sensor) => (
                            <tr key={sensor.id} style={styles.tr}>
                                <td style={styles.td}>{sensor.id}</td>
                                <td style={styles.td}>{sensor.temperature}</td>
                                <td style={styles.td}>{sensor.humidity}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            ) : (
                <p style={styles.noData}>No data available</p>
            )}
        </div>
    );
};

// 🔹 Estilos en objeto JavaScript (Inline Styling)
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

// Efecto Hover para las filas
styles.tr[":hover"] = {
    backgroundColor: "#e0f7fa"
};

export default SensorDataDisplay;

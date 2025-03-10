// event-driven go/cmd/sensor-app/src/infrastructure/api/api.js

export class Api {
  async getSensorData() {
      const response = await fetch('http://api2:8000/send-data');
      const data = await response.json();
      console.log(data);
      return data.sensorData || []; // Si no hay datos, retorna un array vacío
  }
}

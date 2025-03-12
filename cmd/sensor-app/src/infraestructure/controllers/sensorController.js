// /infrastructure/controllers/sensorController.js

export class SensorController {
  constructor(webSocketAdapter, getSensorData) {
    this.webSocketAdapter = webSocketAdapter;
    this.getSensorData = getSensorData;

    this.webSocketAdapter.subscribe(this.handleNewSensorData.bind(this));
  }

  handleNewSensorData(data) {
    // Aquí puedes guardar los datos o pasarlos a un estado global
    this.getSensorData.store(data);
  }
}

// src/application/useCases/GetSensorDataUseCase.js

export class GetSensorDataUseCase {
  constructor(sensorService) {
      this.sensorService = sensorService;
  }

  async execute() {
      return await this.sensorService.getSensorData();
  }
}

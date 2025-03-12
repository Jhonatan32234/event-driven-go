// /application/usecases/getSensorData.js
export class GetSensorData {
    constructor(sensorRepository) {
      this.sensorRepository = sensorRepository;
    }
  
    async execute() {
      return await this.sensorRepository.getAll();
    }
  
    async store(data) {
      await this.sensorRepository.store(data);
    }
  }
  
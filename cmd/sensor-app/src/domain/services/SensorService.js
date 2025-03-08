// src/domain/services/SensorService.js

import { SensorData } from "../models/SensorData";

export class SensorService {
    constructor(api) {
        this.api = api;
    }

    async getSensorData() {
        const response = await this.api.getSensorData();
        return response.map(sensor => new SensorData(sensor.id, sensor.temperature, sensor.humidity));
    }
}

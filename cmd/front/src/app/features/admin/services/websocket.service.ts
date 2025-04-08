import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {
  private socket!: WebSocket;
  private messageSubject = new Subject<string>();

  constructor() {
    this.connect();
  }

  private connect() {
    this.socket = new WebSocket('ws://localhost:8081/ws?topic=nfc');

    this.socket.onopen = () => {
      console.log('Conectado al WebSocket');
    };

    this.socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        const content = JSON.parse(data.content);
        const nfcUid = content.uid;
        this.messageSubject.next(nfcUid);
      } catch (error) {
        console.error('Error al procesar el mensaje:', error);
      }
    };

    this.socket.onerror = (error) => {
      console.error('Error en WebSocket:', error);
    };

    this.socket.onclose = () => {
      console.log('Conexión WebSocket cerrada');
    };
  }

  getMessages(): Observable<string> {
    return this.messageSubject.asObservable();
  }
}

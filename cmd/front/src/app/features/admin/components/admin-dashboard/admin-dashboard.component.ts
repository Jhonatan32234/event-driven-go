import { Component, OnInit, Inject } from '@angular/core';
import { Messaging, getToken, onMessage } from '@angular/fire/messaging';
import { HttpClient } from '@angular/common/http';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-admin-dashboard',
  templateUrl: './admin-dashboard.component.html',
  styleUrls: ['./admin-dashboard.component.css'],
  standalone: false,
})
export class AdminDashboardComponent implements OnInit {
  token: string = '';
  private backendUrl = 'http://localhost:8000';
  private wsUrl = 'ws://localhost:5000/ws'; // URL para la conexión WebSocket
  private socket: WebSocket | null = null;

  // Arreglo para almacenar los mensajes
  public messages: any[] = [];

  constructor(
    @Inject(Messaging) private messaging: Messaging,
    private http: HttpClient
  ) {}

  ngOnInit() {
    //this.showNotificationButton();
    //this.listenForMessages();
    this.listenToSocket(); // Llamada para escuchar el WebSocket
  }

  showNotificationButton() {
    Notification.requestPermission()
      .then(async (permission) => {
        if (permission === 'granted') {
          try {
            const token = await getToken(this.messaging, {
              vapidKey:
                'BE_jQIwsH6tcbrpUexwsWDYfJSknW_S5_7ryOExehA0ddeKw2DAKsmr6mCGl6iZwf8X11X6IiH9jmHh6LwqWHZM',
            });
            if (token) {
              console.log('Token recibido:', token);
              this.token = token;
              this.subscribeToBackend(token);
            }
          } catch (err) {
            console.error('Error obteniendo token de FCM:', err);
          }
        }
      })
      .catch((error) => console.error('Error solicitando permisos:', error));
  }

  private subscribeToBackend(token: string) {
    this.http.post<{ message: string }>(`${this.backendUrl}/subscribe`, { token })
      .subscribe({
        next: (res) => console.log(res.message),
        error: (err) => console.error('🚨 Error en la suscripción:', err),
      });
  }

  listenForMessages() {
    onMessage(this.messaging, (payload) => {
      console.log('Mensaje recibido en primer plano:', payload);
  
      const data = payload.notification?.body;
      if (data) {
        try {
          // Verificar si 'data' ya es un JSON válido
          let parsedData;
          if (typeof data === 'string' && data.startsWith('{')) {
            parsedData = JSON.parse(data); // Intentar parsear JSON
          } else {
            console.warn('El mensaje recibido no está en formato JSON:', data);
            parsedData = this.convertToValidJson(data); // Intentar convertirlo
          }
  
          // Mostrar la notificación con SweetAlert
          Swal.fire({
            title: parsedData.title || 'Nueva Notificación',
            html: `
              <b>ID:</b> ${parsedData.id} <br>
              <b>Descripción:</b> ${parsedData.descripcion} <br>
              <b>Emisor:</b> ${parsedData.emmiter} <br>
              <b>Tema:</b> ${parsedData.topic} <br>
              <b>Fecha:</b> ${parsedData.created_at} <br>
            `,
            icon: 'info',
            confirmButtonText: 'Aceptar'
          });
  
          this.messages.push(parsedData);
        } catch (error) {
          console.error('❌ Error procesando el mensaje:', error, 'Mensaje:', data);
        }
      }
    });
  }
  
  /**
   * Intenta convertir un string con formato incorrecto en un objeto JSON válido
   */
  convertToValidJson(data: string) {
    const keyValuePairs = data.split(',').map(pair => pair.split(':').map(s => s.trim()));
    const jsonObject: any = {};
    
    keyValuePairs.forEach(([key, value]) => {
      if (value.startsWith('"') || value.startsWith("'")) {
        jsonObject[key] = value.replace(/['"]/g, '');
      } else if (!isNaN(Number(value))) {
        jsonObject[key] = Number(value);
      } else {
        jsonObject[key] = value;
      }
    });
  
    return jsonObject;
  }

  listenToSocket() {
    this.socket = new WebSocket(this.wsUrl);
  
    this.socket.onopen = () => {
      console.log("Conexión WebSocket establecida");
    };
  
    this.socket.onmessage = (event) => {
      try {
        let message: any;
        console.log('Mensaje recibido desde WebSocket:', event);
        
        // Verificar si el mensaje empieza con un carácter válido de JSON
        if (event.data && event.data.startsWith('{')) {
          message = JSON.parse(event.data); // Intentar parsear el JSON
          console.log('Mensaje recibido desde WebSocket:', message);
          
          // Mostrar la notificación con SweetAlert
          Swal.fire({
            title: 'Nuevo mensaje WebSocket',
            html: `
              <b>ID:</b> ${message.id} <br>
              <b>Descripción:</b> ${message.descripcion} <br>
              <b>Emisor:</b> ${message.emmiter} <br>
              <b>Tema:</b> ${message.topic} <br>
              <b>Fecha:</b> ${message.created_at} <br>
            `,
            icon: 'info',
            confirmButtonText: 'Aceptar'
          });
  
          // Guardar el mensaje WebSocket en el arreglo
          this.messages.push(message);
        } else {
          console.error('Mensaje no es un JSON válido:', event.data);
        }
      } catch (error) {
        console.error('Error procesando el mensaje WebSocket:', error, 'Mensaje recibido:', event.data);
      }
    };
  
    this.socket.onerror = (error) => {
      console.error('Error en la conexión WebSocket:', error);
    };
  
    this.socket.onclose = () => {
      console.log('Conexión WebSocket cerrada');
    };
  }
  
}

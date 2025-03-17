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

  constructor(
    @Inject(Messaging) private messaging: Messaging,
    private http: HttpClient
  ) {}

  ngOnInit() {
    this.listenForMessages();
    this.showNotificationButton();
  }

  showNotificationButton() {
    Swal.fire({
      title: '¿Deseas activar las notificaciones?',
      text: 'Haz clic en el botón para habilitar las notificaciones.',
      icon: 'info',
      showCancelButton: true,
      confirmButtonText: 'Activar Notificaciones',
      cancelButtonText: 'Cancelar',
    }).then((result) => {
      if (result.isConfirmed) {
        this.requestPermission(); // Llamar a la función para activar notificaciones
      }
    });
  }

  requestPermission() {
    Notification.requestPermission()
      .then(async (permission) => {
        if (permission === 'granted') {
          try {
            const token = await getToken(this.messaging, {
              vapidKey:
                'BE_jQIwsH6tcbrpUexwsWDYfJSknW_S5_7ryOExehA0ddeKw2DAKsmr6mCGl6iZwf8X11X6IiH9jmHh6LwqWHZM',
            });
//'BNiXbBcCoErAiquuylp5PsU2nT8I1Tj4fbX-JPzEj1nyb7A3lQuNxKdZuSy-J4W9QkhPFjT05SQC5s1cv64GlB8'
            if (token) {
              console.log('Token recibido:', token);
              this.token = token;

              this.subscribeToBackend(token);
            }
          } catch (err) {
            console.error('Error obteniendo token de FCM:', err);
          }
        } else {
          console.warn('Permiso de notificaciones no concedido');
        }
      })
      .catch((error) => console.error('Error solicitando permisos:', error));
  }

  private subscribeToBackend(token: string) {
    this.http.post<{ message: string }>(`${this.backendUrl}/api/subscribe`, { token })
      .subscribe({
        next: (res) => console.log(res.message),
        error: (err) => console.error('🚨 Error en la suscripción:', err),
      });
  }

  listenForMessages() {
    onMessage(this.messaging, (payload) => {
      console.log('Mensaje recibido en primer plano:', payload);
  
      Swal.fire({
        title: payload.notification?.title || 'Nueva Temperatura', // Título de la notificación
        text: payload.notification?.body || 'Tienes un nuevo valor de temperatura.', // Cuerpo de la notificación
        icon: 'info', // Icono (puedes cambiarlo a 'success', 'warning', 'error', etc.)
        showConfirmButton: true, // Mostrar botón de confirmación
        confirmButtonText: 'Aceptar', // Texto del botón
        confirmButtonColor: '#3085d6', // Color del botón
        customClass: {
          popup: 'custom-popup', // Clase CSS personalizada para la ventana emergente
          title: 'custom-title', // Clase CSS personalizada para el título
          // content: 'custom-content', // Clase CSS personalizada para el contenido
          confirmButton: 'custom-confirm-button', // Clase CSS personalizada para el botón
        },
        // Animación de entrada
        showClass: {
          popup: 'animate__animated animate__fadeInDown', // Animación de entrada
        },
        // Animación de salida
        hideClass: {
          popup: 'animate__animated animate__fadeOutUp', // Animación de salida
        },
      });
    });
  }
}

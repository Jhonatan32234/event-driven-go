import { Injectable } from '@angular/core';
import { AngularFireMessaging } from '@angular/fire/compat/messaging';
import { take } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class MessagingService {
  constructor(private afMessaging: AngularFireMessaging) {}

  /*requestPermission() {
    this.afMessaging.requestToken
      .pipe(take(1))
      .subscribe(
        (token) => {
          console.log('Token FCM:', token);
        },
        (error) => {
          console.error('Error al obtener el token FCM', error);
        }
      );
  }*/
}

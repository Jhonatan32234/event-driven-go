import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AdminModule } from './features/admin/admin.module';
import { HttpClientModule } from '@angular/common/http';

// Firebase
import { initializeApp, provideFirebaseApp } from '@angular/fire/app';
import { getMessaging, provideMessaging } from '@angular/fire/messaging';

import { environment } from '../environments/environment';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    AdminModule,
    HttpClientModule
  ],
  providers: [
    provideFirebaseApp(() => initializeApp(environment.firebase)),
    provideMessaging(() => getMessaging())
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

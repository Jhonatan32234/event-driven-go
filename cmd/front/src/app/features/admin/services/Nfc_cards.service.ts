import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { NfcCard } from '../models/iNfc_cards';

@Injectable({
  providedIn: 'root'
})
export class NfcCardsService {
  private apiUrl = 'http://localhost:8080/nfc_cards';

  constructor(private http: HttpClient) {}

  createNfcCard(data: NfcCard): Observable<NfcCard> {
    return this.http.post<NfcCard>(`${this.apiUrl}`, data);
  }

  getNfcCardByUID(uid: string): Observable<NfcCard> {
    return this.http.get<NfcCard>(`${this.apiUrl}/${uid}`);
  }

  deleteNfcCard(uid: string): Observable<NfcCard> {
    return this.http.delete<NfcCard>(`${this.apiUrl}/${uid}`);
  }

  getAllNfcCards(): Observable<NfcCard[]> {
    return this.http.get<NfcCard[]>(`${this.apiUrl}`);
  }
}
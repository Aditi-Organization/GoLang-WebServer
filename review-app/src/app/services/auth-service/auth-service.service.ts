import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AuthServiceService {

  apiUrl = 'http://localhost:8080/';
  private userLoggedIn = new Subject<boolean>();

  constructor(private http: HttpClient) {
    this.userLoggedIn.next(false);
  }

  setUserLoggedIn(userLoggedIn: boolean) {
    this.userLoggedIn.next(userLoggedIn);
  }

  getUserLoggedIn(): Observable<boolean> {
    return this.userLoggedIn.asObservable();
  }

  authenticateUser(user: Object): Observable<any> {
    return this.http.post(this.apiUrl + 'api/auth/login', user, {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      })
    });
  }

  registerUser(user: Object): Observable<any> {
    return this.http.post(this.apiUrl + 'api/auth/signup', user, {
      headers: new HttpHeaders({
        "Content-Type": "application/json"
      })
    });
  }


}

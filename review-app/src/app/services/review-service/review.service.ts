import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ReviewService {

  apiUrl = 'http://localhost:8080/';

  constructor(private http: HttpClient) { }

  getReviewsForMovie(id:String): Observable<any> {
    return this.http.get(this.apiUrl + 'api/review/'+id, {
    });
  }
}

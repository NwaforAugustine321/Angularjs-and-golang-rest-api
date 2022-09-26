import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit, Output } from '@angular/core';
import { catchError, throwError } from 'rxjs';
import { Imovie } from './interface';

@Component({
  selector: 'app-movies',
  templateUrl: './movies.component.html',
  styleUrls: ['./movies.component.css'],
})
export class MoviesComponent implements OnInit {
  @Output() movies_list: Imovie[] | undefined = [];
  constructor(private http: HttpClient) {}

  private handleError(error: HttpErrorResponse) {
    return throwError(() => new Error(error.message));
  }

  ngOnInit(): void {
    this.http
      .get<[]>('http://localhost:4000/v1/movies')
      .pipe(catchError(this.handleError))
      .subscribe((data: []): void => {
        this.movies_list = data;
      });
  }
}

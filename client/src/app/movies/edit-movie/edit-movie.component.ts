import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Component, ElementRef, Input, OnInit, Output } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { ActivatedRoute, Params } from '@angular/router';
import { catchError, throwError } from 'rxjs';

@Component({
  selector: 'edit-movie',
  templateUrl: './edit-movie.component.html',
  styleUrls: ['./edit-movie.component.css'],
})
export class EditMovieComponent implements OnInit {
  movie_form: FormGroup;

  constructor(private http: HttpClient, private router: ActivatedRoute) {
    this.movie_form = new FormGroup({
      title: new FormControl(),
      description: new FormControl(),
      year: new FormControl(),
      runTime: new FormControl(),
      rating: new FormControl(),
    });
  }

  private handleError(error: HttpErrorResponse) {
    return throwError(() => new Error(error.message));
  }

  ngOnInit(): void {}

  onSubmit(): void {
    this.router.queryParams.subscribe((param: Params) => {
      if (param['action'] === 'edit') {
        this.http
          .post(
            `http://localhost:4000/v1/movie/edit/${Number(param['id'])}`,
            this.movie_form.value
          )
          .pipe(catchError(this.handleError))
          .subscribe((res) => {
            console.log(res);
          });
      } else if (param['action'] === 'create') {
        this.http
          .post(`http://localhost:4000/v1/movie/create`, this.movie_form.value)
          .pipe(catchError(this.handleError))
          .subscribe((res) => {
            console.log(res);
          });
      } else {
        alert('Invalid url');
      }
    });
  }
}

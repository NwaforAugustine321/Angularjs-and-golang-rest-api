import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Imovie } from '../interface';

@Component({
  selector: 'app-movie-view',
  templateUrl: './movie-view.component.html',
  styleUrls: ['./movie-view.component.css'],
})
export class MovieViewComponent implements OnInit {
  movie_detail: {
    title: string;
    description: string;
    runTime: number;
    rating: number;
    mpaa_rating: string;
    id?: number | null;
  } = {
    title: '',
    description: '',
    runTime: 0,
    rating: 0,
    mpaa_rating: '',
  };
  constructor(private http: HttpClient, private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.route.params.subscribe((param): void => {
      this.http
        .get<Imovie>(`http://localhost:4000/v1/movie/${Number(param['id'])}`)
        .subscribe((movie_detail: Imovie) => {
          this.movie_detail = movie_detail;
        });
    });
  }
}

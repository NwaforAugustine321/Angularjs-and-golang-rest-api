import { Component, Input } from '@angular/core';
import { Imovie } from '../interface';

@Component({
  selector: 'movie-list',
  templateUrl: './movie-list.component.html',
  styleUrls: ['./movie-list.component.css'],
})
export class MovieListComponent {
  @Input('movie') movie!: Imovie | undefined;

  constructor() {}
}

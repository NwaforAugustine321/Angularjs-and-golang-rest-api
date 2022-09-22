import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AdminComponent } from './admin/admin.component';
import { HomeScreenComponent } from './home-screen/home-screen.component';
import { MainComponent } from './main/main.component';
import { EditMovieComponent } from './movies/edit-movie/edit-movie.component';
import { MovieViewComponent } from './movies/movie-view/movie-view.component';
import { MoviesComponent } from './movies/movies.component';
import { NotFoundComponent } from './not-found/not-found.component';

const routes: Routes = [
  {
    path: '',
    component: MainComponent,
    children: [
      {
        path: 'admin',
        component: AdminComponent,
      },
      {
        path: 'movies',
        component: MoviesComponent,
      },
      {
        path: '',
        component: HomeScreenComponent,
      },
      {
        path: 'movie/:id',
        component: MovieViewComponent,
      },

      {
        path: 'movie',
        component: EditMovieComponent,
      },
    ],
  },

  {
    path: '**',
    redirectTo: '',
    component: NotFoundComponent,
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}

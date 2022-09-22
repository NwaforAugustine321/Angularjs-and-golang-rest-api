import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeScreenComponent } from './home-screen/home-screen.component';
import { AdminComponent } from './admin/admin.component';
import { MoviesComponent } from './movies/movies.component';
import { NotFoundComponent } from './not-found/not-found.component';
import { MainComponent } from './main/main.component';
import { MovieListComponent } from './movies/movie-list/movie-list.component';
import { MovieViewComponent } from './movies/movie-view/movie-view.component';
import { EditMovieComponent } from './movies/edit-movie/edit-movie.component';
import { ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    HomeScreenComponent,
    AdminComponent,
    MoviesComponent,
    NotFoundComponent,
    MainComponent,
    MovieListComponent,
    MovieViewComponent,
    EditMovieComponent,
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    ReactiveFormsModule,
  ],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}

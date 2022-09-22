export interface Imovie {
  id: number;
  title: string;
  description: string;
  year: number;
  releaseDate: Date;
  runTime: number;
  rating: number;
  mpaa_rating: string;
  createdAt: Date;
  updatedAt: Date;
  moviesGenre: { [key: string]: string };
}

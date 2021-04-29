import { Component, Input, OnInit } from '@angular/core';
import { MovieObject } from '../movieobject.model';
import { MovieService } from '../services/movie-service/movie.service';

@Component({
  selector: 'app-movie-list',
  templateUrl: './movie-list.component.html',
  styleUrls: ['./movie-list.component.scss']
})

export class MovieListComponent implements OnInit {

  public movies: MovieObject[];
  constructor(public movieService: MovieService) {
  }

  ngOnInit(): void {
    this.getMovieList();
  }
  getMovieList() {
    this.movieService.getAllMovies().subscribe((movieList: MovieObject[]) => {
      this.movies = movieList;
    })
  }

  callMovie() {
    console.log("HERE");
  }
}



import { Component, Input, OnInit } from '@angular/core';
import { MovieObject } from '../movieobject.model';
import { MovieSearchService } from '../services/movie-search-service/movie-search.service';
import { MovieService } from '../services/movie-service/movie.service';

@Component({
  selector: 'app-movie-list',
  templateUrl: './movie-list.component.html',
  styleUrls: ['./movie-list.component.scss']
})

export class MovieListComponent implements OnInit {

  public movies: MovieObject[];
  public textValue: any;

  constructor(public movieService: MovieService, public searchValue: MovieSearchService) {
  }

  ngOnInit(): void {
    this.getMovieList();
  }
  getMovieList() {
    this.movieService.getAllMovies().subscribe((movieList: MovieObject[]) => {
      this.movies = movieList;
    })
  }
  public submitOnClick() {
    this.textValue = (<HTMLInputElement>document.getElementById('searchText')).value;
    this.searchValue.data = this.textValue;
    console.log("searchValue: " + this.searchValue.data);
  }

  callMovie() {
    console.log("HERE");
  }
}



import { Component, OnInit } from '@angular/core';
import { MovieService } from '../services/movie-service/movie.service';
import { SearchObject } from '../searchobject.model';
import { HomepageComponent } from '../homepage/homepage.component';
import { MovieSearchService } from '../services/movie-search-service/movie-search.service';

@Component({
  selector: 'app-movie-search',
  templateUrl: './movie-search.component.html',
  styleUrls: ['./movie-search.component.scss']
})
export class MovieSearchComponent implements OnInit {

  public movie: SearchObject[];
  public movieName: any;
  public textValue: any;

  constructor(public movieService: MovieService, public homepageComponent: HomepageComponent, public searchValue: MovieSearchService) {
    this.movieName = homepageComponent.searchValue.data;
    console.log("Movie Name Here:" + this.movieName);
  }

  ngOnInit(): void {
    this.getMovieList();
  }

  getMovieList() {
    this.movieService.getMovie(this.movieName).subscribe((movie: SearchObject[]) => {
      this.movie = movie;
    })
  }
  public submitOnClick() {
    this.textValue = (<HTMLInputElement>document.getElementById('searchText')).value;
    this.searchValue.data = this.textValue;
    console.log("searchValue: " + this.searchValue.data);
  }
}

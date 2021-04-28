import { Component, OnInit } from '@angular/core';
import { HeaderComponentService } from '../services/header-component-service/header-component.service';
import { HeaderComponent } from '../header/header.component';
import { AuthServiceService } from '../services/auth-service/auth-service.service';
import { MovieObject } from '../movieobject.model';
import { MovieService } from '../services/movie-service/movie.service';

@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.scss']
})
export class HomepageComponent implements OnInit {

  public elements: any[];
  public movies: MovieObject[];

  constructor(public headerComponentService: HeaderComponentService, public headerComponent: HeaderComponent
    , public authServiceService: AuthServiceService, public movieService : MovieService) {
    this.headerComponent.elements[0].style.visibility = "visible";
    this.headerComponent.elements[1].style.visibility = "hidden";
    if (authServiceService.getSessionData() != null) {
      this.headerComponent.elements[0].style.visibility = "hidden";
      this.headerComponent.elements[1].style.visibility = "visible";
    }
    this.movies=[];
  }

  ngOnInit(): void {
    this.movieService.getAllMovies().subscribe((movieList : MovieObject[]) =>{
      this.movies = movieList;
    })

  }

  homeSlider = { items: 1, dots: true, nav: true };

}

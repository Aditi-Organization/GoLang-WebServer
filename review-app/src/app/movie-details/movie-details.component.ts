import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { MovieObject } from '../movieobject.model';
import { MovieService } from '../services/movie-service/movie.service';
import { ReviewObject } from '../reviewobject.model';
import { ReviewService } from '../services/review-service/review.service';

@Component({
  selector: 'app-movie-details',
  templateUrl: './movie-details.component.html',
  styleUrls: ['./movie-details.component.scss']
})
export class MovieDetailsComponent implements OnInit {

  public movieId : String;

  public movieDetails: MovieObject;

  public reviews : ReviewObject[];

  // const fieldset = document.fo
  currentRate = 0;

  constructor(public router : Router, public route : ActivatedRoute, public movieService : MovieService
    , public reviewService : ReviewService) {
    route.url.subscribe(() => {
      this.movieId = route.snapshot.firstChild.url[0].path;
      console.log(this.movieId);
     });
  }

  ngOnInit(): void {

    this.movieService.getMovieDetail(this.movieId).subscribe((movieDet : MovieObject) =>{
      this.movieDetails = movieDet;
      console.log(this.movieDetails);
    });

    this.reviewService.getReviewsForMovie(this.movieId).subscribe((movieReviews : ReviewObject[]) =>{
      this.reviews = movieReviews;
      console.log(this.reviews);
    });

  }

  onSubmit(){
    console.log(this.currentRate);
  }
}

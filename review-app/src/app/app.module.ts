import { BrowserModule } from '@angular/platform-browser';
import { CUSTOM_ELEMENTS_SCHEMA, NgModule } from '@angular/core';


import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { HomepageComponent } from './homepage/homepage.component';
import { LoginComponent } from './login/login.component';
import { ReactiveFormsModule } from '@angular/forms';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { SignupComponent } from './signup/signup.component';
import { HttpClientModule } from '@angular/common/http';
import { OwlModule } from 'ngx-owl-carousel'
import { HeaderComponentService } from './services/header-component-service/header-component.service';
import { AuthServiceService } from './services/auth-service/auth-service.service';
import { MovieDetailsComponent } from './movie-details/movie-details.component';
import { MovieService } from './services/movie-service/movie.service';


@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    FooterComponent,
    HomepageComponent,
    LoginComponent,
    SignupComponent,
    MovieDetailsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    NgbModule,
    HttpClientModule,
    OwlModule
  ],
  schemas:[ CUSTOM_ELEMENTS_SCHEMA],
  providers: [HeaderComponentService, AuthServiceService, MovieService],
  bootstrap: [AppComponent]
})
export class AppModule { }

import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.scss']
})
export class SignupComponent implements OnInit {

  form: FormGroup;
  loading = false;
  submitted = false;
  token: any;
  refreshToken: any;
  id: any

  constructor(private formBuilder: FormBuilder,
    private route: ActivatedRoute,
    private router: Router
    // private authService: AuthService
    ) { }

  ngOnInit(): void {
    this.form = this.formBuilder.group({
      firstName: ['', Validators.required],
      lastName: ['', Validators.required],
      username: ['', Validators.required],
      password: ['', [Validators.required, Validators.minLength(6)]]
  });
  }
  get f() { return this.form.controls; }

    onSubmit() {
        this.submitted = true;

        // // reset alerts on submit
        // this.alertService.clear();

        // // stop here if form is invalid
        if (this.form.invalid) {
            return;
        }

        this.loading = true;

    }

}

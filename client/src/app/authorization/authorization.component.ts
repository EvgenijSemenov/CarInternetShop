import { Component, OnInit } from '@angular/core';
import {AppComponent} from '../app.component';
import {User} from "../entity/user";
import {Router} from "@angular/router";

@Component({
  selector: 'app-authorization',
  templateUrl: './authorization.component.html',
  styleUrls: ['./authorization.component.css']
})
export class AuthorizationComponent implements OnInit {

  private email: string = "";
  private password: string = "";

  constructor(private appComponent: AppComponent, private router: Router) { }

  ngOnInit() {
  }

  login() {
    this.appComponent.authorization.authorizateUser();
    this.router.navigate(['']);
  }

}

import { Component, OnInit } from '@angular/core';
import {Router} from "@angular/router";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {

  public isAuthUser: boolean = false;

  constructor(private router: Router) {}

  ngOnInit() {
    let authToken = localStorage.getItem('auth-token');
    console.log(this.router.url);
  }
}

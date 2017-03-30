import { Component, OnInit } from '@angular/core';
import {Md5} from 'ts-md5/dist/md5';
import {User} from "../entity/user";
import {UserService} from "../service/user.service";
import {Router} from "@angular/router";

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit {

  private user: User = new User(null, "", "", "", "");
  private password: string = "";
  private errorMessage: any;

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit() {
  }

  registrateNewUser(){
    this.user.pass_hash = Md5.hashStr(this.user.email + ":" + this.password).toString();
    this.userService.addUser(this.user).subscribe(
      user => this.router.navigate(['']),
      error =>  this.errorMessage = <any>error
    );
  }

}

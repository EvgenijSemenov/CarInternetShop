import { Component, OnInit } from '@angular/core';
import {CarService} from "../service/car.service";
import {Car} from "../entity/car";
import {Router} from "@angular/router";

@Component({
  selector: 'app-add-advertisement',
  templateUrl: './add-advertisement.component.html'
})

export class AddGoodComponent implements OnInit {

  private car: Car = new Car(null);

  constructor(private carService: CarService, private router: Router) { }

  ngOnInit() {}

  addCar() {
    this.carService.addCar(this.car).subscribe(
      resp => this.router.navigate(['']),
      error =>  console.log(error)
    );
  }

}

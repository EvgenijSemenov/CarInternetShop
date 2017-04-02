import { Component, OnInit } from '@angular/core';
import { CarService } from '../service/car.service';
import {Car} from "../entity/car";
import {Authorization} from "../security/authorization";
import {AppComponent} from "../app.component";

@Component({
  selector: 'app-good-list',
  templateUrl: './good-list.component.html',
  styleUrls: ['./good-list.component.css']
})

export class GoodListComponent implements OnInit {

  public cars: Car[];
  private errorMessage: any;
  private authorization: Authorization = new Authorization();

  constructor(private carService: CarService, private appComponent: AppComponent) {
    this.authorization = appComponent.authorization;
  }

  ngOnInit() {
    this.carService.getAllCars().subscribe(
        cars => this.cars = cars,
        error =>  this.errorMessage = <any> error
    );
  }

  deleteGood(id: number) {
    this.carService.deleteCar(id).subscribe(
      body => this.ngOnInit(),
      error =>  this.errorMessage = <any> error
    );
  }

}

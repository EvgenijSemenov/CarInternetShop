import { Component, OnInit } from '@angular/core';
import { CarService } from '../service/car.service';
import {Car} from "../entity/car";

@Component({
  selector: 'app-good-list',
  templateUrl: './good-list.component.html',
  styleUrls: ['./good-list.component.css']
})

export class GoodListComponent implements OnInit {

  public cars: Car[];
  private errorMessage: any;

  constructor(private carService: CarService) { }

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

import {Component, OnInit} from '@angular/core';
import {Router, ActivatedRoute} from "@angular/router";
import {Car} from "../entity/car";
import {CarService} from "../service/car.service";

@Component({
  selector: 'app-edit-good',
  templateUrl: './edit-good.component.html'
})

export class EditGoodComponent implements OnInit {

  private car = new Car(null);
  private errorMessage: any;

  constructor(private carService: CarService, private route: ActivatedRoute, private router: Router) {}

  ngOnInit() {
    let routeId = this.route.snapshot.params['id'];

    this.carService.getCar(routeId).subscribe(
      car => this.car = car[0],
      error =>  this.errorMessage = <any>error
    );
  }

  updateCar() {
    this.carService.updateCar(this.car).subscribe(
       car => this.router.navigate(['']),
       error =>  this.errorMessage = <any>error
     );
  }

}

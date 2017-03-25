import { Injectable } from '@angular/core';
import {Http, Response, Headers} from "@angular/http";
import {Car} from "../entity/car";
import {Observable} from "rxjs";

@Injectable()
export class CarService {

  constructor(private http: Http) { }

  private carsUri = 'api/v1/car';

  getCar(id: number): Observable<Car> {
    return this.http.get("http://localhost:8080/" + this.carsUri + "/" + id)
      .map(this.extractData)
      .catch(this.handleError);
  }

  getAllCars(): Observable<Car[]> {
    return this.http.get("http://localhost:8080/" + this.carsUri + "/all")
      .map(this.extractData)
      .catch(this.handleError);
  }

  addCar(car: Car): Observable<Car> {
    let body = JSON.stringify(car);
    let headers = new Headers({ 'Content-Type': 'application/json' });

    return this.http.post("http://localhost:8080/" + this.carsUri + "/add", body, headers)
      .map(this.extractData)
      .catch(this.handleError);
  }

  updateCar(car: Car): Observable<Car> {
    let body = JSON.stringify(car);
    let headers = new Headers({ 'Content-Type': 'application/json' });

    return this.http.post('http://localhost:8080/' + this.carsUri+ "/update", body, headers)
      .map(this.extractData)
      .catch(this.handleError);
  }

  deleteCar(id: number): Observable<Car> {
    let car: Car = new Car(id);

    return this.http.post("http://localhost:8080/" + this.carsUri + "/delete", JSON.stringify(car))
      .map(this.extractData)
      .catch(this.handleError);
  }

  private extractData(res: Response) {
    return res.json();
  }

  private handleError (error: Response | any) {
    let errMsg: string;
    if (error instanceof Response) {
      const body = error.json() || '';
      const err = body.error || JSON.stringify(body);
      errMsg = `${error.status} - ${error.statusText || ''} ${err}`;
    } else {
      errMsg = error.message ? error.message : error.toString();
    }
    console.error(errMsg);
    return Observable.throw(errMsg);
  }

}

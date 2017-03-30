import { Injectable } from '@angular/core';
import {Http, Response, Headers} from "@angular/http";
import {Observable} from "rxjs";
import {User} from "../entity/user";

@Injectable()
export class UserService {

  constructor(private http: Http) { }

  private userApiUri = 'api/v1/user';

  addUser(user: User): Observable<User> {
    let body = JSON.stringify(user);
    let headers = new Headers({ 'Content-Type': 'application/json' });

    return this.http.post("http://localhost:8080/" + this.userApiUri + "/add", body, headers)
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

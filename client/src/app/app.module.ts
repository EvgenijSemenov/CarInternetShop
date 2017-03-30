import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import {HttpModule} from '@angular/http';

import { AppComponent } from './app.component';
import { GoodListComponent } from './good-list/good-list.component';
import {CarService} from "./service/car.service";
import { AddGoodComponent } from './add-advertisement/add-advertisement.component';
import { RouterModule }   from '@angular/router';
import { EditGoodComponent } from './edit-good/edit-good.component';
import { AuthorizationComponent } from './authorization/authorization.component';
import { RegistrationComponent } from './registration/registration.component';
import {UserService} from "./service/user.service";

@NgModule({
  declarations: [
    AppComponent,
    GoodListComponent,
    AddGoodComponent,
    EditGoodComponent,
    AuthorizationComponent,
    RegistrationComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    RouterModule.forRoot([
      {
        path: '',
        component: GoodListComponent
      },
      {
        path: 'add-good',
        component: AddGoodComponent
      },
      {
        path: 'edit-good/:id',
        component: EditGoodComponent
      },
      {
        path: 'authorization',
        component: AuthorizationComponent
      },
      {
        path: 'registration',
        component: RegistrationComponent
      }
    ])
  ],
  providers: [CarService, UserService],
  bootstrap: [AppComponent],

})

export class AppModule { }

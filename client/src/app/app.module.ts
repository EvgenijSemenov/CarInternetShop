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

@NgModule({
  declarations: [
    AppComponent,
    GoodListComponent,
    AddGoodComponent,
    EditGoodComponent
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
      }
    ])
  ],
  providers: [CarService],
  bootstrap: [AppComponent],

})

export class AppModule { }

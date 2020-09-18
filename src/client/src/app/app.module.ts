import { NgModule } from "@angular/core";
import { RouterModule } from "@angular/router";
import { ReactiveFormsModule } from "@angular/forms";
import { MatInputModule } from "@angular/material/input";
import { MatSelectModule } from "@angular/material/select";
import { BrowserModule } from "@angular/platform-browser";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";

import { AppComponent } from "./app.component";
import { TopBarComponent } from "./top-bar/top-bar.component";
import { CityCollectionComponent } from "./city-collection/city-collection.component";

@NgModule({
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    RouterModule.forRoot([{ path: "", component: CityCollectionComponent }]),
    MatSelectModule,
    MatInputModule,
  ],
  declarations: [AppComponent, TopBarComponent, CityCollectionComponent],
  bootstrap: [AppComponent],
})
export class AppModule {}

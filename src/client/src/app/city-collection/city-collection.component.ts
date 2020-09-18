import { Component } from "@angular/core";

import {
  CitiesCollection,
  CitiesCollectionList,
  Temperatures,
  CityWithTemp,
} from "src/entities/CityCollection";
import { url } from "src/utils/Url";

// TODO: needs serious refactoring

// type for possibly undefined to null values
type Maybe<T> = T | null | undefined;

/**
 * The cities collection component encapsulates a list of cities along with
 * methods to manipulate these collections on the server backend.
 */
@Component({
  selector: "city-collection",
  templateUrl: "./city-collection.component.html",
  styleUrls: ["./city-collection.component.css"],
})
export class CityCollectionComponent {
  /** the selected collection */
  public selected: CitiesCollection;

  /** the selected temperatures */
  public temperatures: CityWithTemp[];

  /** a list of collections */
  public collections: CitiesCollection[];

  public newCityName: string;
  public newCollectionName: string;

  constructor() {
    this.newCityName = "";
    this.collections = [];
    this.getCollectionList();
  }

  /** performs a request to the backend to retrieve all collections */
  public async getCollectionList(): Promise<void> {
    const fetched = await this.fetchCitiesCollectionList();

    if (fetched) {
      this.selected = fetched.List[0];
      this.collections = fetched.List;
    }
  }

  /** update the new city name */
  public updateNewCity(evt) {
    this.newCityName = evt.target.value;
  }

  /** update the new city name */
  public updateCollection(evt) {
    this.newCollectionName = evt.target.value;
  }

  /** update the new city name */
  public async setTemperatures() {
    // the execution of the update of this list is not bound to a particular model
    // this it always lags one step behind when updating the 'selected' collection

    // this is not a nice way, but it works
    // TODO: rework
    setTimeout(async () => {
      this.temperatures = [];
      const fetchedTemps = await this.fetchTemps(this.selected.ID);
      this.temperatures = fetchedTemps.List.map((t: CityWithTemp) => ({
        Name: t.Name,
        Temp: `${t.Temp}`,
        OwmID: `${t.OwmID}`,
      }));
    }, 50);
  }

  /** remove a collection */
  public async removeCollection() {
    const citiesUrl = url.api(`city/${this.selected.ID}`);
    await fetch(citiesUrl, {
      method: "DELETE",
      cache: "no-cache",
      mode: "cors",
    });
  }

  /** to add a city to the selected collection */
  public async addCity(): Promise<void> {
    await this.postCity(this.newCityName);
    this.temperatures = [];
    this.getCollectionList();
  }

  public async addCollection(): Promise<void> {
    await this.putCollection(this.newCollectionName);
    this.temperatures = [];
    this.getCollectionList();
  }

  private async fetchCitiesCollectionList(): Promise<
    Maybe<CitiesCollectionList>
  > {
    const fetched = await fetch(url.api("cities/all"), {
      method: "GET",
      cache: "no-cache",
      mode: "cors",
    });

    const cities = await fetched.json();

    return cities;
  }

  private async putCollection(name: string): Promise<Maybe<CitiesCollection>> {
    const fetched = await fetch(url.api(`city/${name}`), {
      method: "PUT",
      cache: "no-cache",
      mode: "cors",
    });

    if (fetched.status === 200) {
      const cities = await fetched.json();
      return cities;
    }
  }

  private async postCity(name: string): Promise<Maybe<CitiesCollection>> {
    const fetched = await fetch(url.api(`city/${this.selected.ID}/${name}`), {
      method: "POST",
      cache: "no-cache",
      mode: "cors",
    });

    if (fetched.status === 200) {
      const cities = await fetched.json();
      return cities;
    }
  }

  /** fetches the temperature list belonging to a collection with 'id' */
  private async fetchTemps(id: number): Promise<Temperatures> {
    const fetched = await fetch(url.api(`temp/${id}`), {
      method: "GET",
      cache: "no-cache",
      mode: "cors",
    });

    if (fetched.status === 200) {
      const data = fetched.json();
      return data;
    }
  }
}

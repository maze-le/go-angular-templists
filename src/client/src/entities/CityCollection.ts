import { Model } from "./Model";
import { City } from './City';

export interface CitiesCollection extends Model {
  CityCollectionID: number;
  Name: string;
  Cities: City[];
}

export interface CityWithTemp {
  Name: string;
  Temp: string;
  OwmID: string;
}

export interface Temperatures {
  List: CityWithTemp[];
}

export interface CitiesCollectionList {
  List: CitiesCollection[];
}

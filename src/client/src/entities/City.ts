import { Model } from './Model';

export interface City extends Model {
  CityID: number;
  OwmID: number;
  CityCollectionID: number;
  Name: string;
  Cities: City[] | null;
}

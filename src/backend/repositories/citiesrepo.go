package repositories

/** The cities repository encapsulates database operations for city collections */

import (
	ent "cityserver/entities"
	mid "cityserver/middleware"
	serv "cityserver/services"
)

// bogo return values for errors
var dc []ent.City
var d = ent.CityCollection{
	Name:   "Error",
	Cities: dc,
}

// Create adds a new entity to the database and returns the created collection
func Create(name string) (ent.CityCollection, *mid.HTTPError) {
	var newEntity ent.CityCollection = ent.CityCollection{Name: name, Cities: []ent.City{}}

	dbConnection.Create(&newEntity)

	return newEntity, nil
}

// Read retrieves collections from the database
func Read(id string) (ent.CityCollection, *mid.HTTPError) {
	var collections []ent.CityCollection

	dbConnection.Preload("Cities").Find(&collections, "ID = ?", id)
	if len(collections) == 0 {
		return d, mid.Throw404("cannot find collection with id: " + id)
	}

	return collections[0], nil
}

// Index retrieves all collections from the database
func Index() ([]ent.CityCollection, *mid.HTTPError) {
	var collections []ent.CityCollection

	dbConnection.Preload("Cities").Find(&collections)

	return collections, nil
}

// Update replaces a collection entity on the database and returns the updated collection
func Update(id string, name string) (ent.CityCollection, *mid.HTTPError) {
	var collections []ent.CityCollection
	var selected ent.CityCollection

	dbConnection.Preload("Cities").Find(&collections, "ID = ?", id)
	if len(collections) == 0 {
		return d, mid.Throw404("cannot find collection with id: " + id)
	}

	var cityFromOWM, err = serv.QueryCityByName(name)
	if err != nil {
		return d, mid.Throw404("owm cannot find city: " + name)
	}

	selected = collections[0]

	mid.LogInfo("selected: " + selected.Name)

	var newCity = &ent.City{Name: cityFromOWM.Name, OwmID: cityFromOWM.OwmID}
	dbConnection.Create(&newCity)
	dbConnection.Model(&selected).Association("Cities").Append(newCity)

	return selected, nil
}

// Delete deletes an entity returns the deleted collection
func Delete(id string) ([]ent.CityCollection, *mid.HTTPError) {
	var collections []ent.CityCollection

	dbConnection.Preload("Cities").Find(&collections, "ID = ?", id)
	if len(collections) == 0 {
		return collections, mid.Throw404("cannot find collection with id: " + id)
	}

	dbConnection.Delete(&collections, id)

	return collections, nil
}

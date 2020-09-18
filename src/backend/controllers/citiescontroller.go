package controllers

/**
  The cities controller represents the '/cities/' API-Endpoint. It connects to the
  cities database repository and renders json responses on each HTTP call.
*/

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	ent "cityserver/entities"
	mid "cityserver/middleware"
	repo "cityserver/repositories"
)

// CollectionList encodes a list of temperatures
type CollectionList struct {
	List []ent.CityCollection
}

// CollectionIndex lists all possible collections
func CollectionIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mid.LogRequest(r)

	col, err := repo.Index()

	respondAll(w, col, err)
}

// CreateCollection creates a new city collection
func CreateCollection(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mid.LogRequest(r)

	var name = ps.ByName("name")
	col, err := repo.Create(name)

	respond(w, col, err)
}

// ReadCollection retrieves a city collection
func ReadCollection(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mid.LogRequest(r)

	var id = ps.ByName("id")
	col, err := repo.Read(id)

	respond(w, col, err)
}

// UpdateCollection updates an existing city collection
func UpdateCollection(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mid.LogRequest(r)

	var id = ps.ByName("id")
	var cityID = ps.ByName("cityId")
	col, err := repo.Update(id, cityID)

	respond(w, col, err)
}

// DeleteCollection deletes a city collection
func DeleteCollection(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mid.LogRequest(r)

	var id = ps.ByName("id")
	col, err := repo.Delete(id)

	respondAll(w, col, err)
}

// small helper to declutter the controller methods and unify http json responses
func respond(w http.ResponseWriter, collection ent.CityCollection, err *mid.HTTPError) {
	if err == nil {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(collection)
	} else {
		mid.HandleError(w, err)
	}
}

// small helper to declutter the controller methods and unify http json responses
func respondAll(w http.ResponseWriter, collections []ent.CityCollection, err *mid.HTTPError) {
	if err == nil {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", "*")
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Content-Type", "application/json")

		var clist = &CollectionList{List: collections}

		json.NewEncoder(w).Encode(&clist)
	} else {
		mid.HandleError(w, err)
	}
}

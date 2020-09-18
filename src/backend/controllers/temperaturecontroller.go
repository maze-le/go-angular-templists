package controllers

/**
The temperature controller represents handlers for the '/temp/' API-Endpoint.
*/

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	ent "cityserver/entities"
	mid "cityserver/middleware"
	repo "cityserver/repositories"
	serv "cityserver/services"
)

// TemperatureList encodes a list of temperatures
type TemperatureList struct {
	List []ent.CityWithTemp
}

// GetTemperatures returns an array of temperature values given a city collection id
func GetTemperatures(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mid.LogRequest(r)

	var cityIDList []int
	var temperatureVector []ent.CityWithTemp
	var collectionID = ps.ByName("id")
	var cityCollection, err = repo.Read(collectionID)

	for _, city := range cityCollection.Cities {
		cityIDList = append(cityIDList, city.OwmID)
	}

	temperatureVector, err = serv.QueryCityList(cityIDList)
	if err.Err != nil {
		mid.HandleError(w, err)
	} else {
		respondTemp(w, temperatureVector)
	}
}

func respondTemp(w http.ResponseWriter, temperatures []ent.CityWithTemp) {
	var temperatureReturn = TemperatureList{
		List: temperatures,
	}

	header := w.Header()
	header.Set("Access-Control-Allow-Methods", "*")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(temperatureReturn)
}

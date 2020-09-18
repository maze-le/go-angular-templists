package services

/**
The OWM Service encapsulates methods for http requests to the open weather maps API
*/

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	ent "cityserver/entities"
	mid "cityserver/middleware"
)

type owmMain struct {
	Temp float32
}

type owmReturn struct {
	ID   int
	Name string
	Main owmMain
}

type owmListReturn struct {
	ID   int
	Name string
	Cnt  int
	List []owmReturn
}

var httpClient = &http.Client{Timeout: 20 * time.Second}
var accessKey string
var owmDataURLBase = "https://api.openweathermap.org/data/2.5/weather"
var owmGroupURLBase = "https://api.openweathermap.org/data/2.5/group"

// dummy return for error cases
var d = ent.CityWithTemp{}

// InitializeOWMService initializes the owm service by reading the access key from the env
func InitializeOWMService() {
	accessKey = os.Getenv("OWM_ACCESS_KEY")
	mid.LogInfo("initialized owm service")
}

// QueryCityByName queries the openweather service to find city with 'name'
func QueryCityByName(name string) (ent.CityWithTemp, *mid.HTTPError) {
	var cityURL string = owmCityURL(name)
	var owmReturn = &owmListReturn{}

	mid.LogInfo("performing owm query to: " + cityURL)

	var err = getJSON(cityURL, owmReturn)
	if err.Err != nil {
		return d, err
	}

	if owmReturn.ID == 0 {
		return d, mid.Throw404("City not found!")
	}

	return ent.CityWithTemp{
		Name:  owmReturn.Name,
		OwmID: owmReturn.ID,
	}, nil
}

// QueryCityList queries the openweather service to find a list of cities by their ids 'cityIDs'
func QueryCityList(cityIDs []int) ([]ent.CityWithTemp, *mid.HTTPError) {
	var returnList []ent.CityWithTemp
	var groupURL string = owmGroupURL(cityIDs)
	var listReturn = &owmListReturn{}

	mid.LogInfo("performing owm query to: " + groupURL)

	var err = getJSON(groupURL, listReturn)
	if err.Err != nil {
		mid.LogInfo("FOOOOOO: " + err.Message)
		mid.LogError(err.Err)
		return returnList, err
	}

	for _, entity := range listReturn.List {
		returnList = append(returnList, ent.CityWithTemp{
			Name:  entity.Name,
			OwmID: entity.ID,
			Temp:  entity.Main.Temp,
		})
	}

	return returnList, &mid.HTTPError{Err: nil}
}

// returns the url used to query cities
func owmCityURL(cityName string) string {
	urlstr := []string{owmDataURLBase, "?q=", cityName, "&APPID=", accessKey, "&units=metric"}
	return strings.Join(urlstr, "")
}

// returns the url used to query a list of cities
func owmGroupURL(cityIDs []int) string {
	var IDs []string
	for _, i := range cityIDs {
		IDs = append(IDs, strconv.Itoa(i))
	}

	urlstr := []string{owmGroupURLBase, "?id=", strings.Join(IDs, ","), "&APPID=", accessKey, "&units=metric"}
	return strings.Join(urlstr, "")
}

// http client request wrapped in a json formatter and error handler
func getJSON(url string, target *owmListReturn) *mid.HTTPError {
	r, err := httpClient.Get(url)
	if err != nil {
		return mid.Throw500("owm: too many redirects or other protocol errors")
	}

	switch status := r.StatusCode; status {
	case 200:
		defer r.Body.Close()
		json.NewDecoder(r.Body).Decode(&target)

	case 401:
		return mid.Throw401("owm: access key invalid")

	case 404:
		return mid.Throw404("owm: could not find entity")

	default:
		return mid.Throw500("owm: error while querying the owm service")
	}

	return &mid.HTTPError{Err: nil}
}

package entities

import "github.com/jinzhu/gorm"

// City -- the city represents a city object as queried by the openweather api
type City struct {
	gorm.Model
	CityID           uint
	OwmID            int
	Name             string
	CityCollectionID uint
}

// CityWithTemp -- a city with temperature
type CityWithTemp struct {
	OwmID int
	Name  string
	Temp  float32
}

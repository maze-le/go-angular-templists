package entities

import "github.com/jinzhu/gorm"

// CityCollection -- bundles multiple city ids
type CityCollection struct {
	gorm.Model
	CityCollectionID uint
	Name             string
	Cities           []City
}

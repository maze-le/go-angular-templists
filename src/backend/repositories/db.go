package repositories

import (
	"github.com/jinzhu/gorm"
	// imports the pg dialect explicitly
	_ "github.com/jinzhu/gorm/dialects/postgres"

	ent "cityserver/entities"
	mid "cityserver/middleware"
)

var dbDriver = "postgres"
var dbConnectionOpts = "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres"
var dbConnection *gorm.DB

// InitializeDB brings the database into an initialized state
func InitializeDB(connection *gorm.DB) {
	connection.AutoMigrate(&ent.City{})
	connection.AutoMigrate(&ent.CityCollection{})

	var initialContent []ent.CityCollection

	// populate the database with initial fixtures if it is empty
	connection.Find(&initialContent)
	if len(initialContent) == 0 {
		populateDatabase(connection)
	}
}

// Connect2DB establishes the database connection
func Connect2DB() *gorm.DB {
	var connection, err = gorm.Open(dbDriver, dbConnectionOpts)
	if err != nil {
		mid.LogFatal(err)
	}

	dbConnection = connection
	mid.LogInfo("database connection established")

	return connection
}

// GetDbConnection returns the shared database connection pointer
func GetDbConnection() *gorm.DB {
	return dbConnection
}

// populate Database
func populateDatabase(connection *gorm.DB) {

	// initial fixtures
	var leipzig = ent.City{Name: "Leipzig", OwmID: 6548737}
	var berlin = ent.City{Name: "Berlin", OwmID: 2950159}
	var moscow = ent.City{Name: "Moscow", OwmID: 524901}
	var tokyo = ent.City{Name: "Tokyo", OwmID: 1850144}

	connection.Create(&leipzig)
	connection.Create(&berlin)
	connection.Create(&moscow)
	connection.Create(&tokyo)

	var collections = []ent.CityCollection{
		{Name: "My first collection", Cities: []ent.City{berlin, leipzig}},
		{Name: "The 2nd collection", Cities: []ent.City{moscow, tokyo}},
	}

	for index := range collections {
		connection.Create(&collections[index])
	}

	mid.LogInfo("populated the database with initial values")
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"

	controller "cityserver/controllers"
	mid "cityserver/middleware"
	repo "cityserver/repositories"
	serv "cityserver/services"
)

// index serves as a dummy API root for the router
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "City Collections API.\n")
}

// cors handler
func options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mid.LogInfo("GlobalOPTIONS")

	// Set CORS headers
	header := w.Header()
	header.Set("Access-Control-Allow-Methods", "*")
	header.Set("Access-Control-Allow-Origin", "*")

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprint(w, "")
}

// initializeRouter initializes the httprouter
func initializeRouter() *httprouter.Router {
	router := httprouter.New()

	// enables CORS preflight-requests for all routes and all sources
	router.HandleOPTIONS = true
	router.OPTIONS("/cities/:all", options)
	router.OPTIONS("/city/:all", options)
	router.OPTIONS("/temp/:all", options)

	// root route
	router.GET("/", index)

	// routes for 'controllers/citiescontroller'
	router.GET(
		"/cities/all", controller.CollectionIndex)
	router.GET(
		"/city/:id", controller.ReadCollection)
	router.POST(
		"/city/:id", options)
	router.POST(
		"/city/:id/:cityId", controller.UpdateCollection)
	router.PUT(
		"/city/:name", controller.CreateCollection)
	router.DELETE(
		"/city/:id", controller.DeleteCollection)

	// routes for 'controllers/tempcontroller'
	router.GET(
		"/temp/:id", controller.GetTemperatures)

	mid.LogInfo("initialized routing module")

	return router
}

// initializes the os-signal handling mechanism, to ensure graceful application shutdown
func initializeSignalHandler(connection *gorm.DB) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel

		switch sig {
		// handle SIGINT
		case os.Interrupt:
			mid.LogInfo("received SIGINT signal")
			defer connection.Close()

			mid.LogInfo("graceful application shutdown complete")
			os.Exit(0)

		// handle SIGTERM
		case syscall.SIGTERM:
			mid.LogInfo("received SIGTERM signal")
			defer connection.Close()

			mid.LogInfo("application terminated")
			os.Exit(-1)
		}
	}()

	mid.LogInfo("initalized signal handler")
}

// the application main entry point starts the http router and server and exits
// on fatal server errors
func main() {
	var connection = repo.Connect2DB()
	repo.InitializeDB(connection)
	initializeSignalHandler(connection)
	serv.InitializeOWMService()

	router := initializeRouter()

	mid.LogInfo("starting http server on port: 8082")
	mid.LogFatal(http.ListenAndServe(":8082", router))
}

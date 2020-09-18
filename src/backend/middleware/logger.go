package middleware

/** The logger middleware encapsulates logging methods throughout the application */

import (
	"fmt"
	"log"
	"net/http"
)

// LogRequest implemens a request logger
func LogRequest(r *http.Request) {
	LogInfo("" + r.Method + " " + r.RequestURI + " from " + r.RemoteAddr)
}

// LogInfo logs server operational messages
func LogInfo(message string) {
	var logMessage = fmt.Sprintf("info: %s", message)
	log.Println(logMessage)
}

// LogError logs server errors
func LogError(err error) {
	log.Println("error:")
	log.Println(err)
}

// LogFatal logs critical server errors and exits the process
func LogFatal(err error) {
	log.Println("fatal:")
	log.Fatal(err)
}

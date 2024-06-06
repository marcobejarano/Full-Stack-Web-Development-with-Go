package main

import (
	"bytes"
	"encoding/json"
	"log"
)

// main function to show logging using standard library
func main() {
	ol := log.Default() // Get a new logger that writes to standard error and has the default settings.

	// set log format to - dd/mm/yy hh:mm:ss
	ol.SetFlags(log.LstdFlags)    // Set the output flags for the logger. Here it's set to log the date and time.
	ol.Println("Just a log text") // Print a log message.
	lognumber(ol)                 // Call the lognumber function.
	logjson(ol)                   // Call the logjson function.
}

// logjson function to log json to logger
func logjson(ol *log.Logger) {
	ol.SetFlags(log.Ltime) // Set the output flags for the logger. Here it's set to log the time.

	// JSON string
	ex := `{"name": "Cake","batters":{"batter":[{ "id": "001", "type": "Good Food" }]},"topping":[{ "id": "002", "type": "Syrup" }]}`

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, []byte(ex), "", "\t") // Indent the JSON string for better readability.
	if error != nil {
		ol.Fatalf("Error parsing : %s", error.Error()) // If there's an error in indenting, log the error and exit.
	}

	ol.Println(prettyJSON.String()) // Print the indented JSON string.
}

// lognumber function to log number to logger
func lognumber(ol *log.Logger) {
	ol.SetFlags(log.Lshortfile)       // Set the output flags for the logger. Here it's set to log the file name and line number.
	ol.Printf("This is number %d", 1) // Print a formatted log message.
}

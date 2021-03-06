/*
Name:		<REPLACE_NAME>
Author:		Daniel Middleton <me@daniel-middleton.com>
Description:	<REPLACE_DESC>
Usage:		See ./<REPLACE_NAME> -h or --help
*/
package main

// Import package dependancies
import (
	"fmt"
	"os"
	"time"
)

// Define global variables
var (
	progName string = "<REPLACE_NAME>"
)

// Desc: Prints a given message to screen in a uniform format
// Usage: screenOut("STATUS", "message")
func screenOut(status string, message string) {
	timestamp := time.Now().Format("20060102T150405")
	fmt.Printf("[%v][%v][%v]: %v \n", progName, timestamp, status, message)
}

// Desc: Generically handles potential errors
// Usage: checkStatus(err, "Error message", "Success message")
func checkStatus(err error, errorMsg string, successMsg string) {
	if err == nil {
		screenOut("SUCCESS", successMsg)
	} else {
		screenOut("ERROR", errorMsg+" - Exiting...")
		screenOut("ERROR", "Error report:")
		fmt.Println("\n", err, "\n")
		os.Exit(1)
	}
}

// Desc:
// Usage:
func newFunc() {
	screenOut("INFO", "Started newFunc...")

	// Do stuff
	//	checkStatus(err,
	//		"Error message.",
	//		"Success message.")

	screenOut("INFO", "Finished newFunc.")
}

// Desc: Auto-runs on program execution
// Usage: n/a
func init() {
	// Define and parse command line flags
	//wordPtr := flag.String("word", "foo", "a string")
	//numbPtr := flag.Int("numb", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")
	//flag.Parse()

	//fmt.Println("word:", *wordPtr)
	//fmt.Println("numb:", *numbPtr)
	//fmt.Println("fork:", *boolPtr)
}

// Desc: Auto-runs after init function
// Usage: n/a
func main() {
	fmt.Println()
	screenOut("INFO", progName+" script started...")

	// Execute functions

	screenOut("INFO", progName+" script ended.")
	fmt.Println()
	os.Exit(0)
}

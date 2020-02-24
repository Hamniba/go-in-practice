package main

import (
	"log"
	"os"

	_ "Manning/go-in-action/matchers"
	"Manning/go-in-action/search"
)

// All init functions will get called before the main function
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}

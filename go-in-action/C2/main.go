package main

import (
	"log"
	"os"

	_ "Manning/go-in-action/C2/matchers"
	"Manning/go-in-action/C2/search"
)

// All init functions will get called before the main function
func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}

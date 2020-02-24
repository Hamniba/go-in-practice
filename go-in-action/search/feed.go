package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed type
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// GetFeeds read data from data.json
func GetFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// schedule the file to be closed while the function GetFeeds() returns
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}

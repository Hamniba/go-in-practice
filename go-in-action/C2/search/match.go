package search

import (
	"log"
)

// Result type
type Result struct {
	Field   string
	Content string
}

// Matcher interface
type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

// Match search the term
func Match(matcher Matcher, feed *Feed, term string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, term)
	if err != nil {
		log.Println(err)
		return
	}

	// write the results to the channel
	for _, result := range searchResults {
		results <- result
	}
}

// Display log the search result
func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}

package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

// Run the main search logic
func Run(term string) {
	feeds, err := GetFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// launch the goroutine
		go func(m Matcher, f *Feed) {
			// call Match function in match.go
			Match(m, f, term, results)
			waitGroup.Done()
		}(matcher, feed)

		go func() {
			// block the goroutine until the count for waitGroup is zero
			waitGroup.Wait()
			close(results)
		}()

		Display(results)
	}
}

// Register regist the matcher
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

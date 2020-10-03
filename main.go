package main

import (
	"fmt"

	actions "./actions"
	flickr "./flickr"
)

func main() {
	var apiKey = "0bc2f0f2743df78c0764103b16222110" // todo - to input
	var userID = "106265895@N05"                    // todo - to input
	fmt.Printf("Hi! Fetching photos info for user %s\n", userID)

	var callOptions = flickr.CallOptions{
		APIKey:  apiKey,
		UserID:  userID,
		PerPage: 100,
	}

	// 1. get the # of pages
	var result = flickr.Call("flickr.people.getPhotos", callOptions, 0)

	// 2. call each page data in parallel
	results := actions.Dispatch(result.PagesCount, callOptions)

	// 3. get all names, distinct and aggregate # of birds
	names := actions.GetBirdNames(results)

	fmt.Printf("\n Overall species: %d\n", len(names))

	for _, name := range names {
		fmt.Printf("\t %s\n", name)
	}

	// 4. data per year (todo)
	// 5. when was first occurence? (todo)
}

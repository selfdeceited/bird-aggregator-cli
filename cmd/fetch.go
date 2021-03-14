package cmd

import (
	"fmt"

	actions "github.com/selfdeceited/bird-aggregator-cli/actions"
	flickr "github.com/selfdeceited/bird-aggregator-cli/flickr"
	"github.com/spf13/cobra"
)

var fetchCommand = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch all bird names",
	Run: func(cmd *cobra.Command, args []string) {
		fetch()
	},
}

func fetch() {
	fmt.Printf("Hi! Fetching photos info for user %s\n", userId)

	var callOptions = flickr.CallOptions{
		APIKey:  apiKey,
		UserID:  userId,
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

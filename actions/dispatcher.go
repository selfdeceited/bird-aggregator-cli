package actions

import flickr "../flickr"

func askForPhotos(pageNumber int, callOptions flickr.CallOptions, c chan flickr.CallResult) {
	var result = flickr.Call("flickr.people.getPhotos", callOptions, pageNumber)
	c <- result
}

func Dispatch(pages int, callOptions flickr.CallOptions) []flickr.CallResult {
	c := make(chan flickr.CallResult)
	var results = []flickr.CallResult{}
	for i := 0; i < pages; i++ {
		go askForPhotos(i, callOptions, c)
	}

	for i := 0; i < pages; i++ {
		v := <-c
		results = append(results, v)
	}

	return results
}

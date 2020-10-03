package flickr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "../models"
)

//CallOptions - required
type CallOptions struct {
	APIKey  string
	UserID  string
	PerPage int
}

//CallResult - required
type CallResult struct {
	PagesCount int
	PhotoNames []string
}

//Call - used to call Flickr API
func Call(method string, options CallOptions, pageNumber int) CallResult {
	var url = "https://api.flickr.com/services/rest"

	client := &http.Client{}
	var finalURL = fmt.Sprintf("%s?method=%s&api_key=%s&format=json&nojsoncallback=1&per_page=%d&user_id=%s&page=%d",
		url, method, options.APIKey, options.PerPage, options.UserID, pageNumber)

	req, err := http.NewRequest("GET", finalURL, nil)
	req.Header.Add("Cache-Control", "no-cache")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf(err.Error())
		return CallResult{}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	res := models.PhotosResponse{}
	json.Unmarshal([]byte(body), &res)

	return CallResult{
		PagesCount: res.Photos.Pages,
		PhotoNames: GetPhotoNames(res),
	}
}

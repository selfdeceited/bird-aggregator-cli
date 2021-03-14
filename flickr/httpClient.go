package flickr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/selfdeceited/bird-aggregator-cli/models"
)

type CallOptions struct {
	APIKey  string
	UserID  string
	PerPage int
}

type CallResult struct {
	PagesCount int
	PhotoNames []string
}

func Call(method string, options CallOptions, pageNumber int) CallResult {
	var url = "https://api.flickr.com/services/rest"

	client := &http.Client{}
	var finalURL = fmt.Sprintf("%s?method=%s&api_key=%s&format=json&nojsoncallback=1&per_page=%d&user_id=%s&page=%d",
		url, method, options.APIKey, options.PerPage, options.UserID, pageNumber)

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		fmt.Print(err.Error())
		return CallResult{}
	}

	req.Header.Add("Cache-Control", "no-cache")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		return CallResult{}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
		return CallResult{}
	}

	res := models.PhotosResponse{}
	json.Unmarshal([]byte(body), &res)

	return CallResult{
		PagesCount: res.Photos.Pages,
		PhotoNames: GetPhotoNames(res),
	}
}

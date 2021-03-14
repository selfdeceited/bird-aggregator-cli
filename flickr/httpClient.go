package flickr

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Call(method string, options CallOptions, pageNumber int) CallResult {
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"method":         method,
			"api_key":        options.APIKey,
			"format":         "json",
			"nojsoncallback": "1",
			"per_page":       fmt.Sprint(options.PerPage),
			"user_id":        options.UserID,
			"page":           fmt.Sprint(pageNumber),
		}).
		SetHeader("Cache-Control", "no-cache").
		Get("https://api.flickr.com/services/rest")

	if err != nil {
		fmt.Print(err.Error())
		return CallResult{}
	}

	res := PhotosResponse{}
	json.Unmarshal(resp.Body(), &res)

	return CallResult{
		PagesCount: res.Photos.Pages,
		PhotoNames: GetPhotoNames(res),
	}
}

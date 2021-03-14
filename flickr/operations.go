package flickr

import (
	"strings"

	operators "github.com/selfdeceited/bird-aggregator-cli/operators"
)

func Map(vs []Photo, f func(Photo) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func GetPhotoNames(res PhotosResponse) []string {
	return operators.Filter(Map(res.Photos.Photo, func(v Photo) string {
		return v.Title
	}), func(title string) bool {
		return strings.HasPrefix(title, "B: ")
	})
}

func FlatMapNames(res []CallResult) []string {
	flatNames := []string{}
	for _, v := range res {
		flatNames = append(flatNames, v.PhotoNames...)
	}
	return flatNames
}

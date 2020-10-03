package flickr

import (
	"strings"

	models "../models"
	operators "../operators"
)

func Map(vs []models.Photo, f func(models.Photo) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func GetPhotoNames(res models.PhotosResponse) []string {
	return operators.Filter(Map(res.Photos.Photo, func(v models.Photo) string {
		return v.Title
	}), func(title string) bool {
		return strings.HasPrefix(title, "B: ")
	})
}

func FlatMapNames(res []CallResult) []string {
	vsm := []string{}
	for _, v := range res {
		for _, name := range v.PhotoNames {
			vsm = append(vsm, name)
		}
	}
	return vsm
}

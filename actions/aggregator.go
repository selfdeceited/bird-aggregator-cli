package actions

import (
	"strings"

	flickr "github.com/selfdeceited/bird-aggregator-cli/flickr"
	operators "github.com/selfdeceited/bird-aggregator-cli/operators"
)

func GetBirdNames(callResults []flickr.CallResult) []string {
	photoTitles := flickr.FlatMapNames(callResults)
	birdNames := []string{}
	for _, v := range photoTitles {
		v = strings.TrimPrefix(v, "B: ")
		localNames := operators.Map(strings.Split(v, ","), func(rawName string) string {
			return strings.Trim(rawName, " ")
		})

		for _, name := range localNames {
			if !operators.Include(birdNames, name) {
				birdNames = append(birdNames, name)
			}
		}
	}
	return birdNames
}

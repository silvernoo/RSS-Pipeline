package rss

import (
	"github.com/mmcdole/gofeed"
)

func GetRssItem(feed string) []string {
	var links []string
	fp := gofeed.NewParser()
	feed2, _ := fp.ParseURL(feed)
	for _, i := range feed2.Items {
		links = append(links, i.Enclosures[0].URL)
	}
	return links
}

package feed

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/edubarbieri/ddm/config"
)

type Feed struct {
	Links []Item `xml:"channel>item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	EpisodeID int    `xml:"http://showrss.info episode_id"`
}

// GetFeed process feed
func GetFeed() (Feed, error) {
	resp, err := http.Get(config.Data.FeedURL)
	var feed Feed
	if err != nil {
		return feed, err
	}
	defer resp.Body.Close()
	content, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return feed, err2
	}
	err3 := xml.Unmarshal(content, &feed)
	if err3 != nil {
		return feed, err3
	}
	return feed, nil
}

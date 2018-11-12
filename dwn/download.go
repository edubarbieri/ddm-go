package dwn

import (
	"log"

	"github.com/edubarbieri/ddm/data"
	"github.com/edubarbieri/ddm/feed"
	"github.com/edubarbieri/ddm/trm"
)

// Process download
func Process() {
	download()
	trm.RemoveCompletes()
}

func download() {
	log.Printf("Processing Feed...")
	feed, error := feed.GetFeed()
	if error != nil {
		log.Println("error processing download", error)
		return
	}
	for _, item := range feed.Links {
		_, err := data.GetFeedByEpisodeID(item.EpisodeID)
		if err == nil {
			continue
		}
		log.Printf("Adding Item %v ", item.Title)
		error := trm.AddInTransmission(&item)
		if error == nil {
			data.InsertFeed(item.EpisodeID, item.Title)
		} else {
			log.Printf("error adding %s, %v\n", item.Title, error)
		}
	}
}

package trm

import (
	"fmt"
	"log"

	"github.com/edubarbieri/ddm/config"
	"github.com/edubarbieri/ddm/feed"
	"github.com/odwrtw/transmission"
)

var client *transmission.Client

func getClient() (*transmission.Client, error) {
	if client == nil {
		log.Println("Creating new transmission client")
		conf := transmission.Config{
			Address:  fmt.Sprintf("http://%s/transmission/rpc", config.Data.Transmission.Host),
			User:     config.Data.Transmission.User,
			Password: config.Data.Transmission.Password,
		}
		var err error
		client, err = transmission.New(conf)
		return client, err
	}
	return client, nil
}

//AddInTransmission add new torrent in transmission
func AddInTransmission(item *feed.Item) error {
	t, err := getClient()
	if err != nil {
		log.Println(err)
		return err
	}
	_, err2 := t.Add(item.Link)
	return err2
}

//RemoveCompletes remove completes downloads in transmission
func RemoveCompletes() {
	if !config.Data.Transmission.RemoveCompletes {
		return
	}
	log.Printf("Removing completes...")
	client, err := getClient()
	if err != nil {
		log.Println("Error on get client", err)
		return
	}
	torrents, _ := client.GetTorrents()
	var tRemove []*transmission.Torrent
	for _, t1 := range torrents {
		if t1.IsFinished {
			log.Printf("Torrent %v is complete\n", t1.Name)
			tRemove = append(tRemove, t1)
		}
	}
	errRemove := client.RemoveTorrents(tRemove, false)
	if errRemove != nil {
		log.Println("Error removing torrent ", errRemove)
	}
}

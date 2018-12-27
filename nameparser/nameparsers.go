package nameparser

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/edubarbieri/ddm/data"
	"github.com/edubarbieri/ddm/model"
)

//Process method parse file name, search information about episode and set destination path
func Process(fileData *model.FileData) {
	id, err := getSerieID(fileData)
	if err == nil {
		ep, err := NewTvdbClient().GetEpisode(id, fileData.Season, fileData.Episode)
		if err != nil {
			log.Println("Could not connect to TVDB to get EpisodeData", err)
		}
		if ep.StatusCode == 404 {
			log.Printf("Not find data %v Season %v, Episode %v\n", fileData.Name, fileData.Season, fileData.Episode)
		} else if ep.StatusCode >= 500 {
			log.Panicf("Server error on serch data %v Season %v, Episode %v, Code %v\n", fileData.Name, fileData.Season, fileData.Episode, ep.StatusCode)
		}
		if len(ep.Data) > 0 {
			fileData.EpisodeName = ep.Data[0].EpisodeName
		}
	}
	formatFinalName(fileData)
}

func getSerieID(fileData *model.FileData) (int, error) {
	lowerName := strings.ToLower(fileData.Name)
	searchKey := strings.Join(strings.Split(lowerName, " "), "_")
	s, err := data.GetSerieBySearckKey(searchKey)
	if err == nil {
		fileData.BeautifulName = s.Name
		return s.TvdbID, nil
	}
	log.Printf("Searching serie id with args: %v\n", lowerName)
	tvdbClient := NewTvdbClient()
	resp, err := tvdbClient.SearchSeries(lowerName)
	if err != nil {
		log.Panic("Could not connect to TVDB", err)
	}
	if len(resp.Data) > 0 {
		firstSerie := resp.Data[0]
		data.InsertSerie(&data.Serie{
			Name:      firstSerie.SeriesName,
			TvdbID:    firstSerie.ID,
			SearchKey: searchKey,
		})
		fileData.BeautifulName = firstSerie.SeriesName
		return firstSerie.ID, nil
	}
	return 0, errors.New("could not find serie in TVDB")
}

func formatFinalName(f *model.FileData) {
	//this is format: {name}/Season {season}/{name} - S{season}E{episode} - {title}"
	var name string
	if f.BeautifulName != "" {
		name = f.BeautifulName
	} else {
		name = f.Name
	}

	if f.EpisodeName == "" {
		f.BeautifulName = fmt.Sprintf("%s/Season %02d/%02s - S%02dE%02d", name, f.Season, name, f.Season, f.Episode)
	} else {
		f.BeautifulName = fmt.Sprintf("%s/Season %02d/%02s - S%02dE%02d - %s", name, f.Season, name, f.Season, f.Episode, f.EpisodeName)
	}

}

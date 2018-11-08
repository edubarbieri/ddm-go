package data

import (
	"fmt"
)

// Serie representa as informações de uma serie
type Serie struct {
	ID        int
	Name      string
	TvdbID    int
	SearchKey string
}

func (s Serie) String() string {
	return fmt.Sprintf("Serie ID: %v, Name: %v, TvdbID: %v, SearchKey: %v",
		s.ID, s.Name, s.TvdbID, s.SearchKey)
}

// Feed representa um item de feed ja processado
type Feed struct {
	ID        int
	EpisodeID int
	Title     string
}

func (s Feed) String() string {
	return fmt.Sprintf("Feed ID: %v, EpisodeID: %v, Title: %v", s.ID, s.EpisodeID, s.Title)
}

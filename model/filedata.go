package model

import "fmt"

//FileData represent a information about file to be moved
type FileData struct {
	Path          string
	Name          string
	Season        int
	Episode       int
	NewPath       string
	BeautifulName string
	EpisodeName   string
}

func (s FileData) String() string {
	return fmt.Sprintf("FileData{%#v}", s)
}

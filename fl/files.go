package fl

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/edubarbieri/ddm/config"
)

//FileData represent a information about file to be moved
type FileData struct {
	Path    string
	Name    string
	Season  int
	Episode int
	NewPath string
}

func getFileData(fileName string) (FileData, error) {
	r, _ := regexp.Compile("(.*)[sS](\\d{2})[eE](\\d{2})")
	var fileData FileData
	match := r.MatchString(fileName)
	if match {
		matchs := r.FindStringSubmatch(fileName)
		regReplace, _ := regexp.Compile("[^a-zA-Z0-9]+")
		fileData.Name = strings.Trim(regReplace.ReplaceAllString(matchs[1], " "), " ")
		season, _ := strconv.Atoi(matchs[2])
		fileData.Season = season
		epis, _ := strconv.Atoi(matchs[3])
		fileData.Episode = epis
		return fileData, nil
	}
	return fileData, errors.New("not match fileName")
}
func isValidExt(ext string) bool {
	for _, n := range config.Data.VideoExts {
		if strings.ToLower(ext) == strings.ToLower(n) {
			return true
		}
	}
	return false
}

func setNewPath(fileData *FileData) {

}

func getPedingMove() []FileData {
	var files []FileData
	filepath.Walk(config.Data.SourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isValidExt(filepath.Ext(path)) {
			ep, err := getFileData(info.Name())
			if err == nil {
				ep.Path = path
				setNewPath(&ep)
				files = append(files, ep)
			}
		}
		return nil
	})
	return files
}

func Teste() {
	fmt.Println(getPedingMove())
}

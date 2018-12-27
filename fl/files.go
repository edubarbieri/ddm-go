package fl

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/edubarbieri/ddm/config"
	"github.com/edubarbieri/ddm/model"
	"github.com/edubarbieri/ddm/nameparser"
)

func MovePending() []model.FileData {
	var files []model.FileData
	filepath.Walk(config.Data.SourceFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Panicf("Error reading path %v: %v", path, err)
			return err
		}
		if !info.IsDir() && isValidExt(filepath.Ext(path)) {
			ep, err := GetFileData(path, info)
			if err == nil {
				log.Printf("moving %v to %v", ep.Path, ep.NewPath)
				if !config.Data.TestMode {
					os.Rename(ep.Path, ep.NewPath)
				}
			}
		}
		return nil
	})
	return files
}

//GetFileData Parse file to File Data
func GetFileData(path string, info os.FileInfo) (model.FileData, error) {
	fileData, err := getFileData(info.Name())
	if err != nil {
		return fileData, err
	}
	fileData.Path = path
	setNewPath(&fileData)
	return fileData, nil
}

func getFileData(fileName string) (model.FileData, error) {
	r, _ := regexp.Compile("(.*)[sS](\\d{2})[eE](\\d{2})")
	var fileData model.FileData
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

func setNewPath(fileData *model.FileData) {
	nameparser.Process(fileData)
	newPath := filepath.Join(config.Data.TargetFolder, fileData.BeautifulName)
	fileData.NewPath = SafeExist(newPath)
}
func SafeExist(newPath string) string {
	count := 1
	tmpPath := newPath

	for {
		if e, _ := Exists(tmpPath); !e {
			return tmpPath
		}
		ext := filepath.Ext(newPath)
		tmpPath = strings.Replace(newPath, ext, "", 1) + " " + strconv.Itoa(count) + ext
		count++
	}
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

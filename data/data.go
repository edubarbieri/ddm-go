package data

import (
	"database/sql"
	"log"

	//Import to database driver work
	_ "github.com/mattn/go-sqlite3"
)

var initialSQL = [2]string{
	`
		CREATE TABLE IF NOT EXISTS serie (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			tvdbId INTEGER,
			searchKey TEXT
		)
	`,
	`
		CREATE TABLE IF NOT EXISTS feed (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			episodeId INTEGER,
			title TEXT
		)
	`,
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./ddm.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func init() {
	log.Printf("initializing database...")
	defer log.Printf("database initialized with success...")
	db := openDb()
	defer db.Close()
	for _, sql := range initialSQL {
		_, error := db.Exec(sql)
		if error != nil {
			log.Fatal(error)
		}
	}

}

// ListAllSeries return all series in database
func ListAllSeries() []Serie {

	db := openDb()
	defer db.Close()
	rows, _ := db.Query("select id, name, tvdbId, searchKey from serie order by name")
	defer rows.Close()
	var series []Serie

	for rows.Next() {
		var serie Serie
		rows.Scan(&serie.ID, &serie.Name, &serie.TvdbID, &serie.SearchKey)
		series = append(series, serie)
	}
	return series
}

//SaveSerie Save or update serie
func SaveSerie(serie *Serie) {
	if serie.ID == 0 {
		InsertSerie(serie)
	} else {
		UpdateSerie(serie)
	}
}

// InsertSerie Create new serie in database
func InsertSerie(serie *Serie) {
	db := openDb()
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO serie(name, tvdbId, searchKey) VALUES (?, ?, ?)")

	res, err := stmt.Exec(serie.Name, serie.TvdbID, serie.SearchKey)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	lastID, _ := res.LastInsertId()
	serie.ID = int(lastID)

	tx.Commit()
}

// UpdateSerie Update exist serie in database
func UpdateSerie(serie *Serie) {
	db := openDb()
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("update serie set name = ?, tvdbId = ?, searchKey = ? WHERE id = ?")
	_, err := stmt.Exec(serie.Name, serie.TvdbID, serie.SearchKey, serie.ID)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}

// GetSerieBySearckKey find serie by search key
func GetSerieBySearckKey(searchKey string) (Serie, error) {
	db := openDb()
	defer db.Close()

	var serie Serie
	row := db.QueryRow("select id, name, tvdbId, searchKey from serie where searchKey = ?", searchKey)
	err := row.Scan(&serie.ID, &serie.Name, &serie.TvdbID, &serie.SearchKey)
	if err != nil {
		log.Printf("Could not find serie with searchKey %v", searchKey)
		return serie, err
	}
	return serie, nil
}

// InsertFeed Create new feed in database
func InsertFeed(feed *Feed) {
	db := openDb()
	defer db.Close()
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO feed(episodeId, title)' VALUES (?, ?)")

	res, err := stmt.Exec(feed.EpisodeID, feed.Title)

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	lastID, _ := res.LastInsertId()
	feed.ID = int(lastID)
	tx.Commit()
}

// GetFeedByEpisodeID find feed by episodeId
func GetFeedByEpisodeID(episodeID int) (Feed, error) {
	db := openDb()
	defer db.Close()

	var feed Feed
	row := db.QueryRow("SELECT id, episodeId, title FROM feed where episodeId = ?", episodeID)
	err := row.Scan(&feed.ID, &feed.EpisodeID, &feed.Title)
	if err != nil {
		log.Printf("Could not find feed with episodeId %v", episodeID)
		return feed, err
	}
	return feed, nil
}

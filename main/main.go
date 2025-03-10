package main

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct{
	ID 		int64
	Title 	string
	Artist 	string
	Price 	float32
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to database!")

	albums, err := albumsByArtist("John Coltrane")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	alb, err := albumByID(2)
	if err != nil {
    log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	albID, err := addAlbum(Album{
		Title: "The Modern Sounf of Betty Carter",
		Artist: "Betty Cartr",
		Price: 49.95,
	})
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %d\n", albID)

	albumRemoved, err := removeByID(8)
	if err != nil {
    log.Fatal(err)
	}
	fmt.Printf("Album ID removed: %v\n", albumRemoved)

}

func albumsByArtist(name string) ([]Album, error){
	var albums []Album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil{
		return nil, fmt.Errorf("albumsByArtist: %q: %v", name, err)
	}
	defer rows.Close()

	for rows.Next(){
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil{
			return nil, fmt.Errorf("albumsByArtist: %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err!=nil{
		return nil, fmt.Errorf("albumsByArtist: %q: %v", name, err)
	}
	return albums, nil
}

func albumByID(id int64) (Album, error){
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil{
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumByID: %d: no such album", id)
		}
		return alb, fmt.Errorf("albumByID: %d: %v", id, err)
	}
	return alb, nil
}

func addAlbum(a Album) (int64, error){
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", a.Title, a.Artist, a.Price)
	if err !=nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func removeByID(id int64) (int64, error){
	result, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("removeByID: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("removeByID: %v", err)
	}
	if rowsAffected == 0 {
        return 0, fmt.Errorf("removeByID: no album found with id %d", id)
    }

	return id, nil
}
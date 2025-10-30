package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.NewConfig()
	cfg.User = "root"
	cfg.Passwd = "123456"
	cfg.Addr = "127.0.0.1"
	cfg.DBName = "recordings"

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("连接成功")
	albums, err := albumsByArtist("John Coltrane", db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
	alb, err := albumByID(2, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)
	album := Album{Title: "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99}
	id, err := addAlbum(album, db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("新建id= %d", id)
}

func albumsByArtist(name string, db *sql.DB) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var album Album
		rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func albumByID(id int64, db *sql.DB) (Album, error) {
	var album Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("未找到id=%v", id)
		}
		return album, fmt.Errorf("albumById %v %v", id, err)
	}

	return album, nil
}

func addAlbum(album Album, db *sql.DB) (int64, error) {
	result, err := db.Exec("insert into album(title,artist,price)values(?,?,?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

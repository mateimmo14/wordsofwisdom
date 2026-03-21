package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/mateimmo14/wordsofwisdom/internal/database"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type WisdomResponse struct {
	Wisdom string `json:"wisdom"`
	From   string `json:"author"`
}

var connStr = os.Getenv("DATABASE_URL")

type Wisdom struct {
	Id   int
	Data string
	From string
}

var db, err = sql.Open("postgres", connStr)

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	var Wisdoms []Wisdom

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	tmpl := template.Must(template.ParseFiles("index.html"))
	rows, err := db.Query("SELECT id,data,author FROM wisdoms ORDER BY id DESC")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for rows.Next() {
		var w Wisdom
		err = rows.Scan(&w.Id, &w.Data, &w.From)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		Wisdoms = append(Wisdoms, w)
	}
	err = tmpl.Execute(w, Wisdoms)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
func handleSubmit(w http.ResponseWriter, r *http.Request) {
	queries := database.New(db)
	var resp WisdomResponse
	x, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(x, &resp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	final := sql.NullString{
		String: resp.Wisdom,
		Valid:  true,
	}
	finalfinal := database.AddWisdomParams{
		Data:   final,
		Author: resp.From,
	}
	_, err = queries.AddWisdom(r.Context(), finalfinal)
	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
	}
}

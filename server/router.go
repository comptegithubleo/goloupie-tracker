package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var templates = template.Must(template.ParseGlob("public/**/*.html"))

func Route() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/artist/", ArtistHandler)
	http.HandleFunc("/map/", MapHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		err := templates.ExecuteTemplate(w, "index", getArtists())
		if err != nil {
			fmt.Println(err.Error())
		}

	} else if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		var args args
		json.Unmarshal(body, &args)

		err = templates.ExecuteTemplate(w, "artists", searchArtist(args))
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/artist/")

	err := templates.ExecuteTemplate(w, "artist", getArtist(id))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func MapHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	id := string(body)

	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(getConcertCoordinates(id))
	//id = `[{"lat":"-22.2745264","lon":"166.442419","name":"Nouméa, Province Sud, Nouvelle-Calédonie, 98800, France","Dates":["15-11-2019"]},{"lat":"-17.5373835","lon":"-149.5659964","name":"Papeʻete, Îles du Vent, Polynésie Française, 98714, France","Dates":["16-11-2019"]},{"lat":"20.6308643","lon":"-87.0779503","name":"Playa del Carmen, Solidaridad, Quintana Roo, México","Dates":["05-12-2019","06-12-2019","07-12-2019","08-12-2019","09-12-2019"]}]`
	w.Write(j)
}

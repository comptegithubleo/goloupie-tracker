package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type artist struct {
	Id           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Concerts     map[string][]string `json:"datesLocations"`
}

func getArtists() []artist {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var artists []artist
	json.Unmarshal(responseData, &artists)

	/* for _, artist := range artists {
		json.Unmarshal(getConcertsOf(artist.Id), &artist)
	} */
	return artists
}

func getArtist(id string) artist {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var artist artist
	json.Unmarshal(responseData, &artist)
	json.Unmarshal(getConcertsOf(strconv.Itoa(artist.Id)), &artist) // insert concert dates in artist

	return artist
}

func getConcertsOf(id string) []byte {

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	return responseData
}

type coordinates struct {
	Lat   string `json:"lat"`
	Lon   string `json:"lon"`
	Name  string `json:"name"`
	Dates []string
}

func getConcertCoordinates(artistId string) []coordinates {

	var allCoordinates []coordinates

	for location, dates := range getArtist(artistId).Concerts {
		address := strings.Split(location, "-")
		response, err := http.Get("https://nominatim.openstreetmap.org/search.php?city=" + address[0] + "&country=" + address[1] + "&format=jsonv2")
		if err != nil {
			fmt.Print(err.Error())
		}
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Print(err.Error())
		}

		if string(responseData) == "[]" {
			response, err := http.Get("https://nominatim.openstreetmap.org/search.php?q=" + location + "&format=jsonv2")
			if err != nil {
				fmt.Print(err.Error())
			}
			responseData, err = io.ReadAll(response.Body)
			if err != nil {
				fmt.Print(err.Error())
			}

			if string(responseData) == "[]" {
				continue
			}
		}

		var object []coordinates

		json.Unmarshal(responseData, &object)
		object[0].Dates = dates

		allCoordinates = append(allCoordinates, object[0])

		time.Sleep(1 * time.Second)
	}

	return allCoordinates
}

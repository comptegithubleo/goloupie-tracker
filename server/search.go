package server

import (
	"fmt"
	"strconv"
	"strings"
)

type args struct {
	UserSearch string         `json:"userSearch"`
	Filters    map[string]int `json:"filters"`
}

// if an artist doesnt correspond to filters, return false
func isArtistInFilters(artist artist, filters map[string]int) bool {
	for filter, value := range filters {
		switch filter {
		case "memberCount":
			if len(artist.Members) != value && value > 0 {
				return false
			}
		case "creationDateMin":
			if artist.CreationDate < value && value > 0 {
				return false
			}
		case "creationDateMax":
			if artist.CreationDate > value && value > 0 {
				return false
			}
		case "firstAlbumMin":
			year, err := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
			if err != nil {
				fmt.Print(err.Error())
			}
			if year < value && value > 0 {
				return false
			}
		case "firstAlbumMax":
			year, err := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
			if err != nil {
				fmt.Print(err.Error())
			}
			if year > value && value > 0 {
				return false
			}
		}
	}
	return true
}

// if a search contains part of artist, returns true
func isArtistInSearch(artist artist, search string) bool {
	if search == "" {
		return true
	} else {
		search = strings.ToLower(search)

		// fmt.Printf("%#v", artist)		idea: deserialize to compare
		if strings.Contains(strings.ToLower(artist.Name), search) ||
			strings.Contains(strconv.Itoa(artist.CreationDate), search) {
			return true
		}

		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), search) {
				return true
			}
		}

		for location := range artist.Concerts {
			if strings.Contains(location, search) {
				return true
			}
		}

	}

	return false
}

func searchArtist(args args) []artist {
	artists := getArtists()
	var result []artist

	for _, artist := range artists {
		if isArtistInFilters(artist, args.Filters) {
			if isArtistInSearch(artist, args.UserSearch) {
				result = append(result, artist)
			}
		}
	}

	return result
}

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

func searchArtist(args args) []artist {
	artists := getArtists()
	var result []artist

	// filtering :
	// Initially TRUE, we need filter1 AND filter2 AND filter3 ... to be valid for the artist. Otherwise, filter is not ok (FALSE)
	for _, artist := range artists {
		IsFilterOk := true
	out:
		for filter, value := range args.Filters {
			switch filter {
			case "memberCount":
				if len(artist.Members) != value && value > 0 {
					IsFilterOk = false
					break out
				}
			case "creationDateMin":
				if artist.CreationDate < value && value > 0 {
					IsFilterOk = false
					break out
				}
			case "creationDateMax":
				if artist.CreationDate > value && value > 0 {
					IsFilterOk = false
					break out
				}
			case "firstAlbumMin":
				year, err := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
				if err != nil {
					fmt.Print(err.Error())
				}
				if year < value && value > 0 {
					IsFilterOk = false
					break out
				}
			case "firstAlbumMax":
				year, err := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
				if err != nil {
					fmt.Print(err.Error())
				}
				if year > value && value > 0 {
					IsFilterOk = false
					break out
				}
			}
		}

		// If filters ok, checking search :
		// Inverted logic, initially FALSE, if contains name OR contains member OR contains ... switch to TRUE.
		if IsFilterOk {
			IsInSearch := false

			if args.UserSearch == "" { //empty search, defaults TRUE
				IsInSearch = true
			} else {
				args.UserSearch = strings.ToLower(args.UserSearch)

				// fmt.Printf("%#v", artist)		idea: deserialize to compare
				if strings.Contains(strings.ToLower(artist.Name), args.UserSearch) ||
					strings.Contains(strconv.Itoa(artist.CreationDate), args.UserSearch) {
					IsInSearch = true
				}

				for _, member := range artist.Members {
					if strings.Contains(strings.ToLower(member), args.UserSearch) {
						IsInSearch = true
						break
					}
				}

				for location := range artist.Concerts {
					if strings.Contains(location, args.UserSearch) {
						IsInSearch = true
						break
					}
				}

			}
			if IsInSearch {
				result = append(result, artist)
			}
		}
	}

	return result
}

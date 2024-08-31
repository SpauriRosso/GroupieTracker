package Serveur

import (
	"encoding/json"
	"fmt"
	"groupie-tracker-filters/src/shared"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/SpauriRosso/dotlog"
)

type Sartists struct {
	Id           int
	Image        string
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	FirstAlbum   string
	CreationDate int
}

var Sug []Sartists
var Sco []Location
var Scd []Dates

type LocationResponse struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type DatesResponse struct {
	Index []Dates `json:"index"`
}
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Geocode struct {
	Lat float64
	Lng float64
}

var Geo []Geocode

func FetchApi() {
	resp, err := http.Get(shared.URL)
	if err != nil {
		dotlog.Error("Can't fetch API - line 25 in middleware.go")
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&Sug)
	if err != nil {
		dotlog.Error("Can't fetch API - line 31 in middleware.go")
	}

	go func() {
		resp, err := http.Get(shared.LocationsURL)
		if err != nil {
			dotlog.Error("Can't fetch API - line 43 in middleware.go")
			return
		}
		defer resp.Body.Close()

		var LocationResp LocationResponse
		err = json.NewDecoder(resp.Body).Decode(&LocationResp)
		if err != nil {
			te := fmt.Sprintf("Error decoding locations JSON: %v", err)
			dotlog.Error(te)
			os.Exit(1)
		}
		Sco = LocationResp.Index
	}()

	go func() {
		resp, err := http.Get(shared.DatesURL)
		if err != nil {
			dotlog.Error("Can't fetch API - line 43 in middleware.go")
			return
		}
		defer resp.Body.Close()

		var DatesResp DatesResponse
		err = json.NewDecoder(resp.Body).Decode(&DatesResp)
		if err != nil {
			te := fmt.Sprintf("Error decoding locations JSON: %v", err)
			dotlog.Error(te)
			os.Exit(1)
		}
		Scd = DatesResp.Index
	}()
	dotlog.Info("Synced API")
}

func GetSugg(query string) []string {
	var Suggestions []string
	var contains, lower = strings.Contains, strings.ToLower

	for _, artists := range Sug {
		if contains(lower(artists.Name), lower(query)) {
			Suggestions = append(Suggestions, artists.Name+"  (Band)")
		}
		for _, member := range artists.Members {
			if contains(lower(member), lower(query)) {
				if member == artists.Name {
					Suggestions = append(Suggestions, member+"  (Artist "+artists.Name+")")
				} else {
					Suggestions = append(Suggestions, member+"  (Member of "+artists.Name+")")
				}
			}
		}
		if contains(lower(strconv.Itoa(artists.CreationDate)), lower(query)) {
			Suggestions = append(Suggestions, "  (Band: "+artists.Name+")")
		}
		if contains(lower(artists.FirstAlbum), lower(query)) {
			Suggestions = append(Suggestions, "  (Album from: "+artists.Name+")")
		}
	}
	for _, loc := range Sco {
		for _, loca := range loc.Locations {
			if contains(lower(loca), lower(query)) {
				Suggestions = append(Suggestions, GetArtistName(loc.ID)+" In concert in: "+loca)
			}
		}
	}
	for _, da := range Scd {
		for _, dat := range da.Dates {
			if contains(lower(dat), lower(query)) {
				Suggestions = append(Suggestions, GetArtistName(da.ID)+" live on "+dat)
			}
		}
	}
	//fmt.Println(Suggestions)
	return Suggestions
}

func GetArtistName(ID int) string {
	var artistName string
	for _, v := range Sug {
		if ID == v.Id {
			artistName = v.Name
		}
	}
	return artistName
}

func GetLoc(city string) (float64, float64) {
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%v&key=%v", city, shared.MapAPI)
	resp, err := http.Get(url)
	if err != nil {
		te := fmt.Sprintf("%v", err)
		dotlog.Error(te)
		return 0, 0
	}
	defer resp.Body.Close()

	var result struct {
		Results []struct {
			Geometry struct {
				Location Geocode `json:"location"`
			} `json:"geometry"`
		} `json:"results"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		er := fmt.Sprintf("Error decoding locations JSON: %v", err)
		dotlog.Error(er)
		return 0, 0
	}

	if len(result.Results) > 0 {
		location := result.Results[0].Geometry.Location
		Geo = []Geocode{location}
		return location.Lat, location.Lng
	}

	dotlog.Error("No results found in the geocoding response")
	return 0, 0
}

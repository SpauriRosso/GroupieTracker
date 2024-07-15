package Serveur

import (
	"encoding/json"
	"groupie-tracker/src/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

// structures
type Artist struct {
	Id      int      `json:"id"`
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

func GetId(URL string, name string) int {
	var artists []Artist
	resp, err := http.Get(URL)
	utils.CheckErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	utils.CheckErr(err)

	err = json.Unmarshal(body, &artists)
	utils.CheckErr(err)

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(name)) {
			//fmt.Println("ID:", artist.Id)
			return artist.Id
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(name)) {
				//fmt.Println("ID:", artist.Id)
				return artist.Id
			}
		}
	}
	return 0
}

//func RandArtistPage() {
//	var i = 1
//	var artist Artist
//	URL := "https://groupietrackers.herokuapp.com/api/artists/"
//
//	for i != 4 {
//		URL = "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(utils.GetRandArtist())
//		err := GetInfo(URL, &artist)
//		utils.CheckErr(err)
//		fmt.Printf("%+v\n\n \n", artist)
//		i++
//	}
//}

func GetInfo(URL string, data interface{}) error {
	resp, err := http.Get(URL)
	utils.CheckErr(err)
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
}

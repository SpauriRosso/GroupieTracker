package utils

import (
	"groupie-tracker/src/shared"
	"log"
	"math/rand"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GetRandArtist() string {
	RandID := rand.Intn(51) + 1
	return shared.ArtistURL + strconv.Itoa(RandID)
}

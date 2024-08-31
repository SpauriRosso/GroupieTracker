package utils

import (
	"fmt"
	"github.com/SpauriRosso/dotlog"
	"groupie-tracker-filters/src/shared"
	"log"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
)

func LogError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		te := fmt.Sprintf("%v %v:%v", err, file, line)
		dotlog.Error(te)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func GetRandArtist() string {
	RandID := rand.Intn(51) + 1
	return shared.ArtistURL + strconv.Itoa(RandID)
}

func ExactCity(city string) string {
	return strings.TrimSuffix(city, "-"+strings.Split(city, "-")[len(strings.Split(city, "-"))-1])
}

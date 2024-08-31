package main

import (
	Serveur "groupie-tracker-filters/src/server"
	"groupie-tracker-filters/src/shared"
	"time"

	"github.com/SpauriRosso/dotenv"

	"github.com/SpauriRosso/dotlog"
)

func main() {
	//server := Serveur.NewServer()
	dotenv.Define(".env")
	shared.WeatherAPI = dotenv.GetEnv("WeatherAPI")
	shared.MapAPI = dotenv.GetEnv("MAPS_API_KEY")

	//fmt.Println(Serveur.GetLoc("paris-france"))

	go func() {
		for {
			Serveur.FetchApi()
			time.Sleep(time.Hour)
			dotlog.Info("Fetched API")
		}
	}()

	dotlog.Info("Server started at http://localhost:5826")
	err := startServer(":5826")
	if err != nil {
		dotlog.CheckFuncErr(
			"Something unexpected happened!!",
			func() error { return err },
			"Never mind everything fine",
		)
	}
}

func startServer(addr string) error {
	server := Serveur.NewServer()
	return server.Start(addr)
}

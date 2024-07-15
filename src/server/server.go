package Serveur

import (
	"groupie-tracker/src/shared"
	"groupie-tracker/src/utils"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Server struct {
	mux           *http.ServeMux
	formSubmitted bool
	templateDir   string
	staticDir     string
}

func NewServer() *Server {
	return &Server{
		mux:         http.NewServeMux(),
		templateDir: filepath.Join(shared.BasePath, "assets/static/templates"),
		staticDir:   filepath.Join(shared.BasePath, "assets/static"),
	}
}

type Rad1 struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Rad2 struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Rad3 struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Art struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Date struct {
	Id    int
	Dates any `json:"datesLocations"`
}

// Declaring routes
func (s *Server) routes() {
	// Serve static files
	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(s.staticDir))))

	s.mux.HandleFunc("/", s.rootHandler)
	s.mux.HandleFunc("/home", s.homeHandler)
	s.mux.HandleFunc("/result", s.resultHandler)
	s.mux.HandleFunc("/500", s.errorHandler)
	s.mux.HandleFunc("/404", s.notFoundHandler)
	s.mux.HandleFunc("/405", s.notAllowedHandler)
}

func (s *Server) Start(addr string) error {
	s.routes()
	server := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}
	return server.ListenAndServe()
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	} else {
		s.notFoundHandler(w, r)
	}
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		shared.Prompt = r.FormValue("search")
		s.formSubmitted = true
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	t, err := template.ParseFiles(filepath.Join(s.templateDir, "index.html"))
	if err != nil {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return
	}

	var (
		Random1 Rad1
		Random2 Rad2
		Random3 Rad3
	)
	err = GetInfo(utils.GetRandArtist(), &Random1)
	utils.CheckErr(err)
	err = GetInfo(utils.GetRandArtist(), &Random2)
	utils.CheckErr(err)
	err = GetInfo(utils.GetRandArtist(), &Random3)
	utils.CheckErr(err)

	// assigning existing struct into a single struct
	var RandArt = struct {
		Random1 Rad1
		Random2 Rad2
		Random3 Rad3
	}{
		Random1: Random1,
		Random2: Random2,
		Random3: Random3,
	}

	err = t.Execute(w, RandArt)
	utils.CheckErr(err)
}

func (s *Server) resultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		shared.Prompt = r.FormValue("search")
		s.formSubmitted = true
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	if !s.formSubmitted {
		http.Redirect(w, r, "/405", http.StatusSeeOther)
		return
	}
	s.formSubmitted = false

	var Result Art
	var Concert Date
	shared.ArtistID = GetId(shared.URL, shared.Prompt)
	URL := shared.URL + "/" + strconv.Itoa(shared.ArtistID)
	CURL := shared.ConcertURL + "/" + strconv.Itoa(shared.ArtistID)

	if shared.ArtistID == 0 {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
		return // Stop forgetting about this return !!!!!!!!!!!!!!!!!!!
	}

	GetInfo(URL, &Result)
	GetInfo(CURL, &Concert)

	t, err := template.ParseFiles(filepath.Join(s.templateDir, "result.html"))
	if err != nil {
		log.Printf("Error parsing template result.html: %v", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	var Infos struct {
		Artist   Art
		Concerts Date
	} = struct {
		Artist   Art
		Concerts Date
	}{
		Artist:   Result,
		Concerts: Concert,
	}

	execErr := t.Execute(w, Infos)
	if execErr != nil {
		log.Fatalln(execErr)
	}
}

// Error Handlers
func (s *Server) errorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		shared.Prompt = r.FormValue("search")
		s.formSubmitted = true
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles(filepath.Join(s.templateDir, "error.html"))
	type Message struct {
		Code int
		Err  string
	}
	ErrMsg := Message{Code: http.StatusInternalServerError, Err: "Something unexpected happened"}
	w.WriteHeader(http.StatusInternalServerError)
	t.Execute(w, ErrMsg)
}

func (s *Server) notFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		shared.Prompt = r.FormValue("search")
		s.formSubmitted = true
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles(filepath.Join(s.templateDir, "error.html"))
	type Message struct {
		Code int
		Err  string
	}
	ErrMsg := Message{Code: http.StatusNotFound, Err: "Requested page does not exist"}
	w.WriteHeader(http.StatusNotFound)
	t.Execute(w, ErrMsg)
}

func (s *Server) notAllowedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		shared.Prompt = r.FormValue("search")
		s.formSubmitted = true
		http.Redirect(w, r, "/result", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles(filepath.Join(s.templateDir, "error.html"))
	type Message struct {
		Code int
		Err  string
	}
	ErrMsg := Message{Code: http.StatusMethodNotAllowed, Err: "Used method is not allowed"}
	w.WriteHeader(http.StatusMethodNotAllowed)
	t.Execute(w, ErrMsg)
}

// End of error Handler

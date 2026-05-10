package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	static_fs := os.DirFS("static")
	static_fs_handler := http.FileServer(http.FS(static_fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		static_fs_handler.ServeHTTP(w, r)
	})
	http.HandleFunc("/bracket", bracketHandler)
	http.HandleFunc("/event", eventHandler)

	fmt.Println("Hello RFGC")
	http.ListenAndServe(":8080", nil)
}

func bracketHandler(w http.ResponseWriter, r *http.Request) {
	bracket := r.URL.Query().Get("bracket")
	if bracket == "" {
		http.Error(w, "bracket required", 400)
		return
	}
	io.WriteString(w, "This is a bracket")
	io.WriteString(w, bracket)
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	event := r.URL.Query().Get("event")
	if event == "" {
		http.Error(w, "event required", 400)
		return
	}

	fmt.Println("Returning Event", event)

	var ev Event

	ev.Title = "King's Cup 42"
	ev.Description = "One of the Tournaments of All Time"
	ev.Brackets = make([]Bracket, 3)

	io.WriteString(w, ev.Title)
	io.WriteString(w, ev.Description)

	for i, _ := range ev.Brackets {
		fmt.Fprintf(w, `<div id="bracket-widget" hx-get="/bracket?bracket=%d" hx-trigger="load"></div>`, i)
	}

}

type Event struct {
	Id          uint64
	StartGGSlug string
	Title       string
	Description string
	Date        time.Time
	Link        string

	Brackets []Bracket
}

type Bracket struct {
	Id      uint64
	EventId uint64
	GameId  uint64

	Title       string
	Description string

	Matches []Match
}

type Match struct {
	Id        uint64
	StartGGId uint64

	Player1Id uint64
	Player2Id uint64

	// Pool number descends later into the tournament. So pool 0 is always the final pool that leads to top 8 and finals for the whole bracket.
	Pool int32

	// Round number is basically "Number of games away from the end of the pool", so round 0 is grand finals.
	// Negative numbers indicate loser's side and positive numbers indicate winner's side.
	Round int16

	Player1Score uint8
	Player2Score uint8
}

type Entrant struct {
	Id        uint64
	StartGGId uint64

	Seed int64
	Name string

	PlayerIds []uint64
}

type Player struct {
	Id        uint64
	StartGGId uint64

	GamerTag string
	Prefix   string
}

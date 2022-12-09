package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// PlayerStore stores score information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// Player stores a name with a number of wins.
type Player struct {
	Name string
	Wins int
}

// PlayerServer is a HTTP interface for player information.
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

const jsonContentType = "application/json"

// NewPlayerServer creates a PlayerServer with routing configured.
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	
	router := mux.NewRouter()
	router.HandleFunc("/league", p.leagueHandler).Methods("GET")
	router.HandleFunc("/players/{player}", p.processWin).Methods("POST")
	router.HandleFunc("/players/{player}", p.showScore).Methods("GET")

	p.Handler = router

	return p
}




func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}


func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	score := p.store.GetPlayerScore(vars["player"])

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p.store.RecordWin(vars["player"])
	w.WriteHeader(http.StatusAccepted)
}
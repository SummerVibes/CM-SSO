package test

import (
	"fmt"
	"net/http"
	"testing"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}
type PlayerServer struct {
	store PlayerStore
}
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func TestGETPlayers(t *testing.T) {
	//server := &PlayerServer{}
	//
	//t.Run("returns Pepper's score", func(t *testing.T) {
	//	request := newGetScoreRequest("Pepper")
	//	response := httptest.NewRecorder()
	//
	//	server.ServeHTTP(response, request)
	//
	//	assertResponseBody(t, response.Body.String(), "20")
	//})
	//
	//t.Run("returns Floyd's score", func(t *testing.T) {
	//	request := newGetScoreRequest("Floyd")
	//	response := httptest.NewRecorder()
	//
	//	server.ServeHTTP(response, request)
	//
	//	assertResponseBody(t, response.Body.String(), "10")
	//})
}

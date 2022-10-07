package server

import (
	"encoding/json"
	models "github.com/dghwood/battlesnake-go/models"
	snake "github.com/dghwood/battlesnake-go/snake"
	"log"
	"net/http"
	"os"
)

type Server struct {
	snake snake.Snake
}

func (s *Server) HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := s.snake.Info()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode info response, %s", err)
	}
}

func (s *Server) HandleStart(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode start json, %s", err)
		return
	}

	s.snake.Start(state)

	// Nothing to respond with here
}

func (s *Server) HandleMove(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode move json, %s", err)
		return
	}

	response := s.snake.Move(state)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("ERROR: Failed to encode move response, %s", err)
		return
	}
}

func (s *Server) HandleEnd(w http.ResponseWriter, r *http.Request) {
	state := models.GameState{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		log.Printf("ERROR: Failed to decode end json, %s", err)
		return
	}

	s.snake.End(state)

	// Nothing to respond with here
}

// Start Battlesnake Server
func RunServer(snek snake.Snake) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	s := Server{snek}

	http.HandleFunc("/", s.HandleIndex)
	http.HandleFunc("/start", s.HandleStart)
	http.HandleFunc("/move", s.HandleMove)
	http.HandleFunc("/end", s.HandleEnd)

	log.Printf("Running Battlesnake at http://0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

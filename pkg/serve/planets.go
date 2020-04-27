package serve

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) HandlePlanets(w http.ResponseWriter, r *http.Request) {
	if s.Game == nil {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJson(w, s.Planets)
	default:
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) ServePlanet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(strings.TrimPrefix(r.URL.Path, "/planets/"), 10, 10)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if id < 0 || int(id) >= len(s.Planets) {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		writeJson(w, s.Game.Planets[id])
	case http.MethodPut:
		putJson(w, r.Body, &s.Game.Planets[id])
	default:
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
	}
}

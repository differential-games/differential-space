package serve

import "net/http"

func (s *Server) HandlePlayers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		return
	}

	writeJson(w, s.Players)
}

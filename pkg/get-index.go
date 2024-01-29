package pkg

import (
	"net/http"
)

func (s *Server) GetIndex(w http.ResponseWriter, r *http.Request) {
	data, err := s.conf.Res.ReadFile("res/data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	s.templates.ExecuteTemplate(w, "index.html.tmpl", data)
}

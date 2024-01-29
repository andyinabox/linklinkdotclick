package pkg

import "net/http"

func (s *Server) GetMainJs(w http.ResponseWriter, r *http.Request) {
	js, err := s.conf.Res.ReadFile("res/main.js")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/javascript")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

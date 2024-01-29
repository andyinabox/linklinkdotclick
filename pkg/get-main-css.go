package pkg

import "net/http"

func (s *Server) GetMainCss(w http.ResponseWriter, r *http.Request) {
	css, err := s.conf.Res.ReadFile("res/public/main.css")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(http.StatusOK)
	w.Write(css)
}

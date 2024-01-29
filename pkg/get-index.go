package pkg

import (
	"encoding/json"
	"net/http"
)

type testData struct {
	Links []Link
}

func (s *Server) GetIndex(w http.ResponseWriter, r *http.Request) {

	// load data
	data, err := s.conf.Res.ReadFile("res/static/data.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	testData := testData{}
	err = json.Unmarshal(data, &testData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// send response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	// w.Write()
	s.templates.ExecuteTemplate(w, "index.html.tmpl", testData)
}

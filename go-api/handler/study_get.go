package handler

import (
	"encoding/json"
	"net/http"

	"github.com/marumaro-git/study-go-api/go-api/domain/model"
)

type Server struct {
	DBAccessor DBAccessor
}

type DBAccessor interface {
	GetRecord(r *http.Request) ([]model.Record, error)
}

func (s Server) StudyGetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	res, err := s.DBAccessor.GetRecord(r)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// resをjsonに変換にして返却
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "Failed to encode record: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

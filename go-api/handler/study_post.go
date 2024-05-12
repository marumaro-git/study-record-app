package handler

import (
	"fmt"
	"net/http"
)

func StudyPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Study Post"))
	fmt.Println("Study Post")
}

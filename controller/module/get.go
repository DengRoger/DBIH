package module

import (
	db "DBIH/controller/DB"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	ListContent []string `json:"listContent"`
}

func Getpage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	page, err := strconv.Atoi(vars["page"])
	if err != nil || page < 0 {
		http.Error(w, "invalid page number", http.StatusBadRequest)
		return
	}
	data := db.RGetPage(uid, page)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

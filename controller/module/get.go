package module

import (
	"encoding/json"
	"fmt"

	// "log"
	db "DBIH/controller/DB"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	ListContent []string `json:"listContent"`
}

// func Getpage(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	uid := vars["uid"]
// 	page, _ := strconv.Atoi(vars["page"])
// 	data := db.RGetPage(uid, page)
// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonData)
// 	w.WriteHeader(http.StatusOK)

// }

func Getpage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	page, err := strconv.Atoi(vars["page"])
	fmt.Println(uid, page)
	if err != nil {
		http.Error(w, "invalid page number", http.StatusBadRequest)
		return
	}
	data := db.RGetPage(uid, page)
	fmt.Print(data)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	w.WriteHeader(http.StatusOK)
}

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

func Getpage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
	} 
	// if err != nil {// if the page number is not an integer
	// 	http.Error(w, "Invalid page number", http.StatusBadRequest)
	// 	return
	// }
	tmp := db.RGetPage(uid, page)
	json.NewEncoder(w).Encode(tmp)
	fmt.Println("List sent to client")
}

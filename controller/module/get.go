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
	// type of page is int
	page, _ := strconv.Atoi(vars["page"])
	tmp := db.RGetPage(uid, page)
	fmt.Println(tmp)
	json.NewEncoder(w).Encode(tmp)
	fmt.Println(tmp)
	fmt.Println("List sent to client")
}

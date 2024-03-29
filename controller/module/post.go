package module

import (
	db "DBIH/controller/DB"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type newList struct {
	List []string `json:"listContent"`
}

func ModifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		panic("Method not allowed")
	}
	params := mux.Vars(r)
	var items newList
	err := json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(err.Error())
	}
	db.PInsertEntryList(params["uid"], items.List)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

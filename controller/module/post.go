package module
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "DBEE_HW/controller"
)

type ExChange struct {
	L int      `json:"l"`
	R int      `json:"r"`
	S []string `json:"s"`
}


func ModifyHandler(w http.ResponseWriter, r *http.Request) ExChange {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		panic("Method not allowed") 
	}
	var data ExChange
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
        panic(err.Error())
	}
	// fmt.Printf("L: %d\nR: %d\nS: %v\n", data.L, data.R, data.S)
    request := api.AddData(data.L, data.R, data.S) // int int []string
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonStr, err := json.Marshal(data)
	w.Write(jsonStr)
    return data
}

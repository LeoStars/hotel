package serverRoom

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func CreateRoom(w http.ResponseWriter, r *http.Request){
	fmt.Println(r, "\n", w)
	type Req struct {
		Room string
	}
	var req Req
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(req, r.Body)
}
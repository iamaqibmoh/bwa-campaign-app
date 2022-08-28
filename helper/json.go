package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteToResponseBody(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		log.Println(err.Error())
	}
}

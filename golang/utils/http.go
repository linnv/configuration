package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, obj interface{}) {
	var (
		body []byte
		err  error
	)
	switch obj := obj.(type) {
	case error:
		w.WriteHeader(400)
		body, _ = json.Marshal(map[string]string{
			"error": obj.Error(),
		})
		log.Println(obj)
	case []byte:
		body = obj
	default:
		body, err = json.Marshal(map[string]interface{}{
			"result": obj,
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
	// json
	w.Header().Set("Content-Type", "application/json")
	n, err := w.Write(body)
	if err == nil && n < len(body) {
		err = fmt.Errorf("short written")
	}
	if err != nil {
		log.Println("wrote back to client: ", err)
	}
}

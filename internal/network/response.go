package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GenerateJsonResponse(w http.ResponseWriter, data interface{}, headers map[string]string, timeout int64) {
	timeoutDuration := time.Duration(timeout)
	time.Sleep(timeoutDuration * time.Millisecond)

	jsonResp, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	if headers != nil {
		for header, value := range headers {
			w.Header().Set(header, value)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		fmt.Println(err)
		return
	}
}

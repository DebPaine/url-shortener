package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

var urlMap map[string]string

type RequestBody struct {
	Url string `json:"url"`
	Abc string
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println(reflect.TypeOf(requestBody.Abc), reflect.TypeOf(requestBody.Url))
	fmt.Fprint(w, requestBody)
}

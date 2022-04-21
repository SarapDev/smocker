package network

import (
	"encoding/json"
	"fmt"
	"git.chocodev.kz/rahmet/tools/mocker/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"path"
)

func GetMockData(w http.ResponseWriter, r *http.Request) {
	objIdHex := path.Base(r.URL.Path)
	md := models.MockData{}
	md = md.GetById(objIdHex).(models.MockData)

	GenerateJsonResponse(w, md.Body, md.Header, md.Timeout)
}

func GenerateUrlWithData(w http.ResponseWriter, r *http.Request) {
	var md models.MockData

	err := json.NewDecoder(r.Body).Decode(&md)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rep, err := md.Save()

	if err != nil {
		fmt.Println(err)
		return
	}

	resp := make(map[string]interface{})
	resp["url"] = r.Host + "/mocker/" + rep.InsertedID.(primitive.ObjectID).Hex()

	GenerateJsonResponse(w, resp, nil, 0)
}

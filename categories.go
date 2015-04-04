package main

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func AddCategorieHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	categorie := params["name"]
	err := CategoriesCollection.Insert(bson.M{"name": categorie})
	if err != nil {
		writeError(w, 500, "Error while finding categories", err)
		return
	}

	fmt.Println("Adding categorie: " + categorie)
	w.WriteHeader(http.StatusCreated)
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var categories []struct {
		Name string `json:"name"`
	}
	err := CategoriesCollection.Find(bson.M{}).All(&categories)
	if err != nil {
		writeError(w, 500, "Error while finding categories", err)
		return
	}

	result, err := marshalJson(categories)
	if err != nil {
		writeError(w, 500, "Error while marshalling categories", err)
		return
	}

	fmt.Println(string(result))
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

}

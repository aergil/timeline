package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func AddCategorieHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error while reading request body", err)
	}

	fmt.Println("Adding categorie requested ", string(body))
	newCategorie := &Categorie{}
	err = json.Unmarshal(body, newCategorie)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Error : the categorie is not valid", err)
	}

	if newCategorie.Id.Hex() != "" {
		err = CategorieCollection.UpdateId(newCategorie.Id, newCategorie)
	} else {
		err = CategorieCollection.Insert(newCategorie)
	}

	if err != nil {
		writeError(w, 500, "Error while inserting categorie", err)
		return
	}

	fmt.Println("Adding categorie: ", newCategorie)
	w.WriteHeader(http.StatusCreated)
}

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	categories := []Categorie{}
	err := CategorieCollection.Find(bson.M{}).All(&categories)
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

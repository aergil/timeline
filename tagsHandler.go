package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func AddTagHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error while reading request body", err)
	}

	fmt.Println("Adding tags requested ", string(body))
	newTags := &Tag{}
	err = json.Unmarshal(body, newTags)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Error : the tags is not valid", err)
	}

	if newTags.Id.Hex() != "" {
		err = TagCollection.UpdateId(newTags.Id, newTags)
	} else {
		err = TagCollection.Insert(newTags)
	}

	if err != nil {
		writeError(w, 500, "Error while inserting tags", err)
		return
	}

	fmt.Println("Adding tags: ", newTags)
	w.WriteHeader(http.StatusCreated)
}

func GetTagsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	tags := []Tag{}
	err := TagCollection.Find(bson.M{}).All(&tags)
	if err != nil {
		writeError(w, 500, "Error while finding tags", err)
		return
	}

	result, err := marshalJson(tags)
	if err != nil {
		writeError(w, 500, "Error while marshalling tagss", err)
		return
	}

	fmt.Println(string(result))
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

}

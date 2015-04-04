package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

// ceci est une fonction
func AddEventsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Error while reading request body", err)
	}

	fmt.Println("Adding event requested ", string(body))
	newEvent := &Event{}
	err = json.Unmarshal(body, newEvent)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Error : the event is not valid", err)
	}

	for _, value := range newEvent.Categories {
		count, err := CategoriesCollection.Find(bson.M{"name": value}).Count()
		if err != nil {
			writeError(w, 500, "Error while inserting event", err)
			return
		}
		if count == 0 {
			CategoriesCollection.Insert(bson.M{"name": value})
		}
	}

	if newEvent.Id.Hex() != "" {
		fmt.Println("Update")
		err = EventCollection.UpdateId(newEvent.Id, newEvent)
	} else {
		err = EventCollection.Insert(newEvent)
	}

	if err != nil {
		writeError(w, 500, "Error while inserting event", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func SearchEventsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	name := params["name"]
	fmt.Println("name: ", name)
	events := []Event{}
	err := EventCollection.Find(bson.M{"name": bson.RegEx{".*" + name + ".*", ""}}).All(&events)
	if err != nil {
		writeError(w, 500, "Error while searching events", err)
	}
	eventsJson, _ := json.Marshal(events)
	fmt.Println("eventsJson: ", string(eventsJson))

	w.Write(eventsJson)
}

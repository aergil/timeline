package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

	err = EventCollection.Insert(newEvent)
	if err != nil {
		writeError(w, 500, "Error while inserting event", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

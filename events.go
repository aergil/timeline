package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Events struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string        `json:"name"`
	Start     int           `json:"start"`
	End       int           `json:"end"`
	Ponctuels []Ponctuel    `json:"ponctuels"`
}
type Ponctuel struct {
	Date        int    `json:"date"`
	Description string `json:"description"`
}

func EventsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Println("Events requested")
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		fmt.Println("Error while creating session", err)
		http.Error(w, "Error while creating session", 500)
		return
	}

	c := session.DB("timeline").C("events")
	var events []Events
	err = c.Find(nil).All(&events)
	if err != nil {
		fmt.Println("Error while finding events", err)
		http.Error(w, "Error while finding events", 500)
		return
	}

	result, err := json.Marshal(events)
	if err != nil {
		fmt.Println("Error while marshalling events", err)
		http.Error(w, "Error while finding events", 500)
		return
	}

	fmt.Println(string(result))
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

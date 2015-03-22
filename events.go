package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	start, err1 := strconv.Atoi(params["start"])
	end, err2 := strconv.Atoi(params["end"])
	if err1 != nil || err2 != nil {
		writeError(w, 400, "Error : one of the date are not number: "+params["start"]+params["end"], nil)
	}

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		writeError(w, 500, "Error while creating session", err)
		return
	}

	c := session.DB("timeline").C("events")
	var events []Events
	err = c.Find(bson.M{"end": bson.M{"$gt": start}, "start": bson.M{"$lt": end}}).All(&events)
	if err != nil {
		writeError(w, 500, "Error while finding events", err)
		return
	}

	result, err := marshalJson(events)
	if err != nil {
		writeError(w, 500, "Error while marshalling events", err)
		return
	}

	fmt.Println(string(result))
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func writeError(w http.ResponseWriter, code int, message string, err error) {
	fmt.Println(message, err)
	http.Error(w, message, code)
}

func marshalJson(i interface{}) ([]byte, error) {
	result, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	if string(result) == "null" {
		result = []byte("[]")
	}
	return result, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `json:"name"`
	Start      int           `json:"start"`
	End        int           `json:"end"`
	Ponctuels  []Ponctuel    `json:"ponctuels"`
	Categories []string      `json:"categories"`
}
type Ponctuel struct {
	Date        int    `json:"date"`
	Description string `json:"description"`
}

func EventsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Println("Events requested")
	c := params["categories"]
	categories := strings.Split(c, ",")
	fmt.Println("categories: ", categories)
	start, err1 := strconv.Atoi(params["start"])
	end, err2 := strconv.Atoi(params["end"])
	if err1 != nil || err2 != nil {
		writeError(w, 400, "Error : one of the date are not number: "+params["start"]+params["end"], nil)
	}

	var events []Event
	var err error
	if len(categories) > 0 {
		err = EventCollection.Find(bson.M{"end": bson.M{"$gt": start}, "start": bson.M{"$lt": end}, "categories": bson.M{"$in": categories}}).All(&events)
	} else {
		err = EventCollection.Find(bson.M{"end": bson.M{"$gt": start}, "start": bson.M{"$lt": end}}).All(&events)
	}
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

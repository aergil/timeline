package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
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

	for _, value := range newEvent.Tags {
		count, err := TagCollection.Find(bson.M{"name": value}).Count()
		if err != nil {
			writeError(w, 500, "Error while inserting event", err)
			return
		}
		if count == 0 {
			TagCollection.Insert(bson.M{"name": value})
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

	eventsJson, err := marshalJson(events)
	if err != nil {
		writeError(w, 500, "Error while marshalling events", err)
		return
	}

	fmt.Println("eventsJson: ", string(eventsJson))
	w.Write(eventsJson)
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Println("Events requested")
	c := params["tags"]
	tags := strings.Split(c, ",")
	fmt.Println("tags: ", tags)
	start, err1 := strconv.Atoi(params["start"])
	end, err2 := strconv.Atoi(params["end"])
	if err1 != nil || err2 != nil {
		writeError(w, 400, "Error : one of the date are not number: "+params["start"]+params["end"], nil)
	}

	var events []Event
	var err error
	if len(tags) > 0 && tags[0] != "" {
		fmt.Println("Get event with tags", len(tags))
		err = EventCollection.Find(bson.M{"end": bson.M{"$gt": start}, "start": bson.M{"$lt": end}, "tags": bson.M{"$in": tags}}).All(&events)
	} else {
		fmt.Println("Get event without tags")
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

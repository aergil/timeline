package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestAddEventHandlerBadContent(t *testing.T) {
	Init("127.0.0.1", "timeline_tests", "events", "categories")
	recorder := httptest.NewRecorder()

	r := strings.NewReader("toto")
	req, err := http.NewRequest("POST", "http://localhost", r)
	if err != nil {
		log.Fatal(err)
	}
	AddEventsHandler(recorder, req, nil)

	if recorder.Code != 400 {
		t.Error("Code Should be 400 but ", recorder.Code)
	}
}

func TestAddEventHandler(t *testing.T) {
	Init("127.0.0.1", "timeline_tests", "events", "categories")
	recorder := httptest.NewRecorder()

	eventjson := `{"name":"event","start":2014,"end":2014,"ponctuels":[]}`
	r := strings.NewReader(string(eventjson))
	req, err := http.NewRequest("POST", "http://localhost", r)
	if err != nil {
		log.Fatal(err)
	}
	AddEventsHandler(recorder, req, nil)

	if recorder.Code != 201 {
		t.Error("Code Should be 200 but ", recorder.Code)
	}
}

func TestSearchEventHandler(t *testing.T) {
	Init("127.0.0.1", "timeline_tests", "events", "categories")
	EventCollection.RemoveAll(bson.M{})

	recorder := httptest.NewRecorder()

	eventjson := `{"name":"event","start":2014,"end":2014,"ponctuels":[]}`
	r := strings.NewReader(string(eventjson))
	req, _ := http.NewRequest("POST", "http://localhost", r)
	AddEventsHandler(recorder, req, nil)

	recorder2 := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "http://localhost", nil)
	SearchEventsHandler(recorder2, reqGet, map[string]string{"name": "event"})

	if recorder2.Code != 200 {
		t.Error("Code should be 200 but", recorder2.Code)
	}

	body, _ := ioutil.ReadAll(recorder2.Body)
	events := []Event{}
	err := json.Unmarshal(body, &events)
	if err != nil {
		t.Error("Error :", err)
	}
	if len(events) != 1 {
		t.Error("Error. got: ", string(body), len(events))
	}

}

func TestEventHandlerWithCategories(t *testing.T) {
	Init("127.0.0.1", "timeline_tests", "events", "categories")
	EventCollection.RemoveAll(bson.M{})

	recorder := httptest.NewRecorder()

	eventjson := `{"name":"event1","start":2014,"end":2014,"ponctuels":[], "categories": ["Math","Philo"]}`
	req, _ := http.NewRequest("POST", "http://localhost", strings.NewReader(string(eventjson)))
	AddEventsHandler(recorder, req, nil)

	eventjson = `{"name":"event2","start":2014,"end":2014,"ponctuels":[], "categories": ["Math"]}`
	req, _ = http.NewRequest("POST", "http://localhost", strings.NewReader(string(eventjson)))
	AddEventsHandler(recorder, req, nil)

	eventjson = `{"name":"event3","start":2014,"end":2014,"ponctuels":[], "categories": []}`
	req, _ = http.NewRequest("POST", "http://localhost", strings.NewReader(string(eventjson)))
	AddEventsHandler(recorder, req, nil)

	recorder2 := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "http://localhost", nil)
	EventsHandler(recorder2, reqGet, map[string]string{"start": "0", "end": "2200", "categories": "Math,Philo"})

	if recorder2.Code != 200 {
		t.Error("Code should be 200 but", recorder2.Code)
	}

	body, _ := ioutil.ReadAll(recorder2.Body)
	events := []Event{}
	err := json.Unmarshal(body, &events)
	if err != nil {
		t.Error("Error :", err)
	}
	if len(events) != 2 {
		t.Error("Error. got: ", string(body), len(events))
	}

}

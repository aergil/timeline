package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddEventHandlerBadContent(t *testing.T) {
	Init("127.0.0.1", "timeline_tests", "events")
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
	Init("127.0.0.1", "timeline_tests", "events")
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

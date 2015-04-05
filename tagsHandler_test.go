package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestGetTags(t *testing.T) {
	Init("127.0.0.1", "timeline_tests", "events", "tags")
	TagCollection.RemoveAll(bson.M{})

	recorder := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "http://localhost", strings.NewReader(string(`{"name":"Philosophie"}`)))
	AddTagHandler(recorder, req1, nil)
	req2, _ := http.NewRequest("POST", "http://localhost", strings.NewReader(string(`{"name":"Mathematique"}`)))
	AddTagHandler(recorder, req2, nil)

	// query nil
	recorder2 := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "http://localhost", nil)
	GetTagsHandler(recorder2, reqGet, nil)

	if recorder2.Code != 200 {
		t.Error("Code should be 200 but", recorder2.Code)
	}

	body, _ := ioutil.ReadAll(recorder2.Body)
	tags := []Tag{}
	err := json.Unmarshal(body, &tags)
	if err != nil {
		t.Error("Error :", err)
	}
	if len(tags) != 2 {
		t.Error("Error. got: ", string(body), len(tags))
	}

	// query regex
	recorder3 := httptest.NewRecorder()
	reqGet2, _ := http.NewRequest("GET", "http://localhost", nil)
	GetTagsHandler(recorder3, reqGet2, map[string]string{"query": "oso"})

	if recorder3.Code != 200 {
		t.Error("Code should be 200 but", recorder3.Code)
	}

	body2, _ := ioutil.ReadAll(recorder3.Body)
	tags = []Tag{}
	err = json.Unmarshal(body2, &tags)
	if err != nil {
		t.Error("Error :", err)
	}
	if len(tags) != 1 {
		t.Error("Error. got: ", string(body2), len(tags))
	}

}

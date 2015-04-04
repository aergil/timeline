package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestGetCategories(t *testing.T) {
	Init("127.0.0.1", "timeline_tests", "events", "categories")
	CategoriesCollection.RemoveAll(bson.M{})

	recorder := httptest.NewRecorder()

	req1, _ := http.NewRequest("POST", "http://localhost", strings.NewReader(string(`{"name":"Philosophie"}`)))
	AddCategorieHandler(recorder, req1, nil)
	req2, _ := http.NewRequest("POST", "http://localhost", strings.NewReader(string(`{"name":"Mathematique"}`)))
	AddCategorieHandler(recorder, req2, nil)

	recorder2 := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "http://localhost", nil)
	GetCategoriesHandler(recorder2, reqGet, nil)

	if recorder2.Code != 200 {
		t.Error("Code should be 200 but", recorder2.Code)
	}

	body, _ := ioutil.ReadAll(recorder2.Body)
	categories := []struct {
		Name string `json:"name"`
	}{}
	fmt.Println(string(body))
	err := json.Unmarshal(body, &categories)
	if err != nil {
		t.Error("Error :", err)
	}
	if len(categories) != 2 {
		t.Error("Error. got: ", string(body), len(categories))
	}

}

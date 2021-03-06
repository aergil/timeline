package main

import "gopkg.in/mgo.v2/bson"

type Event struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `json:"name"`
	Start     int           `json:"start"`
	End       int           `json:"end"`
	Ponctuels []Ponctuel    `json:"ponctuels"`
	Tags      []string      `json:"tags"`
}
type Ponctuel struct {
	Date        int    `json:"date"`
	Description string `json:"description"`
}

type Tag struct {
	Id   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `json:"name"`
}

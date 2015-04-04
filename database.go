package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var EventCollection *mgo.Collection
var CategoriesCollection *mgo.Collection

func Init(host, databaseName, eventCollectionName, categoriesCollectionName string) {
	session, err := mgo.Dial(host)
	if err != nil {
		fmt.Println("Error while creating session", err)
		panic(err)
	}
	EventCollection = session.DB(databaseName).C(eventCollectionName)
	CategoriesCollection = session.DB(databaseName).C(categoriesCollectionName)
}

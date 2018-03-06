package main

import (
	"fmt"
  "log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

var _ ReportDatabase = &mongoDB{}

func newMongoDB(addr string) (ReportDatabase, error) {
	conn, err := mgo.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}

	return &mongoDB{
		conn: conn,
		c:    conn.DB("yoga").C("teacher"),
	}, nil
}

func (db *mongoDB) Close() {
	db.conn.Close()
}

func (db *mongoDB) ListReports() ([]*Report, error) {
  var result []*Report
  if err := db.c.Find(nil).Sort("Name").All(&result); err != nil {
    return nil, err
  }
  log.Print(result)
  log.Print(result[0].Name)
  return result, nil
}




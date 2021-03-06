package main

import (
    "net/http"
    "log"
    "html/template"

    "gopkg.in/mgo.v2/bson"
)

var (
    DB      ReportDatabase
)

type Report struct {
    ID      bson.ObjectId `bson:"_id,omitempty"`
    Name     string  `json:"Name,omitempty" bson:"Name,omitempty"`
	  PreparedBy     string  `json:"PreparedBy,omitempty" bson:"PreparedBy,omitempty"`
    System     string  `json:"System,omitempty" bson:"System,omitempty"`
    ProgramLanguage     string  `json:"ProgramLanguage,omitempty" bson:"ProgramLanguage,omitempty"`
    OS     string  `json:"OS,omitempty" bson:"OS,omitempty"`
    Comments     string  `json:"Comments,omitempty" bson:"Comments,omitempty"`
}

type ReportDatabase interface {
    ListReports() ([]*Report, error)
    Close()
}

var templates = template.Must(template.ParseGlob("templates/*"))

func indexPage(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "indexPage", nil)
  if err != nil {
    log.Print("template error: ", err)
  }
}

func reportListPage(w http.ResponseWriter, r *http.Request) {
    reports, err := DB.ListReports()
  	if err != nil {
  		log.Print(err, "could not list reports: %v", err)
  	}
    log.Print("Reports: ", reports)
    err2 := templates.ExecuteTemplate(w, "reportsListPage", reports)
    if err != nil {
      log.Print("template error: ", err2)
    }
}

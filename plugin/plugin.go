package main

import (
	"errors"
	"github.com/rwynn/monstache/monstachemap"
	"time"
)

var WEEKDAYS = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
var MONTHS = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "July", "Aug", "Sep", "Oct", "Nov", "Dec"}
var HOURS = []string{"12am", "1am", "2am", "3am", "4am", "5am", "6am", "7am", "8am", "9am", "10am", "11am", "12pm", "1pm", "2pm", "3pm", "4pm", "5pm", "6pm", "7pm", "8pm", "9pm", "10pm", "11pm"}

// a plugin to convert crime data
func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	doc := input.Document
	if doc["ID"] == "ID" {
		// header line, skip it
		output = &monstachemap.MapperPluginOutput{Skip: true}
		return
	}
	lat, lon := doc["Latitude"], doc["Longitude"]
	if lat != nil && lon != nil {
		latF, ok := lat.(float64)
		if !ok {
			return nil, errors.New("Latitude was not a float64")
		}
		lonF, ok := lon.(float64)
		if !ok {
			return nil, errors.New("Longitude was not a float64")
		}
		doc["Loc"] = map[string]float64{
			"lat": latF,
			"lon": lonF,
		}
	}
	ts := doc["Date"]
	if ts != nil {
		tst, ok := ts.(time.Time)
		if !ok {
			return nil, errors.New("Date was not a time field")
		}
		doc["Hour"] = HOURS[tst.Hour()]
		doc["Weekday"] = WEEKDAYS[tst.Weekday()]
		doc["Month"] = MONTHS[tst.Month()-1]
	}
	delete(doc, "Location")
	delete(doc, "Latitude")
	delete(doc, "Longitude")
	output = &monstachemap.MapperPluginOutput{Document: doc}
	return
}

package main

import (
	"errors"
	"github.com/rwynn/monstache/monstachemap"
	"time"
)

var WEEKDAYS = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
var MONTHS = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "July", "Aug", "Sep", "Oct", "Nov", "Dec"}
var HOURS = []string{"12am", "1am", "2am", "3am", "4am", "5am", "6am", "7am", "8am", "9am", "10am", "11am", "12pm", "1pm", "2pm", "3pm", "4pm", "5pm", "6pm", "7pm", "8pm", "9pm", "10pm", "11pm"}

// based on https://en.wikipedia.org/wiki/Community_areas_in_Chicago
var COM = map[string]string{
	"8":  "Near North Side",
	"32": "Loop",
	"33": "Near South Side",
	"5":  "North Center",
	"6":  "Lake View",
	"7":  "Lincoln Park",
	"21": "Avondale",
	"22": "Logan Square",
	"1":  "Rogers Park",
	"2":  "West Ridge",
	"3":  "Uptown",
	"4":  "Lincoln Square",
	"9":  "Edison Park",
	"10": "Norwood Park",
	"11": "Jefferson Park",
	"12": "Forest Glen",
	"13": "North Park",
	"14": "Albany Park",
	"76": "O'Hare",
	"77": "Edgewater",
	"15": "Portage Park",
	"16": "Irving Park",
	"17": "Dunning",
	"18": "Montclare",
	"19": "Belmont Cragin",
	"20": "Hermosa",
	"23": "Humboldt Park",
	"24": "West Town",
	"25": "Austin",
	"26": "West Garfield Park",
	"27": "East Garfield Park",
	"28": "Near West Side",
	"29": "North Lawndale",
	"30": "South Lawndate",
	"31": "Lower West Side",
	"34": "Armour Square",
	"35": "Douglas",
	"36": "Oakland",
	"37": "Fuller Park",
	"38": "Grand Boulevard",
	"39": "Kenwood",
	"40": "Washington Park",
	"41": "Hyde Park",
	"42": "Woodlawn",
	"43": "South Shore",
	"60": "Bridgeport",
	"69": "Greater Grand Crossing",
	"56": "Garfield Ridge",
	"57": "Archer Heights",
	"58": "Brighton Park",
	"59": "McKinley Park",
	"61": "New City",
	"62": "West Elsdon",
	"63": "Gage Park",
	"64": "Clearing",
	"65": "West Lawn",
	"66": "Chicago Lawn",
	"67": "West Englewood",
	"68": "Englewood",
	"44": "Chatham",
	"45": "Avalon Park",
	"46": "South Chicago",
	"47": "Burnside",
	"48": "Calumet Heights",
	"49": "Roseland",
	"50": "Pullman",
	"51": "South Deering",
	"52": "East Side",
	"53": "West Pullman",
	"54": "Riverdale",
	"55": "Hegewisch",
	"70": "Ashburn",
	"71": "Auburn Gresham",
	"72": "Beverly",
	"73": "Washington Heights",
	"74": "Mount Greenwood",
	"75": "Morgan Park",
}

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
	ca := doc["Community Area"]
	if ca != nil {
		cad := COM[ca.(string)]
		if cad != "" {
			doc["Community Area Name"] = cad
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

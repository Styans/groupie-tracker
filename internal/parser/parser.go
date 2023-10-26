package parser

import (
	"encoding/json"
	"net/http"
	"time"
)

func Parse() ([]Artists, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	loc, err := client.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	date, err := client.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	rel, err := client.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	defer loc.Body.Close()
	defer date.Body.Close()
	defer rel.Body.Close()

	var artists []Artists
	json.NewDecoder(resp.Body).Decode(&artists)
	var (
		locdates IndexRelations
		location IndexLocations
		dates    IndexConcert
	)
	err = json.NewDecoder(loc.Body).Decode(&location)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(date.Body).Decode(&dates)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(rel.Body).Decode(&locdates)
	if err != nil {
		return nil, err
	}
	for i := range artists {
		artists[i].ConcertDates = dates.ConcertsDates[i]
		artists[i].Locations = location.Locations[i]
		artists[i].Relations = locdates.Relations[i]
	}
	return artists, nil
}

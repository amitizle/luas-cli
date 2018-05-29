package luas

import (
	"encoding/xml"
	"fmt"
	"github.com/amitizle/go-luas/internal/http_client"
	"strings"
)

type Stop struct {
	Name        string    `json:"name"`
	NameAbv     string    `json:"name_abv"`
	Line        string    `json:"line"`
	Coordinates []float64 `json:"coordinates"`
}

type StopInfo struct {
	Directions []Direction `xml:"direction"`
	Message    string      `xml:"message"`
	Created    string      `xml:"created,attr"`
	Stop       string      `xml:"stop,attr"`
	StopAbv    string      `xml:"stopAbv,attr"` // Short name for the stop (3 letters)
}

type Direction struct {
	Name  string `xml:"name,attr"` // Inbound/Outbound
	Trams []Tram `xml:"tram"`
}

type Tram struct {
	DueMins     string `xml:"dueMins,attr"` // string because it can be "DUE"
	Destination string `xml:"destination,attr"`
}

func parseLuasResponse(xmlString []byte) (*StopInfo, error) {
	var stopInfo StopInfo
	err := xml.Unmarshal(xmlString, &stopInfo)
	if err != nil {
		return &StopInfo{}, err
	}
	return &stopInfo, nil
}

func AllStops() []*Stop {
	return allStops
}

func GetStop(stopName string) (*Stop, error) {
	stopNameUpcase := strings.ToUpper(stopName)
	allStops := AllStops()
	for _, stop := range allStops {
		if strings.ToUpper(stop.NameAbv) == stopNameUpcase ||
			strings.ToUpper(stop.Name) == stopNameUpcase {
			return stop, nil
		}
	}
	return &Stop{}, fmt.Errorf("no such stop %v", stopNameUpcase)
}

// TODO integration test
func (stop *Stop) Forecast() (*StopInfo, error) {
	var stopInfo StopInfo
	httpClient, err := luas_http_client.NewClient("")
	if err != nil {
		return &stopInfo, err
	}
	resp, err := httpClient.GetForecast(stop.NameAbv)
	if err != nil {
		return &stopInfo, err
	}
	err = xml.Unmarshal(resp, &stopInfo)

	return &stopInfo, err
}

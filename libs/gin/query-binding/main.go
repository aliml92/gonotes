package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


var flights = []Flight{
	{FlightID: "flight0", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight1", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight2", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight3", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight4", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight5", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight6", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight7", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight8", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
	{FlightID: "flight9", DepartureAirportCode: "SFO", ArrivalAirportCode: "LAX"},
}


type Flight struct {
    FlightID              string    `json:"flight_id"`
    DepartureAirportCode  string    `json:"departure_airport_code"`
    ArrivalAirportCode    string    `json:"arrival_airport_code"`
    // FlightNumber          string    `json:"flight_number"`
    // AirlineName           string    `json:"airline_name"`
    // AirlineID             string    `json:"airline_id"`	
    // DepartureTime         time.Time `json:"departure_time"`
    // ArrivalTime           time.Time `json:"arrival_time"`
    // Status                string    `json:"status"`
}


type flightQuery struct {
	From	string	`form:"from"`
	To  	string	`form:"to"`
	Offset	int    	`form:"offset"`
	Limit	int 	`form:"limit"`
}


func FetchFlights(c *gin.Context) {
	var query flightQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("query: %+v", query)
	resp := make([]Flight, 0)

	// 1.filter flights
	for _, flight := range flights {
		if flight.DepartureAirportCode == query.From && flight.ArrivalAirportCode == query.To {
			resp = append(resp, flight)
		}
	}
	// 2. apply offset and limit
	if query.Offset > 0 {
		if query.Offset > len(resp) {
			resp = make([]Flight, 0)
		} else {
			resp = resp[query.Offset:]
		}
	}
	if query.Limit > 0 {
		if query.Limit > len(resp) {
			resp = resp[:]
		} else {
			resp = resp[:query.Limit]
		}
	}
	

	c.JSON(http.StatusOK, resp)
}


func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/flights", FetchFlights)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
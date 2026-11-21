// Package weather provides tools to check current weather.
package weather

// CurrentCondition tells the current weather report using CurrentLocation.
var CurrentCondition string

// CurrentLocation represents the city for which the weather condition is reported.
var CurrentLocation string

// Forecast returns the current city and the current weather condition in that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
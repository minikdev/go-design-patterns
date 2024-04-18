package main

import (
	"example.com/observer-pattern/display"
	"example.com/observer-pattern/subject"
)

func main() {
	weatherData := subject.NewWeatherData()
	display.NewCurrentConditionsDisplay(weatherData)
	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 29.1)
}

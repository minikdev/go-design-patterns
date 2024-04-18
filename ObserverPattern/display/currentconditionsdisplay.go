package display

import (
	"fmt"

	"example.com/observer-pattern/subject"
)

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64
	WeatherData subject.ISubject
}

func NewCurrentConditionsDisplay(weatherData subject.ISubject) *CurrentConditionsDisplay {
	d := &CurrentConditionsDisplay{
		WeatherData: weatherData,
	}
	weatherData.RegisterObserver(d)
	return d
}
func (c *CurrentConditionsDisplay) Update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.pressure = pressure
	c.Display()
}
func (c *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2fF degrees, %.2f%% humidity and %.2f pressure\n", c.temperature, c.humidity, c.pressure)
}

func (c *CurrentConditionsDisplay) GetTemperature() float64 {
	return c.temperature
}
func (c *CurrentConditionsDisplay) GetHumidity() float64 {
	return c.humidity
}
func (c *CurrentConditionsDisplay) GetPressure() float64 {
	return c.pressure
}

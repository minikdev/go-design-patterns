package subject

import "example.com/observer-pattern/observer"

type WeatherData struct {
	observers []observer.IObserver
	temp      float64
	humidity  float64
	pressure  float64
}

func NewWeatherData() *WeatherData {
	return &WeatherData{}
}

func (w *WeatherData) RegisterObserver(o observer.IObserver) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o observer.IObserver) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temp, w.humidity, w.pressure)
	}
}

func (w *WeatherData) MeasurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temp float64, humidity float64, pressure float64) {
	w.temp = temp
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()
}

func (w *WeatherData) GetTemp() float64 {
	return w.temp
}
func (w *WeatherData) GetHumidity() float64 {
	return w.humidity
}
func (w *WeatherData) GetPressure() float64 {
	return w.pressure
}

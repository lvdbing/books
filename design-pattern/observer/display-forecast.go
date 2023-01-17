package main

import "fmt"

type ForecastDisplay struct {
	currentPressure float64
	lastPressure    float64

	weatherData Subject
}

func NewForecastDisplay(weatherData Subject) *ForecastDisplay {
	f := &ForecastDisplay{}
	f.weatherData = weatherData
	f.weatherData.RegisterObserver(f)
	return f
}

func (f *ForecastDisplay) Update(temperature, humidity, pressure float64) {
	f.lastPressure = f.currentPressure
	f.currentPressure = pressure

	f.Display()
}

func (f *ForecastDisplay) Display() {
	fmt.Printf("Forecast display: current pressure is %.2f and last pressure is %.2f\n", f.currentPressure, f.lastPressure)
}

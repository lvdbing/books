package main

import "fmt"

type StatisticsDisplay struct {
	temperature float64
	humidity    float64
	pressure    float64

	weatherData Subject
}

func NewStatisticsDisplay(weatherData Subject) *StatisticsDisplay {
	s := &StatisticsDisplay{}
	s.weatherData = weatherData
	s.weatherData.RegisterObserver(s)
	return s
}

func (s *StatisticsDisplay) Update(temperature, humidity, pressure float64) {
	s.temperature = temperature
	s.humidity = humidity
	s.pressure = pressure

	s.Display()
}

func (s *StatisticsDisplay) Display() {
	fmt.Printf("Statistics display: temperature is %.2f, humidity is %.2f%%, pressure is %.2f\n", s.temperature, s.humidity, s.pressure)
}

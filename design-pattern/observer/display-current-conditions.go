package main

import "fmt"

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
	c := &CurrentConditionsDisplay{}
	c.weatherData = weatherData
	c.weatherData.RegisterObserver(c)
	return c
}

func (c *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	c.temperature = temperature
	c.humidity = humidity

	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: %.2fF degrees and %.2f%% humidity\n", c.temperature, c.humidity)
}

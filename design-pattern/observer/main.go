package main

import "fmt"

func main() {
	weatherData := NewWeatherData()
	current := NewCurrentConditionsDisplay(weatherData)
	forecast := NewForecastDisplay(weatherData)
	statistics := NewStatisticsDisplay(weatherData)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(90, 75, 40.4)
	weatherData.SetMeasurements(70, 55, 20.4)

	fmt.Println()

	fmt.Printf("current: %v\n", current)
	weatherData.RemoveObserver(current)
	fmt.Printf("forecast: %v\n", forecast)
	fmt.Printf("statistics: %v\n", statistics)
	weatherData.RemoveObserver(statistics)
	weatherData.SetMeasurements(10, 65, 30.4)
	weatherData.SetMeasurements(20, 75, 40.4)
	weatherData.SetMeasurements(30, 55, 20.4)
}

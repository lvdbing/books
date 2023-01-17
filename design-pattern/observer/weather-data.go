package main

// WeatherData 主题实现类
type WeatherData struct {
	ConcreteSubject

	temperature float64
	humidity    float64
	pressure    float64
}

func NewWeatherData() *WeatherData {
	wd := &WeatherData{}
	return wd
}

// NotifyObservers 通知观察者数据改变
// 通知观察者数据改变时, 可以主动推送数据, 也可以让观察者自己拉取数据
func (w *WeatherData) NotifyObservers() {
	if !w.HasChanged() {
		return
	}
	for _, ob := range w.Observers {
		ob.Update(w.temperature, w.humidity, w.pressure) // 主动推送改变的数据
	}
	w.ClearChanged()
}

func (w *WeatherData) MeasurementsChanged() {
	w.SetChanged()
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure

	w.MeasurementsChanged()
}

// GetTemperature 获取温度, 用以实现观察者主动拉取数据
func (w *WeatherData) GetTemperature() float64 {
	return w.temperature
}

// GetHumidity 获取湿度, 用以实现观察者主动拉取数据
func (w *WeatherData) GetHumidity() float64 {
	return w.humidity
}

// GetPressure 获取气压, 用以实现观察者主动拉取数据
func (w *WeatherData) GetPressure() float64 {
	return w.pressure
}

package main

// Observer 观察者接口
type Observer interface {
	Update(temp, humidity, pressure float64)
}

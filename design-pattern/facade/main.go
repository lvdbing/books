package main

func main() {
	facade := NewHomeTheaterFacade(Popper{}, Light{}, DVD{})
	facade.WatchMovie("Top Gun")
	facade.EndMovie()
}

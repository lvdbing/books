package main

import "fmt"

type Popper struct{}

func (p *Popper) On() {
	fmt.Println("Popcorn Popper on")
}

func (p *Popper) Pop() {
	fmt.Println("Popcorn Popper popping popcorn!")
}

func (p *Popper) Off() {
	fmt.Println("Popcorn Popper off")
}

type Light struct{}

func (l *Light) Dim(v int) {
	fmt.Printf("Lights dimming to %d%%\n", v)
}

func (l *Light) On() {
	fmt.Println("Lights on")
}

type DVD struct {
	movie string
}

func (d *DVD) On() {
	fmt.Println("DVD Player on")
}

func (d *DVD) Play(movie string) {
	d.movie = movie
	fmt.Printf("DVD Player playing \"%s\"\n", movie)
}

func (d *DVD) Stop() {
	fmt.Printf("DVD Player stopped \"%s\"\n", d.movie)
}

func (d *DVD) Eject() {
	fmt.Println("DVD Player eject")
}

func (d *DVD) Off() {
	fmt.Println("DVD Player off")
}

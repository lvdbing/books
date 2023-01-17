package main

import (
	"fmt"
	"strconv"
)

type Light struct {
	Name string
}

func NewLight(name string) *Light {
	return &Light{Name: name}
}

func (l *Light) On() {
	fmt.Printf("%s is on\n", l.Name)
}

func (l *Light) Off() {
	fmt.Println(l.Name + " is off")
}

type Stereo struct{}

func NewStereo() *Stereo {
	return &Stereo{}
}

func (s *Stereo) On() {
	fmt.Println("Stereo is on")
}

func (s *Stereo) Off() {
	fmt.Println("Stereo is off")
}

func (s *Stereo) SetCD() {
	fmt.Println("Stereo is setting CD")
}

func (s *Stereo) SetVolume(v int) {
	fmt.Println("Stereo set volume " + strconv.Itoa(v))
}

const (
	SPEED_OFF    = 0
	SPEED_LOW    = 1
	SPEED_MEDIUM = 2
	SPEED_HIGH   = 3
)

type Fan struct {
	name  string
	speed int
}

func NewFan(name string) *Fan {
	return &Fan{name: name, speed: SPEED_OFF}
}

func (f *Fan) High() {
	f.speed = SPEED_HIGH
	fmt.Printf("%s is on high\n", f.name)
}

func (f *Fan) Medium() {
	f.speed = SPEED_MEDIUM
	fmt.Printf("%s is on medium\n", f.name)
}

func (f *Fan) Low() {
	f.speed = SPEED_LOW
	fmt.Printf("%s is on low\n", f.name)
}

func (f *Fan) Off() {
	f.speed = SPEED_OFF
	fmt.Printf("%s is off\n", f.name)
}

func (f *Fan) NextSpeed() {
	switch f.speed {
	case SPEED_HIGH:
		f.Off()
	case SPEED_MEDIUM:
		f.High()
	case SPEED_LOW:
		f.Medium()
	case SPEED_OFF:
		f.Low()
	}
}

func (f *Fan) GetSpeed() int {
	return f.speed
}

package main

type Mocha struct {
	beverage Beverage
}

func NewMocha(b Beverage) *Mocha {
	return &Mocha{beverage: b}
}

func (m *Mocha) GetDescription() string {
	return m.beverage.GetDescription() + ", Mocha"
}

func (m *Mocha) Cost() int {
	return 20 + m.beverage.Cost()
}

type Soy struct {
	beverage Beverage
}

func NewSoy(b Beverage) *Soy {
	return &Soy{beverage: b}
}

func (s *Soy) GetDescription() string {
	return s.beverage.GetDescription() + ", Soy"
}

func (s *Soy) Cost() int {
	return 15 + s.beverage.Cost()
}

type Whip struct {
	beverage Beverage
}

func NewWhip(b Beverage) *Whip {
	return &Whip{beverage: b}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", Whip"
}

func (w *Whip) Cost() int {
	return 10 + w.beverage.Cost()
}

type Milk struct {
	beverage Beverage
}

func NewMilk(b Beverage) *Milk {
	return &Milk{beverage: b}
}

func (m *Milk) GetDescription() string {
	return m.beverage.GetDescription() + ", Milk"
}

func (m *Milk) Cost() int {
	return 12 + m.beverage.Cost()
}

package main

import (
	"fmt"
)

func main() {
	d := NewMallardDuck()
	d.Display()
	d.PerformFly()
	d.PerformQuack()
	d.Swim()

	fmt.Println()

	md := NewModelDuck()
	md.Display()
	md.PerformFly()
	md.PerformQuack()
	md.SetFlyBehavior(&FlyRocketPowered{})
	md.PerformFly()

	fmt.Println()

	rd := &RedheadDuck{}
	rd = rd.New()
	rd.Display()
	rd.PerformFly()
	rd.PerformQuack()

	fmt.Println()

	rubberDuck := RubberDuck{}.New()
	rubberDuck.Display()
	rubberDuck.PerformFly()
	rubberDuck.PerformQuack()

	fmt.Println()

	decoyDuck := &DecoyDuck{}
	decoyDuck = decoyDuck.New()
	decoyDuck.Display()
	decoyDuck.PerformFly()
	decoyDuck.PerformQuack()

	fmt.Println()

	king := NewKing()
	king.UseWeapon()

	queen := NewQueen()
	queen.UseWeapon()
	queen.SetWeaponBehavior(&BowAndArrowBehavior{})
	queen.UseWeapon()

	knight := NewKnight()
	knight.UseWeapon()

	troll := NewTroll()
	troll.UseWeapon()
}

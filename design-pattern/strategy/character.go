package main

import "fmt"

// ******* 用游戏角色实现策略模式 *******

type WeaponBehavior interface {
	UseWeapon()
}

type SwordBehavior struct{}

func (s *SwordBehavior) UseWeapon() {
	fmt.Println("Attack with a sword!")
}

type KnifeBehavior struct{}

func (k KnifeBehavior) UseWeapon() {
	fmt.Println("Attack with a knife!")
}

type BowAndArrowBehavior struct{}

func (b *BowAndArrowBehavior) UseWeapon() {
	fmt.Println("Attack with bow and arrow!")
}

type AxeBehavior struct{}

func (a *AxeBehavior) UseWeapon() {
	fmt.Println("Attack with an axe!")
}

type Character struct {
	Weapon WeaponBehavior
}

func (c *Character) UseWeapon() {
	c.Weapon.UseWeapon()
}

func (c *Character) SetWeaponBehavior(wb WeaponBehavior) {
	c.Weapon = wb
}

type King struct {
	Character
}

func NewKing() *King {
	k := &King{}
	fmt.Println("I'm a king")
	k.Weapon = &SwordBehavior{}
	return k
}

type Queen struct {
	Character
}

func NewQueen() *Queen {
	q := &Queen{}
	fmt.Println("I'm a queen")
	q.Weapon = KnifeBehavior{}
	return q
}

type Knight struct {
	Character
}

func NewKnight() *Knight {
	k := &Knight{}
	fmt.Println("I'm a knight")
	k.Weapon = &BowAndArrowBehavior{}
	return k
}

type Troll struct {
	Character
}

func NewTroll() *Troll {
	t := &Troll{}
	fmt.Println("I'm a troll")
	t.Weapon = &AxeBehavior{}
	return t
}

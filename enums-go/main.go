package main

import "fmt"

// weapon type
// Axe
// Sword
// Wooden stick
// Knife

// === CONSTRUCTING ENUM ======

type weaponType int

// const (

// 	Axe         weaponType = 1
// 	Sword       weaponType = 2
// 	WoodenStick weaponType = 3
// 	Knife       weaponType = 4

// )

// ====						==========
	// ===ADDDING A FUNCTION 
// ====						======

func (w weaponType) string() string  {
	switch w {
	case Sword:
		return "SWORD"
		case Axe:
		return "AXE"
	}
	return ""
	
}

// === TRICK THIA WILL INCREMENT ALL BELOW ====
const (
	Axe weaponType = iota
	Sword
	WoodenStick
	Knife
)

func getDamage(weaponType weaponType) int {
	switch weaponType {
	case Axe:
		return 200
	case Sword:
		return 100
	case WoodenStick:
		return 2

	case Knife:
		return 50
	default:
		panic("Weapon does not exist")
	}
}

func main() {
	// ====Trick====
	fmt.Printf("Damage of the weapon(%d) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("Damage of the weapon(%d) (%d):\n", Sword, getDamage(Sword))
	fmt.Printf("Damage of the weapon(%d) (%d):\n", WoodenStick, getDamage(WoodenStick))
	fmt.Printf("Damage of the weapon(%d) (%d):\n", Knife, getDamage(Knife))

	// ======== To print the output of the function above with a string interface ========
	fmt.Printf("Damage of weapon(%s) (%d):\n", Axe, getDamage(Axe))
	fmt.Printf("Damage of weapon(%s) (%d):\n", Sword, getDamage(Sword))

	
	 
	// fmt.Println("Damage of the given weapon:",  getDamage(Axe))
}

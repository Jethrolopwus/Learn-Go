
package main

import "fmt"

// ====== custom types ======
var (
    floatVar32 float32 = 0.2
    floatVar64 float64 = 0.5
    name       string  = "John "
    intVar32   int32   = 1
    intVar64   int64   = 25556
    intVar     int     = -10
    uinVar     uint32  = 1
    uintVar    uint64  = 233
    uint8Var   uint8   = 0x2
    byteVar    byte    = 0x1
    runVar     rune    = 'a'
)

// ========= Structs ============
// definition
type player struct {
    name        string
    health      int
    attackPower float64
}

// ==== struct Receiver function ====
func (p player) getHealth() int {
    return p.health
}

// ======CUSTOME TYPES ======

type Weapon string

func getWeapon(weapon Weapon) string  {
	return string(weapon)
}


func main() {

    // Create the player first
    // player := player{
    //     name:        "Captain John",
    //     health:      101,
    //     attackPower: 50.1,
    // }
// =========== Maps =================
// initializing an empty Map 
	users := map[string]int{}

	users = make(map[string]int)

	users["students"] = 10
	users["Staffs"] = 20
	
	// ==== How to access the value of a map====
	age := users["students"]

	// ==== checking if a key exist or not in a map==========
	age, ok := users["Management"]
	if !ok {
		fmt.Println("Management Does not exist")
	} else {
		fmt.Println("Students existed in the map ", age)
	}
	// ===How to delete in a map====
	delete(users, "Staffs")

	// ===== How to range over a map=====
	for key, v := range users {
		fmt.Printf("The key is %s and the value is %d\n: ", key, v)
	}

    // fmt.Printf("This is the player's Health: %d\n ", player.getHealth())
    // fmt.Printf("This is the player: %+v\n ", player)
    fmt.Printf("The users are: %+v\n ", users)
    fmt.Printf("The users's age is: %d\n ", age)
	fmt.Println(users)



	// ===== Slices ======
	numbers := []int{1, 2,3,4}
	otherNumbers := make([]int, 2 )

	fmt.Println(numbers)
	fmt.Println(otherNumbers)



	// ====ARRAYS ======

	fruits := [2]int{}

	fmt.Printf(" the fruits are %+v\n:",fruits)

	// ====APPENDING A NEW NUMBER INTO A SLICE ====
	numbers = append(numbers, 5)
	numbers = append(numbers, 10)
	fmt.Printf("The appended number is %v\n : ", numbers)



}
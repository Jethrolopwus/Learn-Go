package main

import "fmt"

//==== Any animal that can make a sound is a "SoundMaker"===== Scenario 1======

type SoundMaker interface {
	MakeSound() string
}

type Dog struct {
	Name string
}

// Dog implements SoundMaker by defining MakeSound()
func (d Dog) MakeSound() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

// Cat also implements SoundMaker
func (c Cat) MakeSound() string {
	return "Meow!"
}

// This function works with ANY SoundMaker
func MakePetSound(pet SoundMaker) {
	fmt.Println(pet.MakeSound())
}

// =========scenario 2  Data Storage interfaces =============

type NumberStorer interface {
	getAll() ([] int, error)
	put(int) error
}

// to implement PostgrSQL instead of mongoDB=====

type postgresNumberStore struct{
	// postgress values (db connection)============
}
//     ====== receiver functions for postgress implementation =======

func  (s postgresNumberStore) getAll() ([] int, error) {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, nil
}
func  (s postgresNumberStore) put(number int) error {
		fmt.Println("Store the number into the postgres storage")
		return nil
}



type mongoDBNumberStore struct{
	// some values
}

//======= implementation of the struct above using a reciever function========
func  (m mongoDBNumberStore) getAll() ([] int, error) {
	return []int{1, 2, 3, 4}, nil
}
func  (m mongoDBNumberStore) put(number int) error {
		fmt.Println("Store the number into the mogoDB storage")
		return nil
}

type ApiServer struct{
numberStore NumberStorer
}

func main() {
	// implementing the storage interfaces for mongoDB
	apiServer := ApiServer{
		numberStore: mongoDBNumberStore{},
	}


	// implementing the storage interfaces for postgres
	apiServer = ApiServer{
		numberStore: postgresNumberStore{},
	}

	// postgres logic handles here
	if err := apiServer.numberStore.put(1); err != nil {
		panic(err)
	}
	// mongoDB logic handles here
	numbers, err := apiServer.numberStore.getAll()
	if err != nil {
		panic(err)
	}
	// if No Errors then
	fmt.Println(numbers)



	// ========= DOg implementation============

	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Both can be used as SoundMakers
	MakePetSound(dog)
	MakePetSound(cat)
}

package main

import (
	"fmt"
)
// this is a name struct
type Person struct {
	name    string
	age     int
	gender  string
	alive   bool
	// this is a type struct====================
	health  physicalData
	// This is anonomyous struct===================
	address struct {
		city    string
		country string
	}
}

type physicalData struct {
	height float32
	weight float32
}


type Voter struct {
	Person
	voterID int
}

func userData(p Person) {
	// Person data print=================================
	fmt.Println("User name : ", p.name)
	fmt.Println("User age : ", p.age)
	fmt.Println("User alive : ", p.alive)
	fmt.Println("User gender : ", p.gender)
	fmt.Println("User Height : ", p.health.height)
	fmt.Println("User weight : ", p.health.weight)
	fmt.Println("User weight : ", p.health.weight)
	fmt.Println("User City : ", p.address.city)
	fmt.Println("User Country : ", p.address.country)

	
}

func voterData (v Voter){
	// Voter data print 
	fmt.Println("voter id",v.voterID)

	fmt.Println("User name : ", v.name)
	fmt.Println("User age : ", v.age)
	fmt.Println("User alive : ", v.alive)
	fmt.Println("User gender : ", v.gender)
	fmt.Println("User Height : ", v.health.height)
	fmt.Println("User weight : ", v.health.weight)
	fmt.Println("User weight : ", v.health.weight)
	fmt.Println("User City : ", v.address.city)
	fmt.Println("User Country : ", v.address.country)
	// stuct method test =======================================
	fmt.Println(v.add())
}



// struct methods====================================================
func (p Person) add() string{
	return p.name + p.gender
}

func main() {

	userData(Person{
		name:   "munna",
		age:    33,
		gender: "male",
		alive:  true,
		health: physicalData{
			height: 70.5,
			weight: 5.5,
		},
		address: struct {
			city    string
			country string
		}{
			city:    "comilla",
			country: "bangladesh",
		},
	})


	voterData(Voter{
		Person: Person{
			name:   "11munna",
		age:    133,
		gender: "111male",
		alive:  true,
		health: physicalData{
			height: 70.5,
			weight: 5.5,
		},
		address: struct {
			city    string
			country string
		}{
			city:    "comilla",
			country: "bangladesh",
		},
		},
		voterID: 234223234,
	})


}

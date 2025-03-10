package main

import (
	"fmt"
	"learngo/basics/models"
	// "time"
)

func main() {
	structsAndMapsInGo()
}

func structsAndMapsInGo() {
	/// Maps
	fmt.Println("\n→→→→→→→→→→→ START:BLOCK MAPS IN GO LANG ←←←←←←←←←")
	mapsInGo()

	mapsAreNotReferenceVariables()

	fmt.Println("→→→→→→→→→→→ END:BLOCK MAPS IN GO LANG ←←←←←←←←←")

	/// Custom Data Type struct in Go Lang
	structs()

	/*
		Receiver Functions
	*/
	receiverFunctions()
}

/*
MAPS in GO Lang
*/
func mapsInGo() {
	// only declaration
	var mapIngo map[int]string
	fmt.Println(mapIngo)
	// mapIngo[89] = "Eighty Nine" // 💀 assignment to entry in nil map ❌
	fmt.Println("is mapIngo value == nil ? : ", mapIngo == nil) // true
	fmt.Println(len(mapIngo))
	fmt.Println()

	mapUsingMake := make(map[string]bool)
	fmt.Printf("Map of type %T created using make fn and value is : %v \n", mapUsingMake, mapUsingMake)
	fmt.Println("is mapUsingMake value == nil ?  ", mapUsingMake == nil) // false, so we can put values into map

	subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
	fmt.Println(subjectMarks)
	fmt.Println("Length of subjectMarks Map : ", len(subjectMarks))

	marks := map[string]int{}
	fmt.Println("is marks value == nil ? : ", marks == nil) // false, marks is not nil

	marks["person1"] = 45
	marks["person4"] = 45
	marks["person5"] = 45
	marks["person6"] = 45
	marks["person7"] = 45

	/// Accessing a value from map using key
	personMarks, doesPersonHaveMarks := marks["person6"]
	if doesPersonHaveMarks {
		fmt.Println("person6 Marks : ", personMarks)
	}
	// another way of accessing the value and checking if it exists
	if anotherPersonMarks, doesExists := marks["person7"]; doesExists {
		fmt.Println("Another person marks accessed from a map ", anotherPersonMarks)
	}

	delete(marks, "person7")
	delete(marks, "keyNotExist") // does nothing
	fmt.Printf("Type of marks is %T and len is %v and data \n\t\t %v\n", marks, len(marks), marks)

	fmt.Println()
	// Create a map with string keys and dynamic values
	dynamicMap := make(map[string]interface{})

	// Add key-value pairs to the map
	dynamicMap["name"] = "John"
	dynamicMap["age"] = 30
	dynamicMap["isStudent"] = true
	dynamicMap["map"] = map[string]float64{"sub1": 90.0}
	dynamicMap["slice"] = []float64{90.0}
	dynamicMap["array"] = [...]int{90, 3, 4, 5, 6, 7}

	// Print the map
	fmt.Printf("dynamicMap type : %T with string keys and dynamic values : %v\n", dynamicMap, dynamicMap)
	per, doesExists := dynamicMap["age"].(int)
	if doesExists {
		println(per)
	}

}

/*

	STRUCTS in GO Lang

*/

type Person struct {
	age  int
	name string
}

func structs() {
	fmt.Println("\n→→→→→→→→→→→ START:BLOCK  STRUCTS IN GO LANG ←←←←←←←←←")

	/// Different ways to create objects of type [Person]
	defaultPerson := Person{}
	fmt.Println("defaultPerson: ", defaultPerson)

	var customPerson Person
	customPerson.age = 90
	fmt.Println("customPerson: ", customPerson)

	shortPerson := Person{age: 23, name: "7xp"}
	shortPerson2 := Person{78, "without named params"}
	fmt.Printf("shortPerson: %v\n", shortPerson)
	fmt.Printf("shortPerson2: %+v\n", shortPerson2)

	/*
		Creating the instance of a Person using `new` keyword

		🌻 Use Case: useful if you plan to modify the struct's fields directly
		or if you want to pass the struct around without copying it.
		Pointers are efficient when you have large structs or need to
		share data across multiple functions or goroutines.

		Read more about new() & make() in Pointers section
	*/
	p := new(Person)
	// p := &Person{age: 56} same as above, p hold a pointer to Person
	p.age = 56
	fmt.Println(p.age, p.name)

	structSlice := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("Slice of type Struct ", structSlice)

	/// Robot struct is created in models folder, let's create new Robot
	firstRobot := models.CreateRobot("Robo 1.0", 3.0)
	fmt.Printf("First Robot of type %T created using struct :%v\n", firstRobot, firstRobot)
	firstRobot.Name = "updated Name"
	fmt.Printf("First Robot of type %T created using struct :%v\n", firstRobot, firstRobot)

	/// receiver function on Robot struct
	firstRobot.PrintFormattedRobotDetails()

	// Create another robot
	robot2 := models.Robot{Name: "GPT Robot", Power: 5.0}
	fmt.Println("robot2 Details ", robot2)
	pointerRobot := &robot2
	pointerRobot.Name = "Pointer Robot" // this modifies the robot2.Name as well
	fmt.Println("robot2 Details ", robot2)

	/// structs and maps combination
	fmt.Printf("\n ------> Struct and Maps Combination <------ \n")
	peopleMap := map[string]Person{"india": {name: "Indian"}}
	peopleMap["Nepal"] = Person{name: "Nepali"}
	fmt.Printf("\npeopleMap value is : %v & is peopleMap nil ? %v\n\n", peopleMap, peopleMap == nil)

	var nilPeopleMap map[string]Person
	fmt.Printf("nilPeopleMap value is : %v & is nilPeopleMap nil ? %v\n", nilPeopleMap, nilPeopleMap == nil)
	// nilPeopleMap["China"] = Person{name: "Chinese"} // 🚨 panic: assignment to entry in nil map
	nilPeopleMap = make(map[string]Person)
	// now nilPeopleMap is usable
	nilPeopleMap["China"] = Person{name: "Chinese", age: 89}
	fmt.Printf("updated nilPeopleMap value is : %v & is nilPeopleMap nil ? %v\n", nilPeopleMap, nilPeopleMap == nil)

	if chinaPersonDetails, doesExist := nilPeopleMap["China"]; doesExist {
		fmt.Println("chinaPersonDetails: ", chinaPersonDetails)
	}

	fmt.Println("→→→→→→→→→→→ END:BLOCK  STRUCTS IN GO LANG ←←←←←←←←←")
}

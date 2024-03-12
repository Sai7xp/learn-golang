/*
* Created on 29 Feb 2024
* @author Sai Sumanth
 */

package main

import "fmt"

// lets suppose Student can have marks of type int or float64
// in this case we can make use of GENERICS very well
type Student[T []int | []float64] struct {
	Id    int
	Name  string
	Marks T
}

type CustomData interface {
	int | float64 | int8 | bool | string
}

type MadMan[T CustomData] struct {
	Name          string
	SaidSomething T
}

func genericsWithStructs() {
	fmt.Println("\n----> BLOCK:START GENERICS with struct <------")
	user1 := Student[[]float64]{Id: 4, Name: "user1", Marks: []float64{74.5, 71.8}}
	fmt.Printf("User Details : %+v\n", user1)

	madman1 := MadMan[string]{Name: "madMan1", SaidSomething: "Heyo!"}
	yesOrNoMadMan := MadMan[bool]{Name: "realMadMan", SaidSomething: true}

	fmt.Printf("madman1 Details : %+v\n", madman1)
	fmt.Printf("yesOrNoMadMan Details : %+v\n", yesOrNoMadMan)

	fmt.Printf("----> BLOCK:END GENERICS with struct <------\n")

	/// 💡 type alias with GENERICS
	type MarksType[T int | float64] []T
	marks := MarksType[int]{1, 2, 3, 4, 5}
	marks2 := MarksType[float64]{1.4, 2.8}

	fmt.Println(marks)
	fmt.Println(marks2)

	/// we can use multiple generics as well
	type FrequencyMap map[string]int
	mp1 := FrequencyMap{"a": 8, "b": 89}
	fmt.Printf("mp1 : %+v\n", mp1)

	type CustomMap[T comparable, V int | bool] map[T]V
	freqMap := CustomMap[string, int]{"hello": 3, "world": 5}
	doesExistMap := CustomMap[int, bool]{3: true, 5: false}
	fmt.Printf("CustomMap freq : %+v\n", freqMap)
	fmt.Printf("CustomMap doesExistMap : %+v\n", doesExistMap)
}

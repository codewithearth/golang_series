package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slice in golang")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	list1 := append(fruitList[1:])
	fmt.Println("list1 : ", list1)

	list2 := append(fruitList[1:3])
	fmt.Println("list2 : ", list2)

	list3 := append(fruitList[:3])
	fmt.Println("list3 : ", list3)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 456
	highScores[3] = 867
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from slices based on index

	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)

}

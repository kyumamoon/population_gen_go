package main

import (
	"fmt"
	"math/rand"
	"time"
)

var populationA int = 1000
var tickerstop int = 0

type tests struct {
	Name string
	Age  int
}

func illnesses() int {
	x := 45
	return x
}

func rolldeath() {
	x := rand.Intn(101)

	if x >= 5 && x <= 25 { // 25% chance
		populationA -= 1
	} else if x >= 25 && x <= 50 { // 50% chance
		populationA += 1
	} else if x >= 0 && x <= 5 { // 5% chance
		populationA -= 2
	}
}

func tickerfunc() {
	ticker := time.NewTicker(1 * time.Second)

	for _ = range ticker.C {
		//fmt.Println("tock")
		rolldeath()
		fmt.Println(populationA)
		if tickerstop == 1 {
			ticker.Stop()
		}
	}
	//fmt.Println(populationA)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	dan := tests{"dan", 35}
	fmt.Println(dan.Name, dan.Age)
	fmt.Println(illnesses())
	//maxrolls := 0

	//for maxrolls <= 1000 {
	//rolldeath()
	//maxrolls++
	//}
	go tickerfunc()

	timer := time.NewTimer(10 * time.Second)

	<-timer.C
	fmt.Println("Timer Ended")
	tickerstop++

	select {}
}

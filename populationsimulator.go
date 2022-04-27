package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

var populationtotal int = 1000

type person struct {
	ID          int
	Age         int
	AliveStatus string
	WeeksAlive  int
	Birthdays   int
	Gender      string
	Pregnant    int
	DeathDate   int
	BornDate    int
	Offsprings  int
}

var peoplelist []person
var ogpeoplelist []person
var populationgrowth []int
var weeks int // 1 is equal to 2 weeks
var yearlimit int = 26
var quarterlimit float64 = 6.5
var newyear int = 0
var newquarter float64
var timerring int = 0

func populationage() {

	for i := 1; i < (populationtotal + 1); i++ {
		randage := rand.Intn(70) + 1
		randgender := rand.Intn(2) + 1

		var gender string
		switch randgender {
		case 1:
			gender = "Male"
		case 2:
			gender = "Female"
		}

		temp := person{i, randage, "Alive", 0, 0, gender, 0, 0, weeks, 0}
		peoplelist = append(peoplelist, temp)
		ogpeoplelist = append(ogpeoplelist, temp)
	}

}

func calculatechance(age int) int {
	A := age
	var B int
	if A > 0 && A <= 2 {
		B = 2
	} else if A > 2 && A <= 5 {
		B = 4
	} else if A > 5 && A <= 10 {
		B = 4
	} else if A > 10 && A <= 15 {
		B = 5
	} else if A > 15 && A <= 20 {
		B = 8
	} else if A > 20 && A <= 25 {
		B = 8
	} else if A > 25 && A <= 35 {
		B = 10
	} else if A > 35 && A <= 45 {
		B = 10
	} else if A > 45 && A <= 55 {
		B = 12
	} else if A > 55 && A <= 65 {
		B = 15
	} else if A > 65 && A <= 80 {
		B = 18
	} else if A > 80 {
		B = 20
	}

	return B
}

func birthday() {
	if newyear > yearlimit {
		newyear = 0
		for pos, ages := range peoplelist {
			if ages.AliveStatus == "Alive" {
				peoplelist[pos].Age++
				peoplelist[pos].Birthdays++

				if ages.Pregnant == 1 {
					peoplelist[pos].Pregnant = 0
				}

			} else {
				// skip dead person
			}
		}

	} else {

	}
}

func birthsimulator() {
	var compatiblefemales int
	var compatiblemales int

	for _, ages := range peoplelist {
		if ages.AliveStatus == "Alive" {
			if ages.Gender == "Female" { // if female
				if ages.Age > 18 && ages.Age < 50 { // if old enough
					if ages.Pregnant == 0 { // not pregnant
						compatiblefemales++
					} else {
						// pregnant do nothing
					}
				} else {
					// too young
				}
			} else {
				// if male do nothing
			}
		} else {
			// skip dead females
		}
	}

	for _, ages := range peoplelist {
		if ages.AliveStatus == "Alive" {
			if ages.Gender == "Male" { // if female
				if ages.Age > 18 && ages.Age < 50 { // if old enough
					if ages.Pregnant == 0 { // not pregnant
						compatiblemales++
					} else {
						// pregnant do nothing
					}
				} else {
					// too young
				}
			} else {
				// if male do nothing
			}
		} else {
			// if dead do nothing
		}
	}

	//if compatiblefemales > compatiblemales {
	if newquarter > quarterlimit {
		newquarter = 0
		for pos, ages := range peoplelist {
			if ages.AliveStatus == "Alive" {
				if ages.Gender == "Female" { // if female
					if ages.Age > 18 && ages.Age < 50 { // if old enough
						if ages.Pregnant == 0 { // not pregnant
							if ages.Offsprings < 5 {
								babychance := rand.Intn(100) + 1
								if babychance <= 10 { // 10% chance of making babies
									newbabyID := len(peoplelist) + 1
									randgender := rand.Intn(2) + 1

									var gender string
									switch randgender {
									case 1:
										gender = "Male"
									case 2:
										gender = "Female"
									}

									temp := person{newbabyID, 1, "Alive", 0, 0, gender, 0, 0, weeks, 0}
									peoplelist = append(peoplelist, temp)
									peoplelist[pos].Pregnant = 1
									peoplelist[pos].Offsprings++
								} else {
									//no make babies
								}
							} else {
								// too much kids, lower chance
								babychance := rand.Intn(100) + 1
								if babychance <= 3 { // 10% chance of making babies
									newbabyID := len(peoplelist) + 1
									randgender := rand.Intn(2) + 1

									var gender string
									switch randgender {
									case 1:
										gender = "Male"
									case 2:
										gender = "Female"
									}

									temp := person{newbabyID, 1, "Alive", 0, 0, gender, 0, 0, weeks, 0}
									peoplelist = append(peoplelist, temp)
									peoplelist[pos].Pregnant = 1
									peoplelist[pos].Offsprings++
								} else {
									// make no babies
								}
							}

						} else {
							// pregnant do nothing
						}
					} else {
						// too young
					}
				} else {
					// if male do nothing
				}
			} else {
				// if dead do nothing
			}
		}

	} else {
		// not quarter yet
	}
	//}
	fmt.Printf("Compatible Females: %d \n", compatiblefemales)
	fmt.Printf("Compatible Males: %d \n", compatiblemales)

	var alive int = 0
	for _, ages := range peoplelist {
		if ages.AliveStatus == "Alive" {
			alive++
		}
	}
	// every week (2 weeks) record population
	populationgrowth = append(populationgrowth, alive)
}

func deathsimulation() {
	var totalalive int
	weeks++
	newyear++
	newquarter++

	for _, ages := range peoplelist {
		if ages.AliveStatus == "Alive" {
			totalalive++
		} else {
			//nothing
		}
	}

	for pos, ages := range peoplelist {
		if ages.AliveStatus == "Alive" {
			A := ages.Age
			B := calculatechance(A)
			rolldeath := rand.Intn(2000) + 1
			if rolldeath <= B {
				deathchance := rand.Intn(2) + 1
				if deathchance == 2 {
					peoplelist[pos].AliveStatus = "Deceased"
					peoplelist[pos].DeathDate = weeks
				}
				//peoplelist[pos].AliveStatus = "Deceased"
			} else {
				peoplelist[pos].WeeksAlive++
				peoplelist[pos].WeeksAlive++
			}
		} else {
			// skip dead person
		}
	}

	birthday()

	//for _, ages := range peoplelist { DEBUG REMOVED
	//fmt.Println(ages)
	//}

	fmt.Printf("Starting Population: 300 \n")
	fmt.Printf("Ending Population: %d \n", len(peoplelist))
	fmt.Printf("Total alive population count: %d \n", totalalive)
	Years := float64(weeks) / 26
	fmt.Printf("Week: %d | Year: %f \n", weeks*2, Years)
	populationtotal = totalalive

	if populationtotal > 0 && timerring == 0 {
		birthsimulator()
		deathsimulation()
	} else if populationtotal == 0 || timerring == 1 {
		fmt.Println("First List")
		for _, ages := range ogpeoplelist {
			fmt.Println(ages)
		}

		// create excel

		f := excelize.NewFile()
		index := f.NewSheet("Sheet1")
		index2 := f.NewSheet("Sheet2")

		for pos, ages := range peoplelist {
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(pos), pos)
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(pos), ages.Age)
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(pos), ages.AliveStatus)
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(pos), ages.WeeksAlive)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(pos), ages.Birthdays)
			f.SetCellValue("Sheet1", "F"+strconv.Itoa(pos), ages.Gender)
			f.SetCellValue("Sheet1", "G"+strconv.Itoa(pos), ages.Pregnant)
			f.SetCellValue("Sheet1", "H"+strconv.Itoa(pos), ages.DeathDate)
			f.SetCellValue("Sheet1", "I"+strconv.Itoa(pos), ages.BornDate)
			f.SetCellValue("Sheet1", "J"+strconv.Itoa(pos), ages.Offsprings)
		}

		f.SetActiveSheet(index)

		var weekindex int = 2
		for pos1, num := range populationgrowth {
			f.SetCellValue("Sheet2", "A"+strconv.Itoa(pos1), weekindex)
			f.SetCellValue("Sheet2", "B"+strconv.Itoa(pos1), num)
			weekindex = weekindex + 2
		}

		f.SetActiveSheet(index2)

		if err := f.SaveAs("VirtualWorld7.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
}

func timer1() {
	timer := time.NewTimer(900 * time.Second)
	<-timer.C
	timerring = 1
}

func main() {
	rand.Seed(time.Now().UnixNano())
	populationage()
	fmt.Println(len(peoplelist))
	for _, ages := range peoplelist {
		fmt.Println(ages)
	}

	go timer1()
	deathsimulation()
	fmt.Println(populationtotal)

	select {}
}

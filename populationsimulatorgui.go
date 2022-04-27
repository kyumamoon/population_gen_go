package main

import (
	"fmt"
	"math/rand"
	"time"

	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/360EntSecGroup-Skylar/excelize"
)

var populationname string
var populationcount int // starting pop
var birthrate int       // birth rate out of 100
var maxbirthcount int   // max children
var rarebirthprob int   // rare birth rate out of 100
var finished int = 0
var newbornmortalityrate int
var toddlermortalityrate int
var childmortalityrate int
var youngteenmortalityrate int
var lateteenmortalityrate int
var midtwentiesmortalityrate int
var midthirtiesmortalityrate int
var midfortiesmortalityrate int
var midfiftiesmortalityrate int
var midsixtiesmortalityrate int
var oldergenmortalityrate int
var rareoldergenmortalityrate int
var exporttoexcelenabled string
var excelfileexportname string
var exportdatachoice string
var simulationtime int
var submitbuttontext string = "Submit"
var buttonpressed int = 0
var minageforbirth int
var maxageforbirth int

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

var weeks int // 1 is equal to 2 weeks
var newyear int = 0
var newquarter float64
var populationgrowth []int
var yearlimit int = 26
var quarterlimit float64 = 6.5
var timerring int = 0
var timerring2 int = 0

func stringtoint(strings string) int {
	stringtoint, _ := strconv.Atoi(strings)
	return stringtoint
}

func populationage() {

	for i := 1; i < (populationcount + 1); i++ {
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

func birthsimulator() {
	var compatiblefemales int
	var compatiblemales int

	for _, ages := range peoplelist {
		if ages.AliveStatus == "Alive" {
			if ages.Gender == "Female" { // if female
				if ages.Age > minageforbirth && ages.Age < maxageforbirth { // if old enough
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
					if ages.Age > minageforbirth && ages.Age < maxageforbirth { // if old enough
						if ages.Pregnant == 0 { // not pregnant
							if ages.Offsprings < maxbirthcount {
								babychance := rand.Intn(100) + 1
								if babychance <= birthrate { // 10% chance of making babies
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
								if babychance <= rarebirthprob { // 10% chance of making babies
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
	//fmt.Printf("Compatible Females: %d \n", compatiblefemales)
	//fmt.Printf("Compatible Males: %d \n", compatiblemales)

	var alive int = 0
	for _, ages := range peoplelist {
		if ages.AliveStatus == "Alive" {
			alive++
		}
	}
	// every week (2 weeks) record population
	populationgrowth = append(populationgrowth, alive)
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

func calculatechance(age int) int {
	A := age
	var B int
	if A > 0 && A <= 2 {
		B = newbornmortalityrate
	} else if A > 2 && A <= 5 {
		B = toddlermortalityrate
	} else if A > 5 && A <= 10 {
		B = childmortalityrate
	} else if A > 10 && A <= 15 {
		B = youngteenmortalityrate
	} else if A > 15 && A <= 20 {
		B = lateteenmortalityrate
	} else if A > 20 && A <= 25 {
		B = midtwentiesmortalityrate
	} else if A > 25 && A <= 35 {
		B = midthirtiesmortalityrate
	} else if A > 35 && A <= 45 {
		B = midfortiesmortalityrate
	} else if A > 45 && A <= 55 {
		B = midfiftiesmortalityrate
	} else if A > 55 && A <= 65 {
		B = midsixtiesmortalityrate
	} else if A > 65 && A <= 80 {
		B = oldergenmortalityrate
	} else if A > 80 {
		B = rareoldergenmortalityrate
	}

	return B
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

	//REMOVED DEBUG PRINTS
	//fmt.Printf("Starting Population: 300 \n")
	//fmt.Printf("Ending Population: %d \n", len(peoplelist))
	//fmt.Printf("Total alive population count: %d \n", totalalive)
	Years := float64(weeks) / 26
	//fmt.Printf("Week: %d | Year: %f \n", weeks*2, Years)
	populationcount = totalalive

	//                                                                   IF EVERYONE DIED OR SIMULATION TIME ENDED
	if populationcount > 0 && timerring == 0 {
		birthsimulator()
		deathsimulation()
	} else if populationcount == 0 || timerring == 1 {
		//fmt.Println("First List")
		//for _, ages := range ogpeoplelist {
		//	fmt.Println(ages)
		//}

		// create excel
		if exporttoexcelenabled == "yes" {
			if exportdatachoice == "Alive Data Only" {

				f := excelize.NewFile()
				index := f.NewSheet("Sheet1")  // INFO SHEET
				index2 := f.NewSheet("Sheet2") // PEOPLE DATA
				index3 := f.NewSheet("Sheet3") // POPULATION GROWTH DATA

				f.SetCellValue("Sheet1", "A1", "Population Name")
				f.SetCellValue("Sheet1", "B1", populationname)
				f.SetCellValue("Sheet1", "A2", "Population Starting Amount")
				f.SetCellValue("Sheet1", "B2", len(ogpeoplelist))
				f.SetCellValue("Sheet1", "A3", "Population Ending Amount")
				f.SetCellValue("Sheet1", "B3", len(peoplelist))
				f.SetCellValue("Sheet1", "A4", "Population Age (Weeks)")
				f.SetCellValue("Sheet1", "B4", weeks*2)
				f.SetCellValue("Sheet1", "A5", "Population Age (Years)")
				f.SetCellValue("Sheet1", "B5", Years)

				f.SetActiveSheet(index)

				f.SetCellValue("Sheet2", "A1", "Person ID")
				f.SetCellValue("Sheet2", "B1", "Person Age")
				f.SetCellValue("Sheet2", "C1", "Alive?")
				f.SetCellValue("Sheet2", "D1", "Weeks Survived")
				f.SetCellValue("Sheet2", "E1", "Birthdays Celebrated")
				f.SetCellValue("Sheet2", "F1", "Gender")
				f.SetCellValue("Sheet2", "G1", "Pregnant?")
				f.SetCellValue("Sheet2", "H1", "Births")
				f.SetCellValue("Sheet2", "I1", "Birthday (Week Born)")

				alivepos := 1

				for pos, ages := range peoplelist {
					if ages.AliveStatus == "Alive" {
						f.SetCellValue("Sheet2", "A"+strconv.Itoa(alivepos+1), pos)
						f.SetCellValue("Sheet2", "B"+strconv.Itoa(alivepos+1), ages.Age)
						f.SetCellValue("Sheet2", "C"+strconv.Itoa(alivepos+1), ages.AliveStatus)
						f.SetCellValue("Sheet2", "D"+strconv.Itoa(alivepos+1), ages.WeeksAlive)
						f.SetCellValue("Sheet2", "E"+strconv.Itoa(alivepos+1), ages.Birthdays)
						f.SetCellValue("Sheet2", "F"+strconv.Itoa(alivepos+1), ages.Gender)
						f.SetCellValue("Sheet2", "G"+strconv.Itoa(alivepos+1), ages.Pregnant)
						f.SetCellValue("Sheet2", "H"+strconv.Itoa(alivepos+1), ages.Offsprings)
						f.SetCellValue("Sheet2", "I"+strconv.Itoa(alivepos+1), ages.BornDate)
						alivepos++
					} else {
						// skip if not alive
					}
				}

				if alivepos == 1 {
					f.SetCellValue("Sheet2", "A"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "B"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "C"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "D"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "E"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "F"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "G"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "H"+strconv.Itoa(alivepos+1), "No survivors")
					f.SetCellValue("Sheet2", "I"+strconv.Itoa(alivepos+1), "No survivors")
				}

				f.SetActiveSheet(index2)

				var weekindex int = 2

				f.SetCellValue("Sheet3", "A1", "Week")
				f.SetCellValue("Sheet3", "B1", "Population Alive")

				for pos1, num := range populationgrowth {
					f.SetCellValue("Sheet3", "A"+strconv.Itoa(pos1+2), weekindex)
					f.SetCellValue("Sheet3", "B"+strconv.Itoa(pos1+2), num)
					weekindex = weekindex + 2
				}

				f.SetActiveSheet(index3)
				f.SetActiveSheet(index)

				if err := f.SaveAs(excelfileexportname); err != nil {
					fmt.Println(err)
				}

				fmt.Println("Sim Finished")
				finished = 1

			} else if exportdatachoice == "Both" {

				f := excelize.NewFile()
				index := f.NewSheet("Sheet1")  // INFO SHEET
				index2 := f.NewSheet("Sheet2") // PEOPLE DATA
				index3 := f.NewSheet("Sheet3") // POPULATION GROWTH DATA

				f.SetCellValue("Sheet1", "A1", "Population Name")
				f.SetCellValue("Sheet1", "B1", populationname)
				f.SetCellValue("Sheet1", "A2", "Population Starting Amount")
				f.SetCellValue("Sheet1", "B2", len(ogpeoplelist))
				f.SetCellValue("Sheet1", "A3", "Population Ending Amount")
				f.SetCellValue("Sheet1", "B3", len(peoplelist))
				f.SetCellValue("Sheet1", "A4", "Population Age (Weeks)")
				f.SetCellValue("Sheet1", "B4", weeks*2)
				f.SetCellValue("Sheet1", "A5", "Population Age (Years)")
				f.SetCellValue("Sheet1", "B5", Years)

				f.SetActiveSheet(index)

				f.SetCellValue("Sheet2", "A1", "Person ID")
				f.SetCellValue("Sheet2", "B1", "Person Age")
				f.SetCellValue("Sheet2", "C1", "Alive?")
				f.SetCellValue("Sheet2", "D1", "Weeks Survived")
				f.SetCellValue("Sheet2", "E1", "Birthdays Celebrated")
				f.SetCellValue("Sheet2", "F1", "Gender")
				f.SetCellValue("Sheet2", "G1", "Pregnant?")
				f.SetCellValue("Sheet2", "H1", "Births")
				f.SetCellValue("Sheet2", "I1", "Birthday (Week Born)")
				f.SetCellValue("Sheet2", "J1", "Deathday (Week Died)")

				for pos, ages := range peoplelist {
					f.SetCellValue("Sheet2", "A"+strconv.Itoa(pos+2), pos)
					f.SetCellValue("Sheet2", "B"+strconv.Itoa(pos+2), ages.Age)
					f.SetCellValue("Sheet2", "C"+strconv.Itoa(pos+2), ages.AliveStatus)
					f.SetCellValue("Sheet2", "D"+strconv.Itoa(pos+2), ages.WeeksAlive)
					f.SetCellValue("Sheet2", "E"+strconv.Itoa(pos+2), ages.Birthdays)
					f.SetCellValue("Sheet2", "F"+strconv.Itoa(pos+2), ages.Gender)
					f.SetCellValue("Sheet2", "G"+strconv.Itoa(pos+2), ages.Pregnant)
					f.SetCellValue("Sheet2", "H"+strconv.Itoa(pos+2), ages.Offsprings)
					f.SetCellValue("Sheet2", "I"+strconv.Itoa(pos+2), ages.BornDate)
					f.SetCellValue("Sheet2", "J"+strconv.Itoa(pos+2), ages.DeathDate)
				}

				f.SetActiveSheet(index2)

				var weekindex int = 2

				f.SetCellValue("Sheet3", "A1", "Week")
				f.SetCellValue("Sheet3", "B1", "Population Alive")

				for pos1, num := range populationgrowth {
					f.SetCellValue("Sheet3", "A"+strconv.Itoa(pos1+2), weekindex)
					f.SetCellValue("Sheet3", "B"+strconv.Itoa(pos1+2), num)
					weekindex = weekindex + 2
				}

				f.SetActiveSheet(index3)
				f.SetActiveSheet(index)

				if err := f.SaveAs(excelfileexportname); err != nil {
					fmt.Println(err)
				}

				fmt.Println("Sim Finished")
				finished = 1
			}

		} else {
			// dont want excel files. show results some other way
			fmt.Println("Not exported to excel")
			finished = 1
		}
	}
}

func timer1() {
	timer := time.NewTimer(time.Duration(simulationtime) * time.Second)
	<-timer.C
	timerring = 1
}

func startsimulation() {
	fmt.Println("Starting")
	populationage()
	go timer1()
	deathsimulation()

}

func main() {
	rand.Seed(time.Now().UnixNano())
	myApp := app.New()
	myWindow := myApp.NewWindow("Population Simulator")
	myWindow.Resize(fyne.NewSize(1000, 500))
	myWindow.SetFixedSize(true)

	nameentry := widget.NewEntry()
	nameform := widget.NewFormItem("Population Name: ", nameentry)

	populationcountentry := widget.NewEntry()
	popcountform := widget.NewFormItem("Starting Population Amount: ", populationcountentry)

	infoentry2 := widget.NewLabel("")
	infoform2 := widget.NewFormItem("For rates, enter whole numbers only, no sign (2% = 2):", infoentry2)

	birthrateentry := widget.NewEntry()
	birthrateform := widget.NewFormItem("Birth Rate: ", birthrateentry)

	maxbirthentry := widget.NewEntry()
	maxbirthform := widget.NewFormItem("Max births per female: ", maxbirthentry)

	rarebirthentry := widget.NewEntry()
	rarebirthform := widget.NewFormItem("Chance of having a child after reaching max birth: ", rarebirthentry)

	minageforbirthentry := widget.NewEntry()
	minageforbirthform := widget.NewFormItem("Minimum age for female to give birth (0 - 70): ", minageforbirthentry)

	maxageforbirthentry := widget.NewEntry()
	maxageforbirthform := widget.NewFormItem("Maximum age for female to give birth (0 - 70): ", maxageforbirthentry)

	infoentry := widget.NewLabel("")
	infoform := widget.NewFormItem("Scale is from 0% to 100%", infoentry)

	newborndeathrateentry := widget.NewSlider(0, 2000)
	newborndeathrateform := widget.NewFormItem("Age 0-2 Mortality Rate: ", newborndeathrateentry)

	toddlerdeathrateentry := widget.NewSlider(0, 2000)
	toddlerdeathrateform := widget.NewFormItem("Age 2-5 Mortality Rate: ", toddlerdeathrateentry)

	childmortalityrateentry := widget.NewSlider(0, 2000)
	childmortalityrateform := widget.NewFormItem("Age 5-10 Mortality Rate: ", childmortalityrateentry)

	youngteenmortalityrateentry := widget.NewSlider(0, 2000)
	youngteenmortalityrateform := widget.NewFormItem("Age 10-15 Mortality Rate: ", youngteenmortalityrateentry)

	lateteenmortalityrateentry := widget.NewSlider(0, 2000)
	lateteenmortalityrateform := widget.NewFormItem("Age 15-20 Mortality Rate: ", lateteenmortalityrateentry)

	midtwentiesmortalityrateentry := widget.NewSlider(0, 2000)
	midtwentiesmortalityrateform := widget.NewFormItem("Age 20-25 Mortality Rate: ", midtwentiesmortalityrateentry)

	midthirtiesmortalityrateentry := widget.NewSlider(0, 2000)
	midthirtiesmortalityrateform := widget.NewFormItem("Age 25-35 Mortality Rate: ", midthirtiesmortalityrateentry)

	midfortiesmortalityrateentry := widget.NewSlider(0, 2000)
	midfortiesmortalityrateform := widget.NewFormItem("Age 35-45 Mortality Rate: ", midfortiesmortalityrateentry)

	midfiftiesmortalityrateentry := widget.NewSlider(0, 2000)
	midfiftiesmortalityrateform := widget.NewFormItem("Age 45-55 Mortality Rate: ", midfiftiesmortalityrateentry)

	midsixtiesmortalityrateentry := widget.NewSlider(0, 2000)
	midsixtiesmortalityrateform := widget.NewFormItem("Age 55-65 Mortality Rate: ", midsixtiesmortalityrateentry)

	oldergenmortalityrateentry := widget.NewSlider(0, 2000)
	oldergenmortalityrateform := widget.NewFormItem("Age 65-80 Mortality Rate: ", oldergenmortalityrateentry)

	rareoldergenmortalityrateentry := widget.NewSlider(0, 2000)
	rareoldergenmortalityrateform := widget.NewFormItem("Age 80+ Mortality Rate: ", rareoldergenmortalityrateentry)

	options1 := []string{
		"yes", "no",
	}
	exportoption1entry := widget.NewRadioGroup(options1, func(string) {
	})
	exportoption1form := widget.NewFormItem("Export to Excel?", exportoption1entry)

	options2 := []string{
		"Alive Data Only", "Both",
	}
	exportoption2entry := widget.NewRadioGroup(options2, func(string) {
	})
	exportoption2form := widget.NewFormItem("Export only alive, or both alive and deceased data:", exportoption2entry)

	excelfilenameentry := widget.NewEntry()
	excelfilenameform := widget.NewFormItem("Enter a name to export an excel file: ", excelfilenameentry)

	simulationtimeentry := widget.NewEntry()
	simulationtimeform := widget.NewFormItem("How long the simulation will run (in seconds): ", simulationtimeentry)

	text1 := widget.NewFormItem("! If everyone dies before timer ends, then simulation ends !", widget.NewLabel(""))

	newform := widget.NewForm(nameform, popcountform, infoform2, birthrateform, maxbirthform, rarebirthform, minageforbirthform, maxageforbirthform, infoform, newborndeathrateform,
		toddlerdeathrateform, childmortalityrateform, youngteenmortalityrateform, lateteenmortalityrateform,
		midtwentiesmortalityrateform, midthirtiesmortalityrateform, midfortiesmortalityrateform, midfiftiesmortalityrateform,
		midsixtiesmortalityrateform, oldergenmortalityrateform, rareoldergenmortalityrateform, exportoption1form, excelfilenameform,
		exportoption2form,
		simulationtimeform,
		text1,
	)
	//rightcolumn := container.New(layout.NewVBoxLayout(), widget.NewLabel("Test"))
	leftcolumn := container.New(layout.NewVBoxLayout(), newform)
	submitbutton := widget.NewButton(submitbuttontext, func() {

		if buttonpressed == 0 {
			//buttonpressed = 1
			populationname = nameentry.Text
			if populationname == "" {
				populationname = "defaultname"
			}

			populationcount = stringtoint(populationcountentry.Text)
			if populationcount == 0 {
				populationcount = 100
			}

			birthrate = stringtoint(birthrateentry.Text)
			maxbirthcount = stringtoint(maxbirthentry.Text)
			if maxbirthcount == 0 {
				maxbirthcount = 1
			}
			rarebirthprob = stringtoint(rarebirthentry.Text)
			newbornmortalityrate = int(newborndeathrateentry.Value)
			toddlermortalityrate = int(toddlerdeathrateentry.Value)
			childmortalityrate = int(childmortalityrateentry.Value)
			youngteenmortalityrate = int(youngteenmortalityrateentry.Value)
			lateteenmortalityrate = int(lateteenmortalityrateentry.Value)
			midtwentiesmortalityrate = int(midtwentiesmortalityrateentry.Value)
			midthirtiesmortalityrate = int(midthirtiesmortalityrateentry.Value)
			midfortiesmortalityrate = int(midfortiesmortalityrateentry.Value)
			midfiftiesmortalityrate = int(midfiftiesmortalityrateentry.Value)
			midsixtiesmortalityrate = int(midsixtiesmortalityrateentry.Value)
			oldergenmortalityrate = int(oldergenmortalityrateentry.Value)
			rareoldergenmortalityrate = int(rareoldergenmortalityrateentry.Value)
			excelfileexportname = excelfilenameentry.Text + ".xlsx"
			if excelfileexportname == "" {
				excelfileexportname = "defaultname.xlsx"
			}
			exporttoexcelenabled = exportoption1entry.Selected
			if exporttoexcelenabled == "" {
				exporttoexcelenabled = "yes"
			}

			exportdatachoice = exportoption2entry.Selected
			if exportdatachoice == "" {
				exportdatachoice = "Alive Data Only"
			}
			simulationtime = stringtoint(simulationtimeentry.Text)
			if simulationtime == 0 {
				simulationtime = 30
			}

			minageforbirth = stringtoint(minageforbirthentry.Text)
			maxageforbirth = stringtoint(maxageforbirthentry.Text)
			fmt.Println(populationname, populationcount, birthrate, maxbirthcount, rarebirthprob,
				newbornmortalityrate, toddlermortalityrate, childmortalityrate, youngteenmortalityrate,
				lateteenmortalityrate, midtwentiesmortalityrate, midthirtiesmortalityrate, midfortiesmortalityrate,
				midfiftiesmortalityrate, midsixtiesmortalityrate, oldergenmortalityrate, rareoldergenmortalityrate, exporttoexcelenabled,
				excelfileexportname, exportdatachoice, simulationtime,
			)

			submitbuttontext = "Simulating . . ."

			startsimulation()

		}

	})

	maincontainersub := container.NewVBox(container.New(layout.NewGridLayout(1), leftcolumn), submitbutton)
	maincontainer := container.NewVScroll(maincontainersub)

	// Run Loop
	myWindow.SetContent(maincontainer)
	myWindow.ShowAndRun()

	//myApp.Run()
	tidyUp() // runs when closed.
}

func tidyUp() {
	fmt.Println("Exited")
}

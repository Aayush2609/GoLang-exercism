package main

import "fmt"

// Instructions
// In this exercise you'll be writing code to analyze the production in a car factory.

// Task 1
// Calculate the number of working cars produced per hour

// The cars are produced on an assembly line. The assembly line has a certain speed, that can be changed. The faster the assembly line speed is, the more cars are produced. However, changing the speed of the assembly line also changes the number of cars that are produced successfully, that is cars without any errors in their production.

// Implement a function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, from 0 to 100:

// CalculateWorkingCarsPerHour(1547, 90)
// // => 1392.3
// Note: the return value should be a float64.

// Stuck? Reveal Hints
// Opens in a modal
// Task 2
// Calculate the number of working cars produced per minute

// Task 3
// Calculate the cost of production

// How to debug
// When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that console output will be shown too. You can write to the console using:

// import "fmt"
// fmt.Println("Debug message")
// Note: When debugging with the in-browser editor, using e.g. fmt.Print will not add a newline after the output which can provoke a bug in go test --json and result in messed up test output.


func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64{
	return float64(productionRate) * successRate / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int{
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)/60)
}

func CalculateCost(carsCount int) uint {
	groupCost := float64(carsCount/10)*95000
	individualCost := (carsCount%10)*10000
	return uint(groupCost + float64(individualCost))
}

func main(){
	// fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	// fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println(CalculateCost(37))
}
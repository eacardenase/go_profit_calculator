package main

import "fmt"

func main() {
	fmt.Println("Welcome to Profit Calculator")

	revenue := getUserInput("Enter revenue: ")
	expenses := getUserInput("Enter expenses: ")
	taxRate := getUserInput("Enter tax rate (%): ")

	earningsBeforeTax, profit, ratio := calculateFinantials(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
}

func getUserInput(text string) float64 {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinantials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	return earningsBeforeTax, profit, ratio
}

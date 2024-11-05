package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Profit Calculator")

	revenue, revenueErr := getUserInput("Enter revenue: ")

	if revenueErr != nil {
		fmt.Println("Error:", revenueErr)

		return
	}

	expenses, expensesErr := getUserInput("Enter expenses: ")

	if expensesErr != nil {
		fmt.Println("Error:", expensesErr)

		return
	}

	taxRate, taxErr := getUserInput("Enter tax rate (%): ")

	if taxErr != nil {
		fmt.Println("Error:", taxErr)

		return
	}

	earningsBeforeTax, profit, ratio := calculateFinantials(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Profit Ratio: %.2f\n", ratio)
}

func getUserInput(text string) (float64, error) {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}

func calculateFinantials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	storeResults(earningsBeforeTax, profit, ratio)

	return earningsBeforeTax, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf(`Earnings Before Tax: %.2f
Profit: %.2f
Profit Ratio: %.2f`, ebt, profit, ratio)

	os.WriteFile("results.txt", []byte(results), 0644)
}

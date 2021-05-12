// Created by: haokai
// Created on: May 2021
// This program displays, "Volume-of-a-Pyramid"

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function displays currency
	accountingFormater := accounting.Accounting{Precision: 2}
	// This function does addition
	var length float64
	var width float64
	var height float64
	var volume float64
	// input
	fmt.Println("This program is about Volume of a Pyramid program.")
	fmt.Println()
	fmt.Print("Enter the number of length(cm): ")
	fmt.Scanln(&length)
	fmt.Print("Enter the number of width(cm): ")
	fmt.Scanln(&width)
	fmt.Print("Enter the number of height(cm): ")
	fmt.Scanln(&height)
	fmt.Println()
	// process
	volume = length * width * height / 3
	// output
	fmt.Println("Your pay will be:", accountingFormater.FormatMoney(volume), "cm³")
	fmt.Println("\nDone.")
}

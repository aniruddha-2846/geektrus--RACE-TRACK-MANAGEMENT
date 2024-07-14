package main

import (
	"bufio"
	"fmt"
	"geektrust/subpackages"
	"os"
	"strings"
	"time"
)

func getFirstWord(s string) string {
	words := strings.Fields(s)
	if len(words) > 0 {
		return words[0]
	}
	return ""
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please give the input file path")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening the file")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//instance of REGULAR struct

	regular := subpackages.Regular{
		CurrBikes:         0,
		CurrCars:          0,
		CurrSUVs:          0,
		Customers:         make(map[string]time.Time),
		VehicleInfo:       make(map[string]subpackages.Vehicle),
		AdditionalRevenue: 0,
		TotalRevenue:      0,
	}

	//instance of VIP struct
	vip := subpackages.VIP{
		CurrCars:          0,
		CurrSUVs:          0,
		Customers:         make(map[string]time.Time),
		VehicleInfo:       make(map[string]subpackages.Vehicle),
		AdditionalRevenue: 0,
		TotalRevenue:      0,
	}
	for scanner.Scan() {
		line := scanner.Text()
		firstWord := getFirstWord(line)

		switch firstWord {
		case "BOOK":
			subpackages.ProcessBook(line, &regular, &vip)
		case "ADDITIONAL":
			subpackages.ProcessAdditional(line, &regular, &vip)
		case "REVENUE":
			subpackages.ProcessRevenue(&regular, &vip)

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading the file")
		os.Exit(1)
	}
}

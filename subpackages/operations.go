package subpackages

import (
	"fmt"
	"strings"
	"time"
)

type Regular struct {
	CurrBikes         int
	CurrCars          int
	CurrSUVs          int
	TotalBikes        int
	TotalCars         int
	TotalSUVs         int
	Customers         map[string]time.Time
	VehicleInfo       map[string]Vehicle
	AdditionalRevenue int
	TotalRevenue      int
}

type VIP struct {
	CurrCars          int
	CurrSUVs          int
	TotalCars         int
	TotalSUVs         int
	Customers         map[string]time.Time
	VehicleInfo       map[string]Vehicle
	AdditionalRevenue int
	TotalRevenue      int
}

func ProcessBook(line string, reg *Regular, vip *VIP) {
	//get the booking
	words := strings.Fields(line)
	callie := words[0]
	vehicleType := words[1]
	vehicleNumber := words[2]
	bookingTime := words[3]

	//validate the time
	if !CheckTime(bookingTime, callie) {
		println("INVALID_ENTRY_TIME")
		return
	}

	//check for availability
	if !CheckAvailability(vehicleType, vehicleNumber, bookingTime, reg, vip) {
		println("RACETRACK_FULL")
	} else {
		AddDetails(vehicleType, vehicleNumber, bookingTime, reg, vip)
		println("SUCCESS")
	}

}

func ProcessAdditional(line string, reg *Regular, vip *VIP) {
	//get the additional details
	words := strings.Fields(line)
	vehicleNumber := words[1]
	exitTime := words[2]
	callie := words[0]
	//validate the time
	if !CheckTime(exitTime, callie) {
		println("INVALID_EXIT_TIME")
		return
	}
	// check for availability
	if !CheckAvailability("", vehicleNumber, exitTime, reg, vip) {
		fmt.Printf("RACETRACK_FULL %s\n", vehicleNumber)
	} else {
		CalculateAdditionalCharges(vehicleNumber, exitTime, reg, vip)
		println("SUCCESS")
	}
}

func ProcessRevenue(reg *Regular, vip *VIP) {
	//Revenue from Regular Track
	reg.TotalRevenue = (reg.TotalBikes * REGULAR_BIKE_COST * MinimumTrackDuration) + (reg.TotalCars * REGULAR_CAR_COST * MinimumTrackDuration) + (reg.TotalSUVs * REGULAR_SUV_COST * MinimumTrackDuration)
	reg.TotalRevenue += reg.AdditionalRevenue
	print(reg.TotalRevenue)
	print(" ")
	//Revenue from VIP track
	vip.TotalRevenue = (vip.TotalCars * VIP_CAR_COST * MinimumTrackDuration) + (vip.TotalSUVs * VIP_SUV_COST * MinimumTrackDuration)
	print(vip.TotalRevenue)
}

package subpackages

import (
	"fmt"
	"time"
)

// function to validae the input time
func CheckTime(bookingTime string, callie string) bool {
	//check the HH:MM format
	layout1 := "15:04"
	t, err := time.Parse(layout1, bookingTime)
	if err != nil {
		return false
	}
	//check the constraints
	if t.Before(trackStartingTime) || t.After(trackClosingTime) {
		return false
	}
	if callie == "BOOK" && t.After(lastBookingTime) {
		return false
	}
	return true
}

// function to check the availablity of slots for a booking or extension(ADDITIONAL)
func CheckAvailability(vehicleType string, vehicleNumber string, bookingTime string, reg *Regular, vip *VIP) bool {
	booktime, err := time.Parse("15:04", bookingTime)
	if err != nil {
		fmt.Printf("time parse error: %s", err)
	}
	//checking current vehicles in each track
	var vtype Vehicle
	if vehicleType == "" {
		vtype = GetVehicleTypeFromVehicleNumber(vehicleNumber, reg, vip)
	} else {
		vtype = GetVehicleType(vehicleType)
	}

	checkAndUpdate := func(currentCount, limit int, reg *Regular, vip *VIP, vtype Vehicle, booktime time.Time) bool {
		if currentCount < limit {
			return true
		}
		flag := UpdateCustomersMap(reg, vip, vtype, booktime)
		return flag != 0
	}

	switch vtype {
	case BIKE:
		return checkAndUpdate(reg.CurrBikes, REGULAR_BIKE_LIMIT, reg, vip, vtype, booktime)
	case CAR:
		if reg.CurrCars < REGULAR_CAR_LIMIT || vip.CurrCars < VIP_CAR_LIMIT {
			return true
		} else {
			checkAndUpdate(reg.CurrCars, REGULAR_CAR_LIMIT, reg, vip, vtype, booktime)
		}
	case SUV:
		if reg.CurrSUVs < REGULAR_SUV_LIMIT || vip.CurrSUVs < VIP_SUV_LIMIT {
			return true
		} else {
			return checkAndUpdate(reg.CurrSUVs, REGULAR_SUV_LIMIT, reg, vip, vtype, booktime)
		}
	default:
		return false
	}
	return false
}

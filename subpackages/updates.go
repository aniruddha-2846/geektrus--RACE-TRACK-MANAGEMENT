package subpackages

import "time"

// this function is to delete all the vehicles after 3 hours of time duration
func UpdateCustomersMap(reg *Regular, vip *VIP, vtype Vehicle, bookingtime time.Time) int {
	flag := 0
	updateCustomerData := func(customers map[string]time.Time, vehicleInfo map[string]Vehicle, currBikes, currCars, currSUVs *int) {
		for key, addTime := range customers {
			if bookingtime.Before(addTime.Add(3 * time.Hour)) {
				continue
			}
			delete(customers, key)
			delete(vehicleInfo, key)

			switch vtype {
			case BIKE:
				*currBikes--
			case CAR:
				*currCars--
			case SUV:
				*currSUVs--
			}
			flag++
		}
	}
	// Update regular customers
	updateCustomerData(reg.Customers, reg.VehicleInfo, &reg.CurrBikes, &reg.CurrCars, &reg.CurrSUVs)
	if flag == 0 {
		updateCustomerData(vip.Customers, vip.VehicleInfo, nil, &vip.CurrCars, &vip.CurrSUVs)
	}
	return flag
}

//function to register the detail of a successfull booking
func AddDetails(vehicleType string, vehicleNumber string, bookingTime string, reg *Regular, vip *VIP) {
	vtype := GetVehicleType(vehicleType)
	bookingtime, err := time.Parse("15:04", bookingTime)
	if err != nil {
		return
	}

	updateTrackParameter := func(customerMap map[string]time.Time, vehicleInfoMap map[string]Vehicle, currCount, totalCount *int) {
		*currCount++
		*totalCount++
		customerMap[vehicleNumber] = bookingtime
		vehicleInfoMap[vehicleNumber] = vtype
	}

	switch vtype {
	case BIKE:
		updateTrackParameter(reg.Customers, reg.VehicleInfo, &reg.CurrBikes, &reg.TotalBikes)
	case CAR:
		if reg.CurrCars >= REGULAR_CAR_LIMIT {
			updateTrackParameter(vip.Customers, vip.VehicleInfo, &vip.CurrCars, &vip.TotalCars)
		} else {
			updateTrackParameter(reg.Customers, reg.VehicleInfo, &reg.CurrCars, &reg.TotalCars)
		}
	case SUV:
		if reg.CurrSUVs >= REGULAR_SUV_LIMIT {
			updateTrackParameter(vip.Customers, vip.VehicleInfo, &vip.CurrSUVs, &vip.TotalSUVs)
		} else {
			updateTrackParameter(reg.Customers, reg.VehicleInfo, &reg.CurrSUVs, &reg.TotalSUVs)
		}
	}

}

package subpackages

import "time"

//function to calculate charges of additional stay
func CalculateAdditionalCharges(vehicleNumber string, exitTime string, reg *Regular, vip *VIP) {
	defaultExitTime := GetDefaultExitTime(vehicleNumber, reg, vip)
	newExitTime, err := time.Parse("15:04", exitTime)
	if err != nil {
		return
	}
	additonalTime := newExitTime.Sub(defaultExitTime).Minutes()
	if additonalTime > float64(MaximumFreeExtraDuration) {
		if int(additonalTime)%dividerToCalculateInHours == 0 {
			reg.AdditionalRevenue += int(additonalTime/float64(dividerToCalculateInHours)) * AdditionalChargePerHour
		} else {
			reg.AdditionalRevenue += (int(additonalTime/float64(dividerToCalculateInHours)) + 1) * AdditionalChargePerHour
		}
	}
}

//function to get the default-exit-time of vehicle
func GetDefaultExitTime(vehicleNumber string, reg *Regular, vip *VIP) time.Time {
	result := func(customers map[string]time.Time) (time.Time, bool) {
		defaultTime, found := customers[vehicleNumber]
		return defaultTime, found
	}
	if defaultTime, found := result(reg.Customers); found {
		return defaultTime.Add(3 * time.Hour)
	}
	if defaultTime, found := result(vip.Customers); found {
		return defaultTime.Add(3 * time.Hour)
	}
	return time.Now() //trivial case
}

//function ot get the vehicle-type by a string
func GetVehicleType(vehicleType string) Vehicle {
	switch vehicleType {
	case "BIKE":
		return BIKE
	case "CAR":
		return CAR
	case "SUV":
		return SUV
	default:
		return 4
	}
}

//function to get the behicle-type by vehicle Number
func GetVehicleTypeFromVehicleNumber(vehicleNumber string, reg *Regular, vip *VIP) Vehicle {
	result := func(vehicleInfo map[string]Vehicle) (Vehicle, bool) {
		vtype, found := vehicleInfo[vehicleNumber]
		return vtype, found
	}
	if vtype, found := result(reg.VehicleInfo); found {
		return vtype
	}
	if vtype, found := result(vip.VehicleInfo); found {
		return vtype
	}
	return 4
}

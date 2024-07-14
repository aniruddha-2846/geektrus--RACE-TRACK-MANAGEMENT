package subpackages

import "time"

type Vehicle int

func MustParse(layout, value string) time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return t
}

const (
	BIKE Vehicle = iota + 1
	CAR
	SUV
)

var (
	layout2           = "15:04:05"
	trackStartingTime = MustParse(layout2, "13:00:00")
	lastBookingTime   = MustParse(layout2, "17:00:00")
	trackClosingTime  = MustParse(layout2, "20:00:00")
)

var (
	MinimumTrackDuration      = 3  //in hours
	MaximumFreeExtraDuration  = 15 //in minutes //add this in the whole code later
	AdditionalChargePerHour   = 50 //in Rupees
	dividerToCalculateInHours = 60 //60 minutes in an hour
)

// regular track per vehicle limits
var (
	REGULAR_BIKE_LIMIT = 4
	REGULAR_CAR_LIMIT  = 2
	REGULAR_SUV_LIMIT  = 2
)

// regular price chart
var (
	REGULAR_BIKE_COST = 60
	REGULAR_CAR_COST  = 120
	REGULAR_SUV_COST  = 200
)

// vip track limits per vehicle
var (
	VIP_CAR_LIMIT = 1
	VIP_SUV_LIMIT = 1
)

// vip price chart
var (
	VIP_CAR_COST = 250
	VIP_SUV_COST = 300
)

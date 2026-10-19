package cars

import "math"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var Sucess float64 

    Sucess = float64(successRate) * float64(productionRate) / float64(100)

    return Sucess
    
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var Hours float64

    Hours = float64(productionRate) / float64(60) * successRate / float64(100)
    
    return int(Hours)
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var dec float64

    un := float64(carsCount) / 10
    
	inPart, fracPart := math.Modf(un)

    dec = inPart * float64(9500) + fracPart * float64(10000) 

    dec = dec * float64(10)
    
    return uint(dec)
}

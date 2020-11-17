package fuel

func CalculateFuelRequired(n int64) int64 {
	return int64(n/3 - 2)
}

func CalculateFuelRequiredWithFuelForFuel(moduleMass int64) int64 {
	fuelRequired := CalculateFuelRequired(moduleMass)
	var fuelRequiredWithFuelForFuel int64
	for {
		if fuelRequired <= 0 {
			break
		}
		fuelRequiredWithFuelForFuel += fuelRequired
		fuelRequired = CalculateFuelRequired(fuelRequired)
	}
	return fuelRequiredWithFuelForFuel
}

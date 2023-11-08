package main

func getCar(carModel string) (ICar) {
	if carModel == "Tesla Model S" {
		return newTeslaS()
	} else {
		return newTeslaX()
	}
}
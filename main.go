package main

import "fmt"

const molarMassCH3COOH float64 = 60.0
const molarMassNaHCO3 float64 = 84.0 * 0.10
const molarMassCO2 float64 = 44
const volumicMassCO2 float64 = 1.87

func main() {
	AcidQuantityMl := 150.0 // Quantité d'acide acétique en ml
	coeff := AcidQuantityMl / molarMassCH3COOH
	bicarbonateQuantityG := CalculateBicarbonateQuantity(coeff)
	cO2Quantity := CalculateCO2Quantity(coeff)
	cO2Volume := CalculateCO2Volume(cO2Quantity) // litre

	fmt.Printf("bicabonateQuantityG: %.2fg, cO2Volume = %.2f\n", bicarbonateQuantityG, convertM3ToLiter(cO2Volume))
	fmt.Println(convertM3ToLiter(0.81632653061))
}

func CalculateBicarbonateQuantity(coeff float64) float64 {
	return coeff * molarMassNaHCO3
}

func CalculateCO2Quantity(coeff float64) float64 {
	return coeff * molarMassCO2
}

func CalculateCO2Volume(cO2Quantity float64) float64 {
	return (cO2Quantity / 1000) * volumicMassCO2
}

func convertM3ToLiter(volume float64) float64 {
	return volume * 1000
}

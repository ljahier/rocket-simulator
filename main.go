package main

import (
	"fmt"
	"math"
)

func main() {
	bottleVolume := 1.0 // en litres

	coefficients := calculateStoichiometricCoefficients(bottleVolume)

	fmt.Printf("Coefficients stoechiométriques pour 1 litre - NaHCO3: %d g, CH3COOH: %.2f ml, CO2: %d ml\n",
		coefficients.nahco3Mass, coefficients.ch3coohVolume, coefficients.co2Volume)
}

type StoichiometricCoefficients struct {
	nahco3Mass    int
	ch3coohVolume float64
	co2Volume     int
}

func calculateStoichiometricCoefficients(bottleVolume float64) StoichiometricCoefficients {
	// Masse molaire en g/mol
	molarMassNaHCO3 := 84.0
	molarMassCH3COOH := 60.0

	// Coefficient stoechiométrique
	coeffCO2 := 1

	// Calcul des moles de CO2
	molesCO2 := bottleVolume * float64(coeffCO2)

	// Calcul des masses de réactifs et du volume de CH3COOH
	massNaHCO3 := molesCO2 * molarMassNaHCO3
	volumeCH3COOH := molesCO2 * molarMassCH3COOH * 1000.0 // Convertir moles en ml

	return StoichiometricCoefficients{
		nahco3Mass:    int(math.Round(massNaHCO3)),
		ch3coohVolume: volumeCH3COOH,
		co2Volume:     int(math.Round(bottleVolume * 1000.0)), // Convertir litres en ml
	}
}

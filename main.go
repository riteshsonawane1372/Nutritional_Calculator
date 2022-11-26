package main

import (
	"fmt"
	
)

func main() {

	ns := GetNuritionalScore(NutritionalData{
		Engery:              0,
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcidsGram(500),
		Sodium:              SodiumMilligram(60),
		Fruits:              FruitsPercent(4),
		Fibre:               FibreGram(5),
		Protein:             ProteinGram(2),
		IsWater:             false,
	}, Food)

	fmt.Printf("Nutritonal Score : %d \n",ns.Value)
	fmt.Printf("NutriScore : %s\n",ns.GetNutriScore())

}
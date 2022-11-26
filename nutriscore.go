package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

// Above const are of type ScoreType 
// Iota will give value 1
// then 2 3 and so on they all below to scoretype 
// It is like enum of type ScoreType 

// The Score will be +ve or -ve and it will be a value
type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

var scoreToLetter = [] string{"A","B","C","D"}





// Below are the Units

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcidsGram float64

type SodiumMilligram float64

type FruitsPercent float64

type FibreGram float64

type ProteinGram float64

// Nutritional Data

type NutritionalData struct {
	Engery              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcidsGram
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

// Get Points Method For all

func (e EnergyKJ) GetPoints(st ScoreType) int {

	if st ==Beverage{
		return getPointsFromRange(float64(e),engeryLevelBeverage)
	}
	return getPointsFromRange(float64(e),energyLevels)

}

func (s SugarGram) GetPoints(st ScoreType) int {

	if st==Beverage{
		return getPointsFromRange(float64(s),sugarLevelBeverage)
	}
	return getPointsFromRange(float64(s),sugarLevels)
}

func (sfa SaturatedFattyAcidsGram) GetPoints(st ScoreType) int {

	return getPointsFromRange(float64(sfa),SaturatedFattyAcidsLevels)

}

func (so SodiumMilligram) GetPoints(st ScoreType) int {	

	return getPointsFromRange(float64(so),sodiumLevels)

}

func (fr FruitsPercent) GetPoints(st ScoreType) int {

	if st==Beverage {
		if fr>80{
			return 10
		}else if fr>40{
			return 4
		}else if fr>20{
			return 2
		}
		return 0

	}
	if fr>80{
		return 5
	}else if fr>60{
		return 2	
	}else if fr>40{
		return 1
	}
	return 0
	
}

func (fi FibreGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(fi),fibreLevels)
}

func (p ProteinGram) GetPoints(st ScoreType) int {
	return getPointsFromRange(float64(p),ProteinLevels)
}

// Call From Main Function 

func EnergyFromKcal(kcal float64) EnergyKJ{

	return EnergyKJ(kcal*4.184)
}

func SodiumFromSalt(salt float64) SodiumMilligram{

	return SodiumMilligram(salt/2.5)

}

// Slices For Level of Kcal 

var energyLevels = 					[]float64{3350,3015,2680,2345,2010,1675,1340,1005,670,335}
var sugarLevels =   				[]float64 {45,60,36,31,27,22.5,18,13.5,9,4.5}
var SaturatedFattyAcidsLevels =		[] float64{10,9,8,7,6,5,4,3,2,1}
var sodiumLevels = 					[]float64{900,810,720,630,540,450,360,270,180,90}
var fibreLevels = 					[]float64{4.7,3.7,2.8,1.9,0.9}	
var ProteinLevels = 				[]float64{8,6.4,4.8,3.2,1.6}


var engeryLevelBeverage=			[]float64{270,240,210,180,150,120,90,60,30,0}
var sugarLevelBeverage=				[]float64{13.5,12,10.5,9,7.5,6,4.5,3,2.5,1,0}




func GetNuritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	var positive, Negative, value int

	if st != Water {
		fruitPoint := n.Fruits.GetPoints(st)
		fibrePoint := n.Fibre.GetPoints(st)
		Protein := n.Protein.GetPoints(st)

		Negative = int(n.Engery.GetPoints(st)) + int(n.Sugars.GetPoints(st)) + int(n.SaturatedFattyAcids.GetPoints(st)) + int(n.Sodium.GetPoints(st))

		positive = fibrePoint + Protein + fruitPoint

		if st == Cheese{
			value = Negative-positive
		}else{
			if Negative>=11 && fruitPoint<5{
				value = Negative-positive-fruitPoint
			}else{
				value = Negative-positive
			}
		}

	}
	


	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  Negative,
		ScoreType: st,
	}
}


// this funciton will get us the points from the range we created which are levels 
func getPointsFromRange(v float64,steps[]float64) int {
	lenSteps:= len(steps)

	for i,l:=range steps{
		if v>l{
			return lenSteps-i;
		}
	}
	return 0
}


func (ns NutritionalScore) GetNutriScore() string {

	if ns.ScoreType==Food{
		return scoreToLetter[getPointsFromRange(float64(ns.Value),[]float64{18,10,2,-1})]
	}

	if ns.ScoreType==Water{
		return scoreToLetter[0]
	}
	return scoreToLetter[getPointsFromRange(float64(ns.Value),[]float64{9,5,1,-2})]

}	
package main

// import (
// 	"fmt"
// )

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value int
	Positive int // fibrePoints, protein, fruitPoints, 
	Negative int // Energy, Sugars, SaturatedFattyAcids, Sodium
	ScoreType ScoreType
}

type EnergyKJ float64

type SugarGram float64

type SaturatedFattyAcid float64

type SodiumMilligram float64

type FruitsPercent float64

type FibreGram float64

type ProteinGram float64

var energyLevels = []float64{3350, 3015, 2680, 2345, 2010, 1675, 1340, 1005, 675, 335} // kJ/100g
var sugarLevels = []float64{45, 40, 36, 31, 27, 22.5, 18, 13.5, 9, 4.5} // g/100g
var saturatedFattyAcidLevels = []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1} //g/100g
var sodiumLevels = []float64{900, 810, 720, 630, 540, 450, 360, 270, 180, 90} // mg/100g
var fibreLevels = []float64{4.7, 3.7, 2.8, 1.9, 0.9}
var proteinLevels = []float64{8, 6.4, 4.8, 3.2, 1.6}

var energyLevelsBeverages = []float64{270, 240, 210, 180, 150, 120, 90, 60, 30, 0}
var sugarLevelsBeverages = []float64{13.5, 12, 10.5, 9, 7.5, 6, 4.5, 3, 1.5, 0}
var scoreToLetter = []string{"A", "B", "C", "D", "E", "F"}

type NutritionData struct {
	Energy EnergyKJ
	Sugars SugarGram
	SaturatedFattyAcids SaturatedFattyAcid
	Sodium SodiumMilligram
	Fruits FruitsPercent
	Fiber FibreGram
	Protein ProteinGram
	IsWater bool
}

func GetNutritionalScore(n NutritionData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)

		negative = n.Energy.GetPoints(st) + n.Sugars.GetPoints(st) + n.SaturatedFattyAcids.GetPoints() + n.Sodium.GetPoints()
		positive = n.Fiber.GetPoints() + n.Protein.GetPoints() + fruitPoints

		if st == Cheese {
			value = negative - positive
		} else {
			if negative >= 11 && fruitPoints < 5 {
				value = negative - positive - fruitPoints
			} else {
				value = negative - positive
			}
		}
	} 

	return NutritionalScore{
		value,
		positive,
		negative,
		st,
	}
}

// GetNutriScore alt
func (n *NutritionalScore) GetNutritionGrade() string {
	if n.ScoreType == Food {
		return scoreToLetter[getPointsFromRange(float64(n.Value), []float64{18, 10, 2, -1})]
	}

	if n.ScoreType == Water {
		return scoreToLetter[0]
	}

	return scoreToLetter[getPointsFromRange(float64(n.Value),[]float64{9, 5, 1, -2})]
}

func (e EnergyKJ) GetPoints(st ScoreType) int {
	if st == Water {
		return getPointsFromRange(float64(e), energyLevelsBeverages)
	}

	return getPointsFromRange(float64(e), energyLevels)
}

func (s SugarGram) GetPoints(st ScoreType) int {
	if(st == Water) {
		return getPointsFromRange(float64(s), sugarLevelsBeverages)
	}

	return getPointsFromRange(float64(s), sugarLevels)
}

func (sfa SaturatedFattyAcid) GetPoints() int {
	return getPointsFromRange(float64(sfa), saturatedFattyAcidLevels)
}

func (s SodiumMilligram) GetPoints() int {
	return getPointsFromRange(float64(s), sodiumLevels)
}

func (fg FibreGram) GetPoints() int {
	return getPointsFromRange(float64(fg), fibreLevels)
}

func (f FruitsPercent) GetPoints(st ScoreType) int {
	if st == Beverage {
		if f > 80 {
			return 10
		} else if f > 60 {
			return 4
		} else if f > 40 {
			return 2
		}
		return 0
	}

	if f > 80 {
		return 5
	} else if f > 60 {
		return 2
	} else if f > 40 {
		return 1
	}
	return 0
}

func (pg ProteinGram) GetPoints() int {
	return getPointsFromRange(float64(pg), proteinLevels)
}

// convert Kcal to KJoules
func EngergyFromKcal(kcal float64) EnergyKJ {
	return EnergyKJ(kcal * 4.184)
}

func SodiumFromSalt(saltMg float64) SodiumMilligram {
	return SodiumMilligram(saltMg/2.5);
}

func getPointsFromRange(val float64, steps []float64) int {
	size := len(steps)

	for i, lvl := range(steps) {
		if (val >= lvl) {
			return size - i
		}
	}

	return 0
}

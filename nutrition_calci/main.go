package main

import (
  "fmt"
)

func main() {
  ns := GetNutritionalScore(NutritionData{
    100,
    10,
    2,
    500,
    60,
    4,
    2,
    false,
  }, Food)

  fmt.Println("Nutritional score: ", ns)
  fmt.Println("Nutrtion grade:", ns.GetNutritionGrade())
}
package generators

import (
	"math/rand"
)

var foods = []string{
	"Yam", "Rice", "Pasta", "Noodles", "Fish", "Beans",
	"Onions", "Oats", "Pepper", "Tomato", "Beef", "Pork",
	"Chicken", "Eggs", "Avocado", "Garri", "Plantain",
	"Okro",
}

func GenerateRandomCountry() string {
	return foods[rand.Intn(len(foods))]
}

// data, err := ioutil.ReadFile(fmt.Sprintf("teams/%s", league.Teams))
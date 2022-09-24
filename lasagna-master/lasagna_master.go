package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {

	if averagePrepTime == 0 {
		averagePrepTime = 2
	}

	return averagePrepTime * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingredientsFriend []string, ingredients []string) {
	ingredients[len(ingredients)-1] = ingredientsFriend[len(ingredientsFriend)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountFor2Portions []float64, numberOfPortions int) []float64 {
	var scaledRecipe []float64

	for i := 0; i < len(amountFor2Portions); i++ {
		scaledRecipe = append(scaledRecipe, amountFor2Portions[i]*float64(numberOfPortions)/2)
	}

	return scaledRecipe
}

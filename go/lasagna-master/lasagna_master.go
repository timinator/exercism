package lasagna

func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		return len(layers) * 2
	} else {
		return len(layers) * prepTimePerLayer
	}
}

func Quantities(layers []string) (int, float64) {
	noodleCount := 0
	sauceCount := 0.0
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodleCount += 50
		} else if layers[i] == "sauce" {
			sauceCount += 0.2
		}
	}
	return noodleCount, sauceCount
}

func AddSecretIngredient(friendList []string, myList []string) []string {
	return append(myList, friendList[len(friendList)-1])
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	newAmounts := []float64{}
	multiple := float64(portions / 2)
	for i := 0; i < len(amounts); i++ {
		newAmounts = append(newAmounts, amounts[i]*multiple)
	}
	return newAmounts
}

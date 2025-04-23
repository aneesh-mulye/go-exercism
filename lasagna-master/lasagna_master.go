package lasagna

func PreparationTime(layerNames []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return timePerLayer * len(layerNames)
}

func Quantities(layerNames []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, layer := range layerNames {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(fing []string, ming []string) {
	ming[len(ming)-1] = fing[len(fing)-1]
}

func ScaleRecipe(orig []float64, portions int) []float64 {
	ret := make([]float64, 0, len(orig))
	for _, quantity := range orig {
		ret = append(ret, quantity*float64(portions)/2.0)
	}
	return ret
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

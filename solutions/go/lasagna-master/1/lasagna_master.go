package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodleCnt := 0
	sauceCnt := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleCnt += 50
		}
		if layer == "sauce" {
			sauceCnt += 0.2
		}
	}
	return noodleCnt, sauceCnt
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secret := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, person int) []float64 {
	res := make([]float64, len(quantities))
	cnt := float64(person) / 2.0
	for i, v := range quantities {
		res[i] = v * cnt
	}
	return res
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

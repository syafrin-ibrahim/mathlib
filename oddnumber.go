package mathlib

func OddNumber(n ...int) string {
	var state string
	var result int
	for _, e := range n {
		result += e
	}
	if result%2 == 0 {
		state = "genap"
	} else {
		state = "ganjil"
	}

	return state

}

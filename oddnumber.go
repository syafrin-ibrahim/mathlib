package mathlib

func OddNumber(n ...int) string {
	var state string
	var res int
	for _, e := range n {
		res += e
	}
	if res%2 == 0 {
		state = "genap"
	} else {
		state = "ganjil"
	}

	return state

}

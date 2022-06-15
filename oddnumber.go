package mathlib

func OddNumber(n ...int) []string {
	var state []string

	for _, e := range n {
		if e%2 == 0 {
			state = append(state, "genap")
		} else {
			state = append(state, "ganjil")
		}

	}

	return state

}

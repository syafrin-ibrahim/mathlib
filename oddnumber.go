package mathlib

func OddNumber(n int) string {
	var state string
	result := n % 2

	if result == 0 {
		state = "ini bilangan genap"
	} else {
		state = "ini bilangan ganjil"
	}
	return state

}

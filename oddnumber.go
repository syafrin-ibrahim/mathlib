package mathlib

func OddNumber(number int)string{
	var state string
	result := number % 2

	if result == 0 {
		state = "ini bilangan genap"
	}else{
		state = "ini bilangan ganjil"
	}
	return state
}
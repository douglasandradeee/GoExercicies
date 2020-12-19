package mathy

// Multiplier -  makes multiplication operation between two operators.
func Multiplier(x int, y int) int {
	return x * y
}

// Sum - makes a basic sum operation betwen two operators.
func Sum(x float64, y float64) float64 {
	return x + y
}

// Calculation - Calculation performs any type of mathematical operation, just send the desired function.
func Calculation(function func(int, int) int, numberA int, numberB int) int {
	return function(numberA, numberB)
}

// Divider - is a function with return signature that perfomrs split operation.
func Divider(numeroA int, numeroB int) (result int) {
	result = numeroA / numeroB
	return
}

// DividerWithRest - is a division function that returns result and rest of division, functio with return signature.
func DividerWithRest(numeroA int, numeroB int) (result int, rest int) {
	result = numeroA / numeroB
	rest = numeroA % numeroB
	return
}

package custommath

// o fato de começar com Maiúscula informa que é público e visível para quem importar
func SomaCustomMath[T int | float64](a, b T) T {
	return a + b
}

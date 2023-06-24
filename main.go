package go-cod3r-area

import "math"

// PI é uma proporção numérica definida pela relação entre o prímetro de uma circunferência e seu diâmetro.
const PI = 3.1216

// Circunferencia é responsável por calcular a área de uma circunferência
func Circunferencia(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

// Retangulo é responsável por calcular a área de um retângulo
func Retangulo(base, altura float64) float64 {
	return base * altura
}

// Não é uma função visível
func _TrianguloEquilatero(base, altura float64) float64 {
	return (base * altura) / 2
}

package goarea

import "math"

//Pi é uma proporção númerica definida pela relação entre
//o perímetro de uma circunferência e seu dimentro
const Pi = 3.1416

// Cir é responsável por calcular a área da circunferência
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Calculando a área de um retângulo
//Rect é responsável por calcular a área de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não é uma função visível(publica) o _torna ela privada
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}

package estruturas

import "math"

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}
type Triangulo struct {
	Base   float64
	Altura float64
}

type Forma interface {
	Area() float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) * 0.5
}

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Altura + retangulo.Largura)
}

func Area(retangulo Retangulo) float64 {
	return retangulo.Altura * retangulo.Largura
}

package estruturas

import "testing"

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	result := Perimetro(retangulo)
	esp := 40.0

	if result != esp {
		t.Errorf("resultado %.2f esperado %.2f", result, esp)
	}

}

func Test_area(t *testing.T) {
	retangulo := Retangulo{10.0, 10.0}
	result := Area(retangulo)
	esp := 100.0
	if result != esp {
		t.Errorf("resultado %.2f esperado %.2f", result, esp)
	}
}

func Test_Area(t *testing.T) {
	verificaArea := func(t *testing.T, forma Forma, esp float64) {
		t.Helper()
		result := forma.Area()
		if result != esp {
			t.Errorf("resultado %.2f esperado %.2f", result, esp)
		}
	}

	t.Run("Retângulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		verificaArea(t, retangulo, 72.0)
	})

	t.Run("circulos", func(t *testing.T) {
		circulo := Circulo{10}
		verificaArea(t, circulo, 314.1592653589793)
	})
}

/*Isso se chama Test Table Driven Events!*/
func Test_Table_Driven_Events(t *testing.T) {
	testesArea := []struct {
		nome    string
		forma   Forma
		temArea float64
	}{
		{nome: "Retângulo", forma: Retangulo{Largura: 12, Altura: 6}, temArea: 72.0},
		{nome: "Círculo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}
	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			result := tt.forma.Area()
			if result != tt.temArea {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt.forma, result, tt.temArea)
			}
		})
	}

}

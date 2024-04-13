package main

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Run("Coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}

		resultado := Soma(numeros)
		esperado := 15

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

	t.Run("Coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}
		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}

	})
}
func TestSomaTudo(t *testing.T) {

	recebido := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(recebido, esperado) {
		t.Errorf("recebido %v esperado %v", recebido, esperado)
	}
}
func Test_soma_todo_resto(t *testing.T) {
	rest := SomaTodoResto([]int{1, 2}, []int{0, 9})
	esp := []int{2, 9}

	if !reflect.DeepEqual(rest, esp) {
		t.Errorf("recebido %v esperado %v", rest, esp)

	}
}''

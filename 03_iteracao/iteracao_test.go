package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 5)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}
func TestInfinito(t *testing.T) {
	repeticoes := ForInfinito(2000)
	esperado := 2000
	if repeticoes != esperado {
		t.Errorf("esperado '%d' mas obteve '%d'", esperado, repeticoes)
	}
}

/*
Para executar o benchMark no CLI do Go é necessario utilizar o comando

	go test -bench=.

sendo que o ponto é a pasta atual
*/
func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 1500)
	}
}

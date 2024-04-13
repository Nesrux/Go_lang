package main

import "testing"

// Para rodar um Test precisa de 3 coisas
// 1) o nome do arquivo dever algo como **_test.go
// 2) a funcão que vais er testada precisa começar com a palavra Test
// 3) a funcção tem que receber um ponteiro de t *testing.T
func Test_Ola(t *testing.T) {
	resultado := Ola("mundo")
	esperado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("Resultado %s, esperado %s", resultado, esperado)
	}
}

func Test_Ola2(t *testing.T) {
	result := Ola("João")
	esperado := "Olá, João"
	if result != esperado {
		t.Errorf("resultado %s, esperado %s", result, esperado)
	}
}

func Test_ola3(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, res, esp string) {
		t.Helper()
		if res != esp {
			t.Errorf("resultado %s esperado %s", res, esp)
		}
	}

	t.Run("diz oká para as pessoas", func(t *testing.T) {
		res := Ola("Chris")
		esp := "Olá, Chris"
		verificaMensagemCorreta(t, res, esp)
	})
	t.Run("diz Olá mundo quando uma string vazia for passada", func(t *testing.T) {
		res := Ola("")
		esp := "Olá, Mundo"
		verificaMensagemCorreta(t, res, esp)
	})
}

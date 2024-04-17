package maps

import "testing"

func Test_Busca(t *testing.T) {
	dicionario := Dicionario{"Test": "Isso é apenas um teste"}
	t.Run("Palavra conhecida", func(t *testing.T) {
		res, _ := dicionario.Busca("Test")
		esp := "Isso é apenas um teste"

		comparaStrings(t, res, esp)
	})

	t.Run("Palavra desconhecida", func(t *testing.T) {
		_, resultado := dicionario.Busca("desconhecido")
		comparaErro(t, resultado, ErrNaoEncontrado)
	})

}

func Test_Adiciona(t *testing.T) {
	dicionario := Dicionario{}
	dicionario.Adiciona("teste", "isso é apenas um teste")

	esperado := "isso é apenas um teste"
	result, err := dicionario.Busca("teste")
	if err != nil {
		t.Fatal(ErrNaoEncontrado)
	}
	if esperado != result {
		t.Errorf("resultado %s, esperado %s", result, esperado)
	}
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado %s, esperado %s dado %s", resultado, esperado, "test")
	}
}
func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

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
	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "isso é apenas um test"

		dicionario.Adiciona(palavra, definicao)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"

		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "teste novo")

		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})
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

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()
	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("Deveria ter encontrado palavra adicionada: ", err)
	}
	if definicao != resultado {
		t.Errorf("resultado %s, esperado %s", resultado, definicao)
	}
}

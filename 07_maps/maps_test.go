package maps

import "testing"

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado %s, esperado %s dado %s", resultado, esperado, "test")
	}
}
func Test_Busca(t *testing.T) {
	dicionario := map[string]string{"Test": "Isso é apenas um teste"}

	res := Busca(dicionario, "Test")
	esp := "Isso é apenas um teste"

	comparaStrings(t, res, esp)
}

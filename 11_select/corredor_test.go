package corredor

import "testing"

func Test_Corredor(t *testing.T) {
	urlLenta := "http://www.facebook.com"
	urlRapida := "http://www.quii.co.uk"

	esperado := urlRapida
	resultado := Corredor(urlLenta, urlRapida)

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)

	}
}

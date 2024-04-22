package main

import (
	"bytes"
	"testing"
)

func Test_Contagem(t *testing.T) {
	buffer := &bytes.Buffer{}
	Contagem(buffer)

	rest := buffer.String()
	esperado := `3
	2
	1
	Vai!`

	if rest != esperado {
		t.Errorf("Resultado %s, Esperado %s", rest, esperado)
	}
}

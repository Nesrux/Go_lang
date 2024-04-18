package main

import (
	"bytes"
	"testing"
)

func Test_Cumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "João")

	result := buffer.String()
	esp := "Olá, João"

	if result != esp {
		t.Errorf("resultado '%s' esperado '%s'", result, esp)
	}
}

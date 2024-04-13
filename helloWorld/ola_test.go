package main
import "testing"
//Para rodar um Test precisa de 3 coisas
// 1) o nome do arquivo dever algo como **_test.go
// 2) a funcão que vais er testada precisa começar com a palavra Test
// 3) a funcção tem que receber um ponteiro de t *testing.T
func Test_Ola(t *testing.T){
	resultado := Ola();
	esperado := "Olá mundo!"

	if(resultado != esperado){
		t.Errorf("Resultado %s, esperado %s", resultado, esperado)
	}
}